package service

import (
	"context"
	"database/sql"
	"dc-transaction-service/app/model/entity"
	"dc-transaction-service/app/model/request"
	"dc-transaction-service/app/repository"
	"dc-transaction-service/config"
	"errors"
	"fmt"
	"net/http"
)

type goodsServiceImpl struct {
	cfg       *config.ApplicationConfig
	goodsRepo repository.GoodsRepositoryInterface
}

func NewGoodsServiceImpl(
	cfg *config.ApplicationConfig,
	goodsRepo repository.GoodsRepositoryInterface) GoodsServiceInterface {

	return &goodsServiceImpl{
		cfg:       cfg,
		goodsRepo: goodsRepo,
	}
}

func (m *goodsServiceImpl) GoodsCreate(ctx context.Context, req request.GoodsRequest) (entity.Goods, int, error) {
	data := entity.Goods{
		DateProcess:    req.DateProcess,
		TransactionUID: req.TransactionUID,
		CategoryID:     req.CategoryID,
		Status:         req.Status,
		ExpiredDate:    req.ExpiredDate,
	}

	uid, err := m.goodsRepo.InsertGoods(ctx, data)
	if err != nil {
		return entity.Goods{}, http.StatusInternalServerError, fmt.Errorf("uc.GoodsCreate (InsertGoods) : %v", err)
	}

	data.UID = uid
	return data, http.StatusCreated, nil
}

func (m *goodsServiceImpl) GoodsGetAll(ctx context.Context) ([]entity.Goods, int, error) {
	data, err := m.goodsRepo.GetAllGoods(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.GoodsGetAll (GetAllGoods) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *goodsServiceImpl) GoodsGetByID(ctx context.Context, uid string) (entity.Goods, int, error) {
	data, err := m.goodsRepo.GetGoodsByUID(ctx, uid)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.Goods{}, http.StatusNotFound, fmt.Errorf("uc.GoodsGetByID : goods not found")
	case err != nil:
		return entity.Goods{}, http.StatusInternalServerError, fmt.Errorf("uc.GoodsGetByID (GetGoodsByUID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *goodsServiceImpl) GoodsUpdate(ctx context.Context, uid string, req request.GoodsRequest) (int, error) {
	resultRows, err := m.goodsRepo.UpdateGoodsByUID(ctx, entity.Goods{
		UID:            uid,
		DateProcess:    req.DateProcess,
		TransactionUID: req.TransactionUID,
		CategoryID:     req.CategoryID,
		Status:         req.Status,
		ExpiredDate:    req.ExpiredDate,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.GoodsUpdate (UpdateGoodsByUID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.GoodsUpdate : goods not found")
	}

	return http.StatusOK, nil
}

func (m *goodsServiceImpl) GoodsDelete(ctx context.Context, uid string) (int, error) {
	resultRows, err := m.goodsRepo.DeleteGoodsByUID(ctx, uid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.GoodsDelete (DeleteGoodsByUID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.GoodsDelete : goods not found")
	}

	return http.StatusOK, nil
}
