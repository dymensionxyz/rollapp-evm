package repository

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type DB struct {
	conn *leveldb.DB
}

func NewLevelDB(dbPath string) (*DB, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v", err)
	}
	return &DB{conn: db}, nil
}

func (db *DB) Get(promptID uint64) (Answer, error) {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, promptID)
	b, err := db.conn.Get(key, nil)
	if err != nil {
		return Answer{}, fmt.Errorf("get answer from DB: %v", err)
	}
	var a Answer
	a.MustFromBytes(b)
	return a, nil
}

func (db *DB) Save(promptID uint64, a Answer) error {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, promptID)
	return db.conn.Put(key, a.MustToBytes(), nil)
}

func (db *DB) Close() error {
	return db.conn.Close()
}

type Answer struct {
	Answer      string `json:"answer,omitempty"`
	MessageID   string `json:"message_id,omitempty"`
	ThreadID    string `json:"thread_id,omitempty"`
	RunID       string `json:"run_id,omitempty"`
	AssistantID string `json:"assistant_id,omitempty"`
}

// MustToBytes converts SubmitPromptResponse to bytes
func (a Answer) MustToBytes() []byte {
	b, err := json.Marshal(a)
	if err != nil {
		panic(fmt.Errorf("marshal answer: %w", err))
	}
	return b
}

// MustFromBytes converts bytes to SubmitPromptResponse
func (a *Answer) MustFromBytes(data []byte) {
	err := json.Unmarshal(data, a)
	if err != nil {
		panic(fmt.Errorf("unmarshal answer: %w", err))
	}
}
