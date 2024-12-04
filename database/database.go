package database

import (
	"errors"
)

type InMemoryDB struct {
	data        map[string]int       // main storage
	transaction map[string]*int      // transaction storage
	inProgress  bool                 // transaction state
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data:        make(map[string]int),
		transaction: make(map[string]*int),
		inProgress:  false,
	}
}

func (db *InMemoryDB) Get(key string) *int {
	if db.inProgress {
        if _, exists := db.transaction[key]; exists {
            return nil // Return nil for uncommitted keys
        }
	}
	if val, exists := db.data[key]; exists {
		return &val
	}
	return nil
}

func (db *InMemoryDB) Put(key string, value int) error {
	if !db.inProgress {
		return errors.New("no transaction in progress")
	}
	db.transaction[key] = &value
	return nil
}

func (db *InMemoryDB) BeginTransaction() error {
	if db.inProgress {
		return errors.New("transaction already in progress")
	}
	db.inProgress = true
	db.transaction = make(map[string]*int)
	return nil
}

func (db *InMemoryDB) Commit() error {
	if !db.inProgress {
		return errors.New("no transaction in progress")
	}
	for key, value := range db.transaction {
		if value == nil {
			delete(db.data, key)
		} else {
			db.data[key] = *value
		}
	}
	db.inProgress = false
	db.transaction = nil
	return nil
}

func (db *InMemoryDB) Rollback() error {
	if !db.inProgress {
		return errors.New("no transaction in progress")
	}
	db.inProgress = false
	db.transaction = nil
	return nil
}
