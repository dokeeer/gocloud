package controllers

import (
	"fmt"
	"gocloud/services"
	"net/http"
)

func HandleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	// Пример обработки HTTP запроса и вызова функции сервиса
	services.ProcessHTTPRequest(w, r)
	fmt.Fprintf(w, "HTTP request handled")
}
