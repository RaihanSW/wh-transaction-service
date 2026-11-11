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

type goodsStatusRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewGoodsStatusRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) GoodsStatusRepositoryInterface {
	return &goodsStatusRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *goodsStatusRepository) InsertGoodsStatus(ctx context.Context, data entity.GoodsStatus) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var id int
	err := r.db.QueryRowContext(ctx, query.InsertGoodsStatus, data.Name, r.serviceName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo.InsertGoodsStatus (Scan): %w", err)
	}

	return id, nil
}

func (r *goodsStatusRepository) GetAllGoodsStatus(ctx context.Context) ([]entity.GoodsStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllGoodsStatus)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllGoodsStatus (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.GoodsStatus
	for rows.Next() {
		var data entity.GoodsStatus
		if err := rows.Scan(&data.ID, &data.Name, &data.CreatedAt, &data.CreatedBy, &data.UpdatedAt, &data.UpdatedBy); err != nil {
			return nil, fmt.Errorf("repo.GetAllGoodsStatus (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllGoodsStatus (Rows): %w", err)
	}

	return list, nil
}

func (r *goodsStatusRepository) GetGoodsStatusByID(ctx context.Context, id int) (entity.GoodsStatus, error) {
	var data entity.GoodsStatus

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetGoodsStatusByID, id).Scan(
		&data.ID,
		&data.Name,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.GoodsStatus{}, fmt.Errorf("repo.GetGoodsStatusByID (Scan): %w", err)
	}

	return data, nil
}

func (r *goodsStatusRepository) UpdateGoodsStatusByID(ctx context.Context, data entity.GoodsStatus) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateGoodsStatusByID, data.Name, r.serviceName, data.ID)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateGoodsStatusByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateGoodsStatusByID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *goodsStatusRepository) DeleteGoodsStatusByID(ctx context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteGoodsStatusByID, id)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteGoodsStatusByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteGoodsStatusByID (RowsAffected): %w", err)
	}

	return ra, nil
}
