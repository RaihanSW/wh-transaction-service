package repository

import (
	"context"
	"dc-transaction-service/app/model/entity"
)

type TransactionStatusRepositoryInterface interface {
	InsertTransactionStatus(ctx context.Context, data entity.TransactionStatus) (int, error)
	GetAllTransactionStatus(ctx context.Context) ([]entity.TransactionStatus, error)
	GetTransactionStatusByID(ctx context.Context, id int) (entity.TransactionStatus, error)
	UpdateTransactionStatusByID(ctx context.Context, data entity.TransactionStatus) (int64, error)
	DeleteTransactionStatusByID(ctx context.Context, id int) (int64, error)
}
