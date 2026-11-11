package service

import (
	"context"
	"dc-transaction-service/app/model/entity"
	"dc-transaction-service/app/model/request"
)

type TransactionStatusServiceInterface interface {
	TransactionStatusCreate(ctx context.Context, req request.TransactionStatusRequest) (entity.TransactionStatus, int, error)
	TransactionStatusGetAll(ctx context.Context) ([]entity.TransactionStatus, int, error)
	TransactionStatusGetByID(ctx context.Context, id int) (entity.TransactionStatus, int, error)
	TransactionStatusUpdate(ctx context.Context, id int, req request.TransactionStatusRequest) (int, error)
	TransactionStatusDelete(ctx context.Context, id int) (int, error)
}
