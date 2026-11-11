package repository

import (
	"context"
	"dc-transaction-service/app/model/entity"
)

type GoodsRepositoryInterface interface {
	InsertGoods(ctx context.Context, data entity.Goods) (string, error)
	GetAllGoods(ctx context.Context) ([]entity.Goods, error)
	GetGoodsByUID(ctx context.Context, uid string) (entity.Goods, error)
	UpdateGoodsByUID(ctx context.Context, data entity.Goods) (int64, error)
	DeleteGoodsByUID(ctx context.Context, uid string) (int64, error)
}
