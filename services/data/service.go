package data

import (
	"fmt"
	"github.com/boltdb/bolt"
	"gocloud/services/db"
)

func Read(key string, userId, storageId int) (string, error) {
	var value string
	if checkAccess(userId, storageId) == false {
		return value, fmt.Errorf("user has not access to storage")
	}
	storage, err := openStorage("../../storages/test.db")
	if err != nil {
		return value, err
	}
	value, err = db.ReadRowFromStorage(storage, key, "mybucket") // TODO необходимо ли на этом уровне разделение на бакеты? Учитывая что на каждого юзера отдельное хранилище - не будет ли логичнее бакеты оставить только для разделения данных под капотом?
	if err != nil {
		return value, err
	}
	return value, nil
}

func Write(key string, value string, userId, storageId int) error {
	if checkAccess(userId, storageId) == false {
		return fmt.Errorf("user has not access to storage")
	}
	storage, err := openStorage("../../storages/test.db")
	if err != nil {
		return err
	}
	err = db.WriteRowToStorage(storage, key, value, "mybucket")
	if err != nil {
		return err
	}
	return nil
}

func checkAccess(userId, storageId int) bool {
	return true
}

func openStorage(storagePath string) (*bolt.DB, error) {
	tmpDB, err := bolt.Open(storagePath, 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	return tmpDB, nil
}

func createStorage() {

}

func getStoragePathById() {

}
