package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os"
)

func ReadRowFromStorage(db *bolt.DB, key string, bucket string) (string, error) {
	var value string
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
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

func WriteRowToStorage(db *bolt.DB, key string, value string, bucket string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		if err := bucket.Put([]byte(key), []byte(value)); err != nil {
			return err
		}

		return nil
	})
}

func UpdateRowInStorage(db *bolt.DB, key, newValue string, bucket string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		if bucket.Get([]byte(key)) == nil {
			return fmt.Errorf("key '%s' not found in storage", key)
		}
		return bucket.Put([]byte(key), []byte(newValue))
	})
}

func DeleteRowInStorage(db *bolt.DB, key string, bucket string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		if bucket.Get([]byte(key)) == nil {
			return fmt.Errorf("key '%s' not found in storage", key)
		}
		return bucket.Delete([]byte(key))
	})
}

func createStorage(path string) (bool, error) {
	dbFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return false, err
	}
	defer dbFile.Close()
	return true, nil
}

func GetStorageData(db *bolt.DB, bucket string) (map[string]string, error) {
	data := make(map[string]string)

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		return bucket.ForEach(func(k, v []byte) error {
			data[string(k)] = string(v)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
