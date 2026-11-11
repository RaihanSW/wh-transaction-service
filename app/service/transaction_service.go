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

type transactionServiceImpl struct {
	cfg             *config.ApplicationConfig
	transactionRepo repository.TransactionRepositoryInterface
}

func NewTransactionServiceImpl(
	cfg *config.ApplicationConfig,
	transactionRepo repository.TransactionRepositoryInterface) TransactionServiceInterface {

	return &transactionServiceImpl{
		cfg:             cfg,
		transactionRepo: transactionRepo,
	}
}

func (m *transactionServiceImpl) TransactionCreate(ctx context.Context, req request.TransactionRequest) (entity.Transaction, int, error) {
	data := entity.Transaction{
		Code:        req.Code,
		DateProcess: req.DateProcess,
		TtlWeight:   req.TtlWeight,
		Status:      req.Status,
		EndDate:     req.EndDate,
		TtlVolume:   req.TtlVolume,
	}

	uid, err := m.transactionRepo.InsertTransaction(ctx, data)
	if err != nil {
		return entity.Transaction{}, http.StatusInternalServerError, fmt.Errorf("uc.TransactionCreate (InsertTransaction) : %v", err)
	}

	data.UID = uid
	return data, http.StatusCreated, nil
}

func (m *transactionServiceImpl) TransactionGetAll(ctx context.Context) ([]entity.Transaction, int, error) {
	data, err := m.transactionRepo.GetAllTransactions(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("uc.TransactionGetAll (GetAllTransactions) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *transactionServiceImpl) TransactionGetByID(ctx context.Context, uid string) (entity.Transaction, int, error) {
	data, err := m.transactionRepo.GetTransactionByUID(ctx, uid)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entity.Transaction{}, http.StatusNotFound, fmt.Errorf("uc.TransactionGetByID : transaction not found")
	case err != nil:
		return entity.Transaction{}, http.StatusInternalServerError, fmt.Errorf("uc.TransactionGetByID (GetTransactionByUID) : %v", err)
	}

	return data, http.StatusOK, nil
}

func (m *transactionServiceImpl) TransactionUpdate(ctx context.Context, uid string, req request.TransactionRequest) (int, error) {
	resultRows, err := m.transactionRepo.UpdateTransactionByUID(ctx, entity.Transaction{
		UID:         uid,
		Code:        req.Code,
		DateProcess: req.DateProcess,
		TtlWeight:   req.TtlWeight,
		Status:      req.Status,
		EndDate:     req.EndDate,
		TtlVolume:   req.TtlVolume,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.TransactionUpdate (UpdateTransactionByUID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.TransactionUpdate : transaction not found")
	}

	return http.StatusOK, nil
}

func (m *transactionServiceImpl) TransactionDelete(ctx context.Context, uid string) (int, error) {
	resultRows, err := m.transactionRepo.DeleteTransactionByUID(ctx, uid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.TransactionDelete (DeleteTransactionByUID) : %v", err)
	}
	if resultRows < 1 {
		return http.StatusNotFound, fmt.Errorf("uc.TransactionDelete : transaction not found")
	}

	return http.StatusOK, nil
}
