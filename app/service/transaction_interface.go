package service

import (
	"context"
	"dc-transaction-service/app/model/entity"
	"dc-transaction-service/app/model/request"
)

type TransactionServiceInterface interface {
	TransactionCreate(ctx context.Context, req request.TransactionRequest) (entity.Transaction, int, error)
	TransactionGetAll(ctx context.Context) ([]entity.Transaction, int, error)
	TransactionGetByID(ctx context.Context, uid string) (entity.Transaction, int, error)
	TransactionUpdate(ctx context.Context, uid string, req request.TransactionRequest) (int, error)
	TransactionDelete(ctx context.Context, uid string) (int, error)
}
