package service

import (
	"context"
	"dc-transaction-service/app/model/entity"
	"dc-transaction-service/app/model/request"
)

type GoodsStatusServiceInterface interface {
	GoodsStatusCreate(ctx context.Context, req request.GoodsStatusRequest) (entity.GoodsStatus, int, error)
	GoodsStatusGetAll(ctx context.Context) ([]entity.GoodsStatus, int, error)
	GoodsStatusGetByID(ctx context.Context, id int) (entity.GoodsStatus, int, error)
	GoodsStatusUpdate(ctx context.Context, id int, req request.GoodsStatusRequest) (int, error)
	GoodsStatusDelete(ctx context.Context, id int) (int, error)
}
