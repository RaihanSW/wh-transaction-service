package repository

import (
	"context"
	"dc-transaction-service/app/model/entity"
)

type TransactionRepositoryInterface interface {
	InsertTransaction(ctx context.Context, data entity.Transaction) (string, error)
	GetAllTransactions(ctx context.Context) ([]entity.Transaction, error)
	GetTransactionByUID(ctx context.Context, uid string) (entity.Transaction, error)
	UpdateTransactionByUID(ctx context.Context, data entity.Transaction) (int64, error)
	DeleteTransactionByUID(ctx context.Context, uid string) (int64, error)
}
