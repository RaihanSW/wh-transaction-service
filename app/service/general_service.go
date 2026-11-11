package service

import (
	"context"
	"dc-transaction-service/app/repository"
	"dc-transaction-service/config"
	"fmt"
	"net/http"
)

type generalServiceImpl struct {
	cfg         *config.ApplicationConfig
	generalRepo repository.GeneralRepositoryInterface
}

func NewGeneralServiceImpl(
	cfg *config.ApplicationConfig,
	generalRepo repository.GeneralRepositoryInterface) GeneralServiceInterface {

	return &generalServiceImpl{
		cfg:         cfg,
		generalRepo: generalRepo,
	}
}

func (m *generalServiceImpl) PingDB(ctx context.Context) (int, error) {
	err := m.generalRepo.PingDB(ctx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("uc.PingDB (PingDB) : %v", err)
	}
	return http.StatusOK, nil
}
