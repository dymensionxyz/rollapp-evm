package repository

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type DB struct {
	conn *leveldb.DB
}

func NewLevelDB() (*DB, error) {
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v", err)
	}
	return &DB{conn: db}, nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	return db.conn.Get(key, nil)
}

func (db *DB) Save(key, value []byte) error {
	return db.conn.Put(key, value, nil)
}

func (db *DB) Close() error {
	return db.conn.Close()
}
