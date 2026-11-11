package repository

import (
	"context"
	"dc-transaction-service/app/model/entity"
)

type GoodsStatusRepositoryInterface interface {
	InsertGoodsStatus(ctx context.Context, data entity.GoodsStatus) (int, error)
	GetAllGoodsStatus(ctx context.Context) ([]entity.GoodsStatus, error)
	GetGoodsStatusByID(ctx context.Context, id int) (entity.GoodsStatus, error)
	UpdateGoodsStatusByID(ctx context.Context, data entity.GoodsStatus) (int64, error)
	DeleteGoodsStatusByID(ctx context.Context, id int) (int64, error)
}
