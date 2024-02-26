package services

import (
	"context"
	"net/http"
)

// Функция обработки HTTP запроса
func ProcessHTTPRequest(w http.ResponseWriter, r *http.Request) {
	// Логика обработки HTTP запроса
}

// Функция обработки gRPC запроса
func ProcessGRPCRequest(ctx context.Context, request interface{}) {
	// Логика обработки gRPC запроса
}
