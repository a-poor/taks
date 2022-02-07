package lib

import (
	badger "github.com/dgraph-io/badger/v3"
)

// TaksDB manages the connection to the task DB
type TaksDB struct {
	Path string
	DB   *badger.DB
}

func OpenDB(path string) (*TaksDB, error) {
	opt := badger.DefaultOptions(path)
	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}
	return &TaksDB{
		Path: path,
		DB:   db,
	}, nil
}

func (db *TaksDB) Close() error {
	return db.DB.Close()
}

func (db *TaksDB) GetTask(id string) (*Task, error) {
	var t *Task
	err := db.DB.View(func(txn *badger.Txn) error {
		// Get the value for the given key.
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		// Read the value from the item, if no error occurred.
		err = item.Value(func(val []byte) error {
			t, err = TaskFromBytes(val)
			return err
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Return the unmarshaled task.
	return t, nil
}

func (db *TaksDB) PutTask(t Task) error {
	id, task, err := t.MarshalBytes()
	if err != nil {
		return err
	}
	return db.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(id, task)
	})
}

func (db *TaksDB) DeleteTask(id string) error {
	return db.DB.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(id))
	})
}
