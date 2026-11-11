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

type goodsCategoryServiceImpl struct {
	cfg               *config.ApplicationConfig
	goodsCategoryRepo repository.GoodsCategoryRepositoryInterface
}

func NewGoodsCategoryServiceImpl(
	cfg *config.ApplicationConfig,
	goodsCategoryRepo repository.GoodsCategoryRepositoryInterface) GoodsCategoryServiceInterface {

	return &goodsCategoryServiceImpl{
		cfg:               cfg,
		goodsCategoryRepo: goodsCategoryRepo,
	}
}

func (m *goodsCategoryServiceImpl) GoodsCategoryCreate(ctx context.Context, req request.GoodsCategoryRequest) (entity.GoodsCategory, int, error) {
	data := entity.GoodsCategory{Name: req.Name}

	id, err := m.goodsCategoryRepo.InsertGoodsCategory(ctx, data)
	if err != nil {
		return entity.GoodsCategory{}, http.StatusInternalServerError, fmt.Errorf("uc.GoodsCategoryCreate (InsertGoodsCategory) : %v", err)
	}

	data.ID = id
	return data, http.StatusCreated, nil
}

func (m *goodsCategoryServiceImpl) GoodsCategoryGetAll(ctx context.Context) ([]entity.GoodsCategory, int, error) {
	data, err := m.goodsCategoryRepo.GetAllGoodsCategory(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.GoodsCategoryGetAll (GetAllGoodsCategory) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *goodsCategoryServiceImpl) GoodsCategoryGetByID(ctx context.Context, id int) (entity.GoodsCategory, int, error) {
	data, err := m.goodsCategoryRepo.GetGoodsCategoryByID(ctx, id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.GoodsCategory{}, http.StatusNotFound, fmt.Errorf("uc.GoodsCategoryGetByID : goods category not found")
	case err != nil:
		return entity.GoodsCategory{}, http.StatusInternalServerError, fmt.Errorf("uc.GoodsCategoryGetByID (GetGoodsCategoryByID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *goodsCategoryServiceImpl) GoodsCategoryUpdate(ctx context.Context, id int, req request.GoodsCategoryRequest) (int, error) {
	resultRows, err := m.goodsCategoryRepo.UpdateGoodsCategoryByID(ctx, entity.GoodsCategory{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.GoodsCategoryUpdate (UpdateGoodsCategoryByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.GoodsCategoryUpdate : goods category not found")
	}

	return http.StatusOK, nil
}

func (m *goodsCategoryServiceImpl) GoodsCategoryDelete(ctx context.Context, id int) (int, error) {
	resultRows, err := m.goodsCategoryRepo.DeleteGoodsCategoryByID(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.GoodsCategoryDelete (DeleteGoodsCategoryByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.GoodsCategoryDelete : goods category not found")
	}

	return http.StatusOK, nil
}
