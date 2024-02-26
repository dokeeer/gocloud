package controllers

import (
	"context"
	"fmt"
	"gocloud/services"
)

func HandleGRPCRequest(ctx context.Context, request interface{}) {
	// Пример обработки gRPC запроса и вызова функции сервиса
	services.ProcessGRPCRequest(ctx, request)
	fmt.Println("gRPC request handled")
}
