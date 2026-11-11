package repository

import (
	"context"
	"dc-transaction-service/app/model/entity"
)

type GoodsCategoryRepositoryInterface interface {
	InsertGoodsCategory(ctx context.Context, data entity.GoodsCategory) (int, error)
	GetAllGoodsCategory(ctx context.Context) ([]entity.GoodsCategory, error)
	GetGoodsCategoryByID(ctx context.Context, id int) (entity.GoodsCategory, error)
	UpdateGoodsCategoryByID(ctx context.Context, data entity.GoodsCategory) (int64, error)
	DeleteGoodsCategoryByID(ctx context.Context, id int) (int64, error)
}
