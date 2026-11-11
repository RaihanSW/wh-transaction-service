package service

import (
	"context"
	"dc-transaction-service/app/model/entity"
	"dc-transaction-service/app/model/request"
)

type GoodsServiceInterface interface {
	GoodsCreate(ctx context.Context, req request.GoodsRequest) (entity.Goods, int, error)
	GoodsGetAll(ctx context.Context) ([]entity.Goods, int, error)
	GoodsGetByID(ctx context.Context, uid string) (entity.Goods, int, error)
	GoodsUpdate(ctx context.Context, uid string, req request.GoodsRequest) (int, error)
	GoodsDelete(ctx context.Context, uid string) (int, error)
}
