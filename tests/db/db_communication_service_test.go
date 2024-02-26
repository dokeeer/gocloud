package db

import (
	"github.com/boltdb/bolt"
	"gocloud/services/db"
	"testing"
)

func TestWriteToAndReadRowFromStorage(t *testing.T) {
	tmpDB, err := bolt.Open("../../storages/test.db", 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpDB.Close()

	key := "testKey"
	value := "testValue"
	if err := db.WriteRowToStorage(tmpDB, key, value); err != nil {
		t.Fatal(err)
	}

	readValue, err := db.ReadRowFromStorage(tmpDB, key)
	if err != nil {
		t.Fatal(err)
	}

	if readValue != value {
		t.Errorf("Expected value %s, got %s", value, readValue)
	}
}

func TestKeyFromStorage_KeyNotFound(t *testing.T) {
	tmpDB, err := bolt.Open("../../storages/test.db", 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpDB.Close()

	_, err = db.ReadRowFromStorage(tmpDB, "nonExistentKey")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestUpdateRowInStorage_Success(t *testing.T) {
	tmpDB, err := bolt.Open("../../storages/test.db", 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpDB.Close()

	if err := db.WriteRowToStorage(tmpDB, "1", "value1"); err != nil {
		t.Fatal(err)
	}

	newValue := "newValue"
	if err := db.UpdateRowInStorage(tmpDB, "1", newValue); err != nil {
		t.Fatal(err)
	}

	updatedValue, err := db.ReadRowFromStorage(tmpDB, "1")
	if err != nil {
		t.Fatal(err)
	}
	if updatedValue != newValue {
		t.Errorf("Expected updated value '%s', got '%s'", newValue, updatedValue)
	}
}

func TestUpdateRowInStorage_KeyNotFound(t *testing.T) {
	tmpDB, err := bolt.Open("../../storages/test.db", 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpDB.Close()

	err = db.UpdateRowInStorage(tmpDB, "nonExistentKey", "value")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDeleteRowInStorage_Success(t *testing.T) {
	tmpDB, err := bolt.Open("../../storages/test.db", 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpDB.Close()

	if err := db.WriteRowToStorage(tmpDB, "1", "value1"); err != nil {
		t.Fatal(err)
	}

	if err := db.DeleteRowInStorage(tmpDB, "1"); err != nil {
		t.Fatal(err)
	}

	_, err = db.ReadRowFromStorage(tmpDB, "1")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDeleteRowInStorage_KeyNotFound(t *testing.T) {
	tmpDB, err := bolt.Open("../../storages/test.db", 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpDB.Close()

	err = db.DeleteRowInStorage(tmpDB, "nonExistentKey")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
