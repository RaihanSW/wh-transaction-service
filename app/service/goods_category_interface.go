package service

import (
	"context"
	"dc-transaction-service/app/model/entity"
	"dc-transaction-service/app/model/request"
)

type GoodsCategoryServiceInterface interface {
	GoodsCategoryCreate(ctx context.Context, req request.GoodsCategoryRequest) (entity.GoodsCategory, int, error)
	GoodsCategoryGetAll(ctx context.Context) ([]entity.GoodsCategory, int, error)
	GoodsCategoryGetByID(ctx context.Context, id int) (entity.GoodsCategory, int, error)
	GoodsCategoryUpdate(ctx context.Context, id int, req request.GoodsCategoryRequest) (int, error)
	GoodsCategoryDelete(ctx context.Context, id int) (int, error)
}
