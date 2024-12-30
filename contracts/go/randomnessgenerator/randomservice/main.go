package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
	"net/http"
)

type Config struct {
	HTTPServerAddr string
	LevelDBPath    string
}

type RandomService struct {
	db *leveldb.DB
}

type RandomnessResponse struct {
	RequestID  string
	Randomness *big.Int
}

func NewRandomnessService(config Config) (*RandomService, error) {
	rs := &RandomService{}
	db, err := leveldb.OpenFile(config.LevelDBPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	rs.db = db
	return rs, nil
}

func generateRequestID() string {
	return uuid.New().String()
}

func generateRandomness() *big.Int {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("error generating randomness: %v", err)
	}
	randomness := new(big.Int).SetBytes(randomBytes)
	return randomness
}

func (rs *RandomService) handleGenerate(w http.ResponseWriter, r *http.Request) {
	requestID := generateRequestID()
	randomness := generateRandomness()

	data := RandomnessResponse{
		RequestID:  requestID,
		Randomness: randomness,
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		http.Error(w, "error encoding data", http.StatusInternalServerError)
		log.Printf("error encoding data: %v", err)
		return
	}

	err = rs.db.Put([]byte(requestID), buf.Bytes(), nil)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		log.Printf("error writing to db: %v", err)
		return
	}

	response := fmt.Sprintf(`{"requestID": "%s", "randomness": "%s"}`, requestID, randomness.String())
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

func (rs *RandomService) handleRandomness(w http.ResponseWriter, r *http.Request) {
	requestID := r.URL.Query().Get("request_id")
	if requestID == "" {
		http.Error(w, "requestID is required", http.StatusBadRequest)
		return
	}

	data, err := rs.db.Get([]byte(requestID), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			http.Error(w, "requestID not found", http.StatusNotFound)
		} else {
			http.Error(w, "database error", http.StatusInternalServerError)
		}
		log.Printf("error retrieving data for requestID %s: %v", requestID, err)
		return
	}

	var decodedData RandomnessResponse
	dec := gob.NewDecoder(bytes.NewReader(data))
	err = dec.Decode(&decodedData)
	if err != nil {
		http.Error(w, "error decoding data", http.StatusInternalServerError)
		log.Printf("error decoding data: %v", err)
		return
	}

	randomnessStr := decodedData.Randomness.String()

	response := struct {
		RequestID  string `json:"request_id"`
		Randomness string `json:"randomness"`
	}{
		RequestID:  decodedData.RequestID,
		Randomness: randomnessStr,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding response to JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

func (rs *RandomService) StartHTTPServer(ctx context.Context, address string) error {
	server := &http.Server{
		Addr: address,
	}

	http.HandleFunc("/generate", rs.handleGenerate)
	http.HandleFunc("/randomness", rs.handleRandomness)

	go func() {
		<-ctx.Done()
		_ = server.Shutdown(context.Background())
	}()

	log.Printf("starting http server at %s", address)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := Config{
		HTTPServerAddr: ":8081",
		LevelDBPath:    "db1",
	}

	rs, err := NewRandomnessService(config)
	if err != nil {
		log.Fatalf("error initializing randomnessservice: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := rs.StartHTTPServer(ctx, config.HTTPServerAddr); err != nil {
		log.Fatalf("error starting http server: %v", err)
	}
}
