package db

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func ReadRowFromStorage(db *bolt.DB, key string) (string, error) {
	var value string

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("mybucket"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		val := bucket.Get([]byte(key))
		if val == nil {
			return fmt.Errorf("key not found: %s", key)
		}

		value = string(val)
		return nil
	})

	return value, err
}

func WriteRowToStorage(db *bolt.DB, key string, value string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("mybucket"))
		if err != nil {
			return err
		}

		if err := bucket.Put([]byte(key), []byte(value)); err != nil {
			return err
		}

		return nil
	})

	return err
}

func UpdateRowInStorage(db *bolt.DB, key, newValue string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("mybucket"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		if bucket.Get([]byte(key)) == nil {
			return fmt.Errorf("key '%s' not found in storage", key)
		}
		return bucket.Put([]byte(key), []byte(newValue))
	})
}

func DeleteRowInStorage(db *bolt.DB, key string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("mybucket"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		if bucket.Get([]byte(key)) == nil {
			return fmt.Errorf("key '%s' not found in storage", key)
		}
		return bucket.Delete([]byte(key))
	})
}
