package service

// =============================================================================
// CONTOH UNIT TEST — GoodsService (CRUD di "Transaction".goods).
//
// Jalankan:  go test ./app/service/ -run TestGoods -v
// =============================================================================

import (
	"context"
	"database/sql"
	"dc-transaction-service/app/model/entity"
	"dc-transaction-service/app/model/request"
	"dc-transaction-service/app/repository"
	"dc-transaction-service/config"
	"errors"
	"net/http"
	"testing"
)

type stubGoodsRepo struct {
	repository.GoodsRepositoryInterface

	insertUID  string
	insertData entity.Goods
	insertErr  error
	getAllData []entity.Goods
	getAllErr  error
	getData    entity.Goods
	getErr     error
	updateRows int64
	updateErr  error
	deleteUID  string
	deleteRows int64
	deleteErr  error
}

func (s *stubGoodsRepo) InsertGoods(ctx context.Context, data entity.Goods) (string, error) {
	s.insertData = data
	return s.insertUID, s.insertErr
}

func (s *stubGoodsRepo) GetAllGoods(ctx context.Context) ([]entity.Goods, error) {
	return s.getAllData, s.getAllErr
}

func (s *stubGoodsRepo) GetGoodsByUID(ctx context.Context, uid string) (entity.Goods, error) {
	return s.getData, s.getErr
}

func (s *stubGoodsRepo) UpdateGoodsByUID(ctx context.Context, data entity.Goods) (int64, error) {
	return s.updateRows, s.updateErr
}

func (s *stubGoodsRepo) DeleteGoodsByUID(ctx context.Context, uid string) (int64, error) {
	s.deleteUID = uid
	return s.deleteRows, s.deleteErr
}

func newGoodsService(repo *stubGoodsRepo) GoodsServiceInterface {
	return NewGoodsServiceImpl(&config.ApplicationConfig{ServiceName: "test-service"}, repo)
}

// -----------------------------------------------------------------------------

func TestGoodsCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 201, uid dari DB dipasang ke entity", func(t *testing.T) {
		repo := &stubGoodsRepo{insertUID: "FEDCBA0987654321"}
		svc := newGoodsService(repo)

		data, status, err := svc.GoodsCreate(ctx, request.GoodsRequest{
			TransactionUID: "TRX-UID-1",
			CategoryID:     1,
			Status:         1,
		})
		if err != nil || status != http.StatusCreated {
			t.Fatalf("expect 201/nil, got %d/%v", status, err)
		}
		if data.UID != "FEDCBA0987654321" {
			t.Fatalf("expect uid from repo, got %s", data.UID)
		}
		if repo.insertData.TransactionUID != "TRX-UID-1" || repo.insertData.CategoryID != 1 {
			t.Fatalf("entity not mapped correctly: %+v", repo.insertData)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubGoodsRepo{insertErr: errors.New("db down")}
		svc := newGoodsService(repo)

		_, status, err := svc.GoodsCreate(ctx, request.GoodsRequest{TransactionUID: "T1", CategoryID: 1, Status: 1})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestGoodsGetAll(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200 + list", func(t *testing.T) {
		repo := &stubGoodsRepo{getAllData: []entity.Goods{{UID: "G1"}, {UID: "G2"}}}
		svc := newGoodsService(repo)

		data, status, err := svc.GoodsGetAll(ctx)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if len(data) != 2 {
			t.Fatalf("expect 2 items, got %d", len(data))
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubGoodsRepo{getAllErr: errors.New("conn lost")}
		svc := newGoodsService(repo)

		_, status, err := svc.GoodsGetAll(ctx)
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestGoodsGetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("ketemu -> 200 + data", func(t *testing.T) {
		repo := &stubGoodsRepo{getData: entity.Goods{UID: "G1", TransactionUID: "T1"}}
		svc := newGoodsService(repo)

		data, status, err := svc.GoodsGetByID(ctx, "G1")
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if data.UID != "G1" || data.TransactionUID != "T1" {
			t.Fatalf("unexpected data: %+v", data)
		}
	})

	t.Run("ErrNoRows -> 404", func(t *testing.T) {
		repo := &stubGoodsRepo{getErr: sql.ErrNoRows}
		svc := newGoodsService(repo)

		_, status, err := svc.GoodsGetByID(ctx, "NOPE")
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestGoodsUpdate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200", func(t *testing.T) {
		repo := &stubGoodsRepo{updateRows: 1}
		svc := newGoodsService(repo)

		status, err := svc.GoodsUpdate(ctx, "G1", request.GoodsRequest{TransactionUID: "T1", CategoryID: 2, Status: 2})
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubGoodsRepo{updateRows: 0}
		svc := newGoodsService(repo)

		status, err := svc.GoodsUpdate(ctx, "NOPE", request.GoodsRequest{})
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubGoodsRepo{updateErr: errors.New("db down")}
		svc := newGoodsService(repo)

		status, err := svc.GoodsUpdate(ctx, "G1", request.GoodsRequest{})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestGoodsDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200, uid diteruskan ke repo", func(t *testing.T) {
		repo := &stubGoodsRepo{deleteRows: 1}
		svc := newGoodsService(repo)

		status, err := svc.GoodsDelete(ctx, "G1")
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if repo.deleteUID != "G1" {
			t.Fatalf("expect deleted uid G1, got %s", repo.deleteUID)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubGoodsRepo{deleteRows: 0}
		svc := newGoodsService(repo)

		status, err := svc.GoodsDelete(ctx, "NOPE")
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}
