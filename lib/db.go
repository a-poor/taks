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
	opt.Logger = nil // TODO: Currently removing the logger. Update this?
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

func (db *TaksDB) NewTaskID() (uint64, error) {
	return 0, nil
}

func (db *TaksDB) Validate() error {
	return db.DB.View(func(txn *badger.Txn) error {
		// Get the value for the given key.
		item, err := txn.Get(MetadataID)
		if err != nil {
			return err
		}

		// Read the value from the item, if no error occurred.
		err = item.Value(func(val []byte) error {
			_, err = MetadataFromBytes(val)
			return err
		})

		return nil
	})
}

func (db *TaksDB) ListTasks() ([]*Task, error) {
	// Create a slice to store the tasks
	var ts []*Task

	// Iterate over the tasks in the DB
	err := db.DB.View(func(txn *badger.Txn) error {
		// Create an iterator for the DB
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		// Iterate over the tasks
		for it.Seek([]byte(TaskPrefix)); it.ValidForPrefix([]byte(TaskPrefix)); it.Next() {
			// Get the item...
			item := it.Item()

			// Attempt to get the value as bytes...
			err := item.Value(func(v []byte) error {
				// Unmarshal the task...
				t, err := TaskFromBytes(v)
				if err != nil {
					return err
				}

				// Success! Add the task to the slice
				ts = append(ts, t)
				return nil
			})

			// Unmarshal failed, return the error
			if err != nil {
				return err
			}
		}
		return nil
	})

	// Error occurred, return it
	if err != nil {
		return nil, err
	}

	// Success! Return the tasks
	return ts, nil
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

func (db *TaksDB) PutTask(t *Task) error {
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

func (db *TaksDB) GetMetadata() (*AppMetadata, error) {
	var m *AppMetadata
	err := db.DB.View(func(txn *badger.Txn) error {
		// Get the value for the given key.
		item, err := txn.Get(MetadataID)
		if err != nil {
			return err
		}

		// Read the value from the item, if no error occurred.
		err = item.Value(func(val []byte) error {
			m, err = MetadataFromBytes(val)
			return err
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Return the unmarshaled metadata.
	return m, nil
}

func (db *TaksDB) PutMetadata(m *AppMetadata) error {
	id, body, err := m.MarshalBytes()
	if err != nil {
		return err
	}
	return db.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(id, body)
	})
}

func (db *TaksDB) DeleteMetadata() error {
	return db.DB.Update(func(txn *badger.Txn) error {
		return txn.Delete(MetadataID)
	})
}
