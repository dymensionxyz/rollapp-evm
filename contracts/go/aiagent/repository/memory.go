package repository

import (
	"fmt"
	"sync"
)

type MemoryDB struct {
	data map[string][]byte
	mu   sync.RWMutex
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		data: make(map[string][]byte),
	}
}

func (db *MemoryDB) Get(key []byte) ([]byte, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, exists := db.data[string(key)]
	if !exists {
		return nil, fmt.Errorf("key not found")
	}
	return value, nil
}

func (db *MemoryDB) Save(key, value []byte) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[string(key)] = value
	return nil
}

func (db *MemoryDB) Close() error {
	return nil
}
