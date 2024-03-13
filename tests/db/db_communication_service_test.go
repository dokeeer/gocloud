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
	if err := db.WriteRowToStorage(tmpDB, key, value, "mybucket"); err != nil {
		t.Fatal(err)
	}

	readValue, err := db.ReadRowFromStorage(tmpDB, key, "mybucket")
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

	_, err = db.ReadRowFromStorage(tmpDB, "nonExistentKey", "mybucket")
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

	if err := db.WriteRowToStorage(tmpDB, "1", "value1", "mybucket"); err != nil {
		t.Fatal(err)
	}

	newValue := "newValue"
	if err := db.UpdateRowInStorage(tmpDB, "1", newValue, "mybucket"); err != nil {
		t.Fatal(err)
	}

	updatedValue, err := db.ReadRowFromStorage(tmpDB, "1", "mybucket")
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

	err = db.UpdateRowInStorage(tmpDB, "nonExistentKey", "value", "mybucket")
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

	if err := db.WriteRowToStorage(tmpDB, "1", "value1", "mybucket"); err != nil {
		t.Fatal(err)
	}

	if err := db.DeleteRowInStorage(tmpDB, "1", "mybucket"); err != nil {
		t.Fatal(err)
	}

	_, err = db.ReadRowFromStorage(tmpDB, "1", "mybucket")
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

	err = db.DeleteRowInStorage(tmpDB, "nonExistentKey", "mybucket")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func testCRUD(t *testing.T) {
	t.Run("TestWriteToAndReadRowFromStorage", TestWriteToAndReadRowFromStorage)
	t.Run("TestKeyFromStorage_KeyNotFound", TestKeyFromStorage_KeyNotFound)
	t.Run("TestUpdateRowInStorage_Success", TestUpdateRowInStorage_Success)
	t.Run("TestUpdateRowInStorage_KeyNotFound", TestUpdateRowInStorage_KeyNotFound)
	t.Run("TestDeleteRowInStorage_Success", TestDeleteRowInStorage_Success)
	t.Run("TestDeleteRowInStorage_KeyNotFound", TestDeleteRowInStorage_KeyNotFound)
}
