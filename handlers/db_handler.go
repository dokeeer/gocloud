package handlers

import (
	"encoding/json"
	"gocloud/services/data"
	"net/http"
)

type ReadExpectedData struct {
	StorageId int    `json:"storageId"`
	Key       string `json:"key"`
}

type WriteExpectedData struct {
	StorageId int    `json:"storageId"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type UpdateExpectedData struct {
	StorageId int    `json:"storageId"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type DeleteExpectedData struct {
	StorageId int    `json:"storageId"`
	Key       string `json:"key"`
}

type CreateExpectedData struct {
	StorageName string `json:"StorageName"`
}

func Read(w http.ResponseWriter, r *http.Request) {
	var requestData ReadExpectedData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}
	value, err := data.Read(requestData.Key, 0, requestData.StorageId) // TODO на основании каких данных получаем id пользователя для проверки разрешений на хранилище?
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func Write(w http.ResponseWriter, r *http.Request) {
	var requestData WriteExpectedData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}
	err = data.Write(requestData.Key, requestData.Value, 0, requestData.StorageId) // TODO на основании каких данных получаем id пользователя для проверки разрешений на хранилище?
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Update(http.ResponseWriter, *http.Request) {

}

func Delete(http.ResponseWriter, *http.Request) {

}

func CreateStorage(http.ResponseWriter, *http.Request) {

}
