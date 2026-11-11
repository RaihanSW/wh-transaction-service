package repository

import (
	"context"
	"database/sql"
	"dc-transaction-service/app/model/entity"
	"dc-transaction-service/app/repository/query"
	"dc-transaction-service/config"
	"fmt"
	"time"
)

type transactionRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewTransactionRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) TransactionRepositoryInterface {
	return &transactionRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *transactionRepository) InsertTransaction(ctx context.Context, data entity.Transaction) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var uid string
	err := r.db.QueryRowContext(ctx, query.InsertTransaction,
		data.Code,
		data.DateProcess,
		data.TtlWeight,
		data.Status,
		data.EndDate,
		data.TtlVolume,
		r.serviceName,
	).Scan(&uid)
	if err != nil {
		return "", fmt.Errorf("repo.InsertTransaction (Scan): %w", err)
	}

	return uid, nil
}

func (r *transactionRepository) GetAllTransactions(ctx context.Context) ([]entity.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllTransactions)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllTransactions (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.Transaction
	for rows.Next() {
		var data entity.Transaction
		if err := scanTransaction(rows, &data); err != nil {
			return nil, fmt.Errorf("repo.GetAllTransactions (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllTransactions (Rows): %w", err)
	}

	return list, nil
}

func (r *transactionRepository) GetTransactionByUID(ctx context.Context, uid string) (entity.Transaction, error) {
	var data entity.Transaction

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetTransactionByUID, uid).Scan(
		&data.UID,
		&data.Code,
		&data.DateProcess,
		&data.TtlWeight,
		&data.Status,
		&data.EndDate,
		&data.TtlVolume,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("repo.GetTransactionByUID (Scan): %w", err)
	}

	return data, nil
}

func (r *transactionRepository) UpdateTransactionByUID(ctx context.Context, data entity.Transaction) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateTransactionByUID,
		data.Code,
		data.DateProcess,
		data.TtlWeight,
		data.Status,
		data.EndDate,
		data.TtlVolume,
		r.serviceName,
		data.UID,
	)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateTransactionByUID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateTransactionByUID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *transactionRepository) DeleteTransactionByUID(ctx context.Context, uid string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteTransactionByUID, uid)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteTransactionByUID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteTransactionByUID (RowsAffected): %w", err)
	}

	return ra, nil
}

func scanTransaction(rows *sql.Rows, data *entity.Transaction) error {
	return rows.Scan(
		&data.UID,
		&data.Code,
		&data.DateProcess,
		&data.TtlWeight,
		&data.Status,
		&data.EndDate,
		&data.TtlVolume,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
}
