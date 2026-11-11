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

type transactionStatusRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewTransactionStatusRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) TransactionStatusRepositoryInterface {
	return &transactionStatusRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *transactionStatusRepository) InsertTransactionStatus(ctx context.Context, data entity.TransactionStatus) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var id int
	err := r.db.QueryRowContext(ctx, query.InsertTransactionStatus, data.Name, r.serviceName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo.InsertTransactionStatus (Scan): %w", err)
	}

	return id, nil
}

func (r *transactionStatusRepository) GetAllTransactionStatus(ctx context.Context) ([]entity.TransactionStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllTransactionStatus)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllTransactionStatus (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.TransactionStatus
	for rows.Next() {
		var data entity.TransactionStatus
		if err := rows.Scan(&data.ID, &data.Name, &data.CreatedAt, &data.CreatedBy, &data.UpdatedAt, &data.UpdatedBy); err != nil {
			return nil, fmt.Errorf("repo.GetAllTransactionStatus (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllTransactionStatus (Rows): %w", err)
	}

	return list, nil
}

func (r *transactionStatusRepository) GetTransactionStatusByID(ctx context.Context, id int) (entity.TransactionStatus, error) {
	var data entity.TransactionStatus

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetTransactionStatusByID, id).Scan(
		&data.ID,
		&data.Name,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.TransactionStatus{}, fmt.Errorf("repo.GetTransactionStatusByID (Scan): %w", err)
	}

	return data, nil
}

func (r *transactionStatusRepository) UpdateTransactionStatusByID(ctx context.Context, data entity.TransactionStatus) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateTransactionStatusByID, data.Name, r.serviceName, data.ID)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateTransactionStatusByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateTransactionStatusByID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *transactionStatusRepository) DeleteTransactionStatusByID(ctx context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteTransactionStatusByID, id)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteTransactionStatusByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteTransactionStatusByID (RowsAffected): %w", err)
	}

	return ra, nil
}
