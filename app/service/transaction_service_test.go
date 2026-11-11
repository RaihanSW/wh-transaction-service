package service

// =============================================================================
// CONTOH UNIT TEST — TransactionService (CRUD di "Transaction".transaction).
//
// Pola: stub embed interface (nil), implement hanya method yang dipakai,
// tangkap argumen untuk assert.
//
// Jalankan:  go test ./app/service/ -run TestTransaction -v
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

type stubTransactionRepo struct {
	repository.TransactionRepositoryInterface

	insertUID  string
	insertData entity.Transaction
	insertErr  error
	getAllData []entity.Transaction
	getAllErr  error
	getData    entity.Transaction
	getErr     error
	updateRows int64
	updateErr  error
	deleteUID  string
	deleteRows int64
	deleteErr  error
}

func (s *stubTransactionRepo) InsertTransaction(ctx context.Context, data entity.Transaction) (string, error) {
	s.insertData = data
	return s.insertUID, s.insertErr
}

func (s *stubTransactionRepo) GetAllTransactions(ctx context.Context) ([]entity.Transaction, error) {
	return s.getAllData, s.getAllErr
}

func (s *stubTransactionRepo) GetTransactionByUID(ctx context.Context, uid string) (entity.Transaction, error) {
	return s.getData, s.getErr
}

func (s *stubTransactionRepo) UpdateTransactionByUID(ctx context.Context, data entity.Transaction) (int64, error) {
	return s.updateRows, s.updateErr
}

func (s *stubTransactionRepo) DeleteTransactionByUID(ctx context.Context, uid string) (int64, error) {
	s.deleteUID = uid
	return s.deleteRows, s.deleteErr
}

func newTransactionService(repo *stubTransactionRepo) TransactionServiceInterface {
	return NewTransactionServiceImpl(&config.ApplicationConfig{ServiceName: "test-service"}, repo)
}

// -----------------------------------------------------------------------------

func TestTransactionCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 201, uid dari DB dipasang ke entity", func(t *testing.T) {
		repo := &stubTransactionRepo{insertUID: "ABCDEF1234567890"}
		svc := newTransactionService(repo)

		data, status, err := svc.TransactionCreate(ctx, request.TransactionRequest{
			Code:      "TRX-001",
			TtlWeight: 100,
			Status:    1,
			TtlVolume: 50,
		})
		if err != nil || status != http.StatusCreated {
			t.Fatalf("expect 201/nil, got %d/%v", status, err)
		}
		if data.UID != "ABCDEF1234567890" {
			t.Fatalf("expect uid from repo, got %s", data.UID)
		}
		if repo.insertData.Code != "TRX-001" || repo.insertData.TtlWeight != 100 {
			t.Fatalf("entity not mapped correctly: %+v", repo.insertData)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubTransactionRepo{insertErr: errors.New("db down")}
		svc := newTransactionService(repo)

		_, status, err := svc.TransactionCreate(ctx, request.TransactionRequest{Code: "TRX-001", Status: 1})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestTransactionGetAll(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200 + list", func(t *testing.T) {
		repo := &stubTransactionRepo{getAllData: []entity.Transaction{{UID: "U1"}, {UID: "U2"}}}
		svc := newTransactionService(repo)

		data, status, err := svc.TransactionGetAll(ctx)
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if len(data) != 2 {
			t.Fatalf("expect 2 items, got %d", len(data))
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubTransactionRepo{getAllErr: errors.New("conn lost")}
		svc := newTransactionService(repo)

		_, status, err := svc.TransactionGetAll(ctx)
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestTransactionGetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("ketemu -> 200 + data", func(t *testing.T) {
		repo := &stubTransactionRepo{getData: entity.Transaction{UID: "U1", Code: "TRX-001"}}
		svc := newTransactionService(repo)

		data, status, err := svc.TransactionGetByID(ctx, "U1")
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if data.UID != "U1" || data.Code != "TRX-001" {
			t.Fatalf("unexpected data: %+v", data)
		}
	})

	t.Run("ErrNoRows -> 404", func(t *testing.T) {
		repo := &stubTransactionRepo{getErr: sql.ErrNoRows}
		svc := newTransactionService(repo)

		_, status, err := svc.TransactionGetByID(ctx, "NOPE")
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}

func TestTransactionUpdate(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200", func(t *testing.T) {
		repo := &stubTransactionRepo{updateRows: 1}
		svc := newTransactionService(repo)

		status, err := svc.TransactionUpdate(ctx, "U1", request.TransactionRequest{Code: "TRX-009", Status: 2})
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubTransactionRepo{updateRows: 0}
		svc := newTransactionService(repo)

		status, err := svc.TransactionUpdate(ctx, "NOPE", request.TransactionRequest{})
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})

	t.Run("repo error -> 500", func(t *testing.T) {
		repo := &stubTransactionRepo{updateErr: errors.New("db down")}
		svc := newTransactionService(repo)

		status, err := svc.TransactionUpdate(ctx, "U1", request.TransactionRequest{})
		if err == nil || status != http.StatusInternalServerError {
			t.Fatalf("expect 500/error, got %d/%v", status, err)
		}
	})
}

func TestTransactionDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("sukses -> 200, uid diteruskan ke repo", func(t *testing.T) {
		repo := &stubTransactionRepo{deleteRows: 1}
		svc := newTransactionService(repo)

		status, err := svc.TransactionDelete(ctx, "U1")
		if err != nil || status != http.StatusOK {
			t.Fatalf("expect 200/nil, got %d/%v", status, err)
		}
		if repo.deleteUID != "U1" {
			t.Fatalf("expect deleted uid U1, got %s", repo.deleteUID)
		}
	})

	t.Run("rows affected 0 -> 404", func(t *testing.T) {
		repo := &stubTransactionRepo{deleteRows: 0}
		svc := newTransactionService(repo)

		status, err := svc.TransactionDelete(ctx, "NOPE")
		if err == nil || status != http.StatusNotFound {
			t.Fatalf("expect 404/error, got %d/%v", status, err)
		}
	})
}
