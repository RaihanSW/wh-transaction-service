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

type goodsStatusServiceImpl struct {
	cfg             *config.ApplicationConfig
	goodsStatusRepo repository.GoodsStatusRepositoryInterface
}

func NewGoodsStatusServiceImpl(
	cfg *config.ApplicationConfig,
	goodsStatusRepo repository.GoodsStatusRepositoryInterface) GoodsStatusServiceInterface {

	return &goodsStatusServiceImpl{
		cfg:             cfg,
		goodsStatusRepo: goodsStatusRepo,
	}
}

func (m *goodsStatusServiceImpl) GoodsStatusCreate(ctx context.Context, req request.GoodsStatusRequest) (entity.GoodsStatus, int, error) {
	data := entity.GoodsStatus{Name: req.Name}

	id, err := m.goodsStatusRepo.InsertGoodsStatus(ctx, data)
	if err != nil {
		return entity.GoodsStatus{}, http.StatusInternalServerError, fmt.Errorf("uc.GoodsStatusCreate (InsertGoodsStatus) : %v", err)
	}

	data.ID = id
	return data, http.StatusCreated, nil
}

func (m *goodsStatusServiceImpl) GoodsStatusGetAll(ctx context.Context) ([]entity.GoodsStatus, int, error) {
	data, err := m.goodsStatusRepo.GetAllGoodsStatus(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.GoodsStatusGetAll (GetAllGoodsStatus) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *goodsStatusServiceImpl) GoodsStatusGetByID(ctx context.Context, id int) (entity.GoodsStatus, int, error) {
	data, err := m.goodsStatusRepo.GetGoodsStatusByID(ctx, id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.GoodsStatus{}, http.StatusNotFound, fmt.Errorf("uc.GoodsStatusGetByID : goods status not found")
	case err != nil:
		return entity.GoodsStatus{}, http.StatusInternalServerError, fmt.Errorf("uc.GoodsStatusGetByID (GetGoodsStatusByID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *goodsStatusServiceImpl) GoodsStatusUpdate(ctx context.Context, id int, req request.GoodsStatusRequest) (int, error) {
	resultRows, err := m.goodsStatusRepo.UpdateGoodsStatusByID(ctx, entity.GoodsStatus{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.GoodsStatusUpdate (UpdateGoodsStatusByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.GoodsStatusUpdate : goods status not found")
	}

	return http.StatusOK, nil
}

func (m *goodsStatusServiceImpl) GoodsStatusDelete(ctx context.Context, id int) (int, error) {
	resultRows, err := m.goodsStatusRepo.DeleteGoodsStatusByID(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.GoodsStatusDelete (DeleteGoodsStatusByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.GoodsStatusDelete : goods status not found")
	}

	return http.StatusOK, nil
}
