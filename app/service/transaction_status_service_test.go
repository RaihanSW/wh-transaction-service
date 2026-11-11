package service

// =============================================================================
// CONTOH UNIT TEST — TransactionStatusService (lookup id+name).
//
// Jalankan:  go test ./app/service/ -run TestTransactionStatus -v
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

type stubTransactionStatusRepo struct {
	repository.TransactionStatusRepositoryInterface

	insertID   int
	insertData entity.TransactionStatus
	insertErr  error
	getAllData []entity.TransactionStatus
	getAllErr  error
	getData    entity.TransactionStatus
	getErr     error
	updateRows int64
	updateErr  error
	deleteID   int
	deleteRows int64
	deleteErr  error
}

func (s *stubTransactionStatusRepo) InsertTransactionStatus(ctx context.Context, data entity.TransactionStatus) (int, error) {
	s.insertData = data
	return s.insertID, s.insertErr
}

func (s *stubTransactionStatusRepo) GetAllTransactionStatus(ctx context.Context) ([]entity.TransactionStatus, error) {
	return s.getAllData, s.getAllErr
}

func (s *stubTransactionStatusRepo) GetTransactionStatusByID(ctx context.Context, id int) (entity.TransactionStatus, error) {
	return s.getData, s.getErr
}

func (s *stubTransactionStatusRepo) UpdateTransactionStatusByID(ctx context.Context, data entity.TransactionStatus) (int64, error) {
	return s.updateRows, s.updateErr
}

func (s *stubTransactionStatusRepo) DeleteTransactionStatusByID(ctx context.Context, id int) (int64, error) {
	s.deleteID = id
	return s.deleteRows, s.deleteErr
}

func newTransactionStatusService(repo *stubTransactionStatusRepo) TransactionStatusServiceInterface {
	return NewTransactionStatusServiceImpl(&config.ApplicationConfig{ServiceName: "test-service"}, repo)
}

// -----------------------------------------------------------------------------

func TestTransactionStatusCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 201, id dari DB dipasang ke entity", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{insertID: 7}
		svc := newTransactionStatusService(repo)

		data, status, err := svc.TransactionStatusCreate(ctx, request.TransactionStatusRequest{Name: "Done"})
		if err != nil || status != http.StatusCreated {
			t.Fatalf("expect 201/nil, got %d/%v", status, err)
		}
		if data.ID != 7 || data.Name != "Done" {
			t.Fatalf("unexpected data: %+v", data)
		}
		if repo.insertData.Name != "Done" {
			t.Fatalf("entity not mapped correctly: %+v", repo.insertData)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{insertErr: errors.New("db down")}
		svc := newTransactionStatusService(repo)

		_, status, err := svc.TransactionStatusCreate(ctx, request.TransactionStatusRequest{Name: "X"})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestTransactionStatusGetAll(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200 + list", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{getAllData: []entity.TransactionStatus{{ID: 1}, {ID: 2}}}
		svc := newTransactionStatusService(repo)

		data, status, err := svc.TransactionStatusGetAll(ctx)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if len(data) != 2 {
			t.Fatalf("expect 2 items, got %d", len(data))
		}
	})
}

func TestTransactionStatusGetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("ketemu -> 200 + data", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{getData: entity.TransactionStatus{ID: 1, Name: "Draft"}}
		svc := newTransactionStatusService(repo)

		data, status, err := svc.TransactionStatusGetByID(ctx, 1)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if data.ID != 1 || data.Name != "Draft" {
			t.Fatalf("unexpected data: %+v", data)
		}
	})

	t.Run("ErrNoRows -> 404", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{getErr: sql.ErrNoRows}
		svc := newTransactionStatusService(repo)

		_, status, err := svc.TransactionStatusGetByID(ctx, 99)
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestTransactionStatusUpdate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{updateRows: 1}
		svc := newTransactionStatusService(repo)

		status, err := svc.TransactionStatusUpdate(ctx, 1, request.TransactionStatusRequest{Name: "Renamed"})
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{updateRows: 0}
		svc := newTransactionStatusService(repo)

		status, err := svc.TransactionStatusUpdate(ctx, 99, request.TransactionStatusRequest{Name: "X"})
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestTransactionStatusDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200, id diteruskan ke repo", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{deleteRows: 1}
		svc := newTransactionStatusService(repo)

		status, err := svc.TransactionStatusDelete(ctx, 1)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if repo.deleteID != 1 {
			t.Fatalf("expect deleted id 1, got %d", repo.deleteID)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubTransactionStatusRepo{deleteRows: 0}
		svc := newTransactionStatusService(repo)

		status, err := svc.TransactionStatusDelete(ctx, 99)
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}
