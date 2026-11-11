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

type transactionStatusServiceImpl struct {
	cfg                   *config.ApplicationConfig
	transactionStatusRepo repository.TransactionStatusRepositoryInterface
}

func NewTransactionStatusServiceImpl(
	cfg *config.ApplicationConfig,
	transactionStatusRepo repository.TransactionStatusRepositoryInterface) TransactionStatusServiceInterface {

	return &transactionStatusServiceImpl{
		cfg:                   cfg,
		transactionStatusRepo: transactionStatusRepo,
	}
}

func (m *transactionStatusServiceImpl) TransactionStatusCreate(ctx context.Context, req request.TransactionStatusRequest) (entity.TransactionStatus, int, error) {
	data := entity.TransactionStatus{Name: req.Name}

	id, err := m.transactionStatusRepo.InsertTransactionStatus(ctx, data)
	if err != nil {
		return entity.TransactionStatus{}, http.StatusInternalServerError, fmt.Errorf("uc.TransactionStatusCreate (InsertTransactionStatus) : %v", err)
	}

	data.ID = id
	return data, http.StatusCreated, nil
}

func (m *transactionStatusServiceImpl) TransactionStatusGetAll(ctx context.Context) ([]entity.TransactionStatus, int, error) {
	data, err := m.transactionStatusRepo.GetAllTransactionStatus(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.TransactionStatusGetAll (GetAllTransactionStatus) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *transactionStatusServiceImpl) TransactionStatusGetByID(ctx context.Context, id int) (entity.TransactionStatus, int, error) {
	data, err := m.transactionStatusRepo.GetTransactionStatusByID(ctx, id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.TransactionStatus{}, http.StatusNotFound, fmt.Errorf("uc.TransactionStatusGetByID : transaction status not found")
	case err != nil:
		return entity.TransactionStatus{}, http.StatusInternalServerError, fmt.Errorf("uc.TransactionStatusGetByID (GetTransactionStatusByID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *transactionStatusServiceImpl) TransactionStatusUpdate(ctx context.Context, id int, req request.TransactionStatusRequest) (int, error) {
	resultRows, err := m.transactionStatusRepo.UpdateTransactionStatusByID(ctx, entity.TransactionStatus{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.TransactionStatusUpdate (UpdateTransactionStatusByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.TransactionStatusUpdate : transaction status not found")
	}

	return http.StatusOK, nil
}

func (m *transactionStatusServiceImpl) TransactionStatusDelete(ctx context.Context, id int) (int, error) {
	resultRows, err := m.transactionStatusRepo.DeleteTransactionStatusByID(ctx, id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.TransactionStatusDelete (DeleteTransactionStatusByID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.TransactionStatusDelete : transaction status not found")
	}

	return http.StatusOK, nil
}
