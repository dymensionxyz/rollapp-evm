package repository

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"randomnessgenerator/agent/external"
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

func (db *DB) Get(randID uint64) (Answer, error) {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, randID)
	b, err := db.conn.Get(key, nil)
	if err != nil {
		return Answer{}, fmt.Errorf("get answer from DB: %v", err)
	}
	var a Answer
	a.MustFromBytes(b)
	return a, nil
}

func (db *DB) Save(randID uint64, a Answer) error {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, randID)
	return db.conn.Put(key, a.MustToBytes(), nil)
}

func (db *DB) Close() error {
	return db.conn.Close()
}

type Answer struct {
	external.RandomnessResponse
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
