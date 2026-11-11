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

type goodsRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewGoodsRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) GoodsRepositoryInterface {
	return &goodsRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *goodsRepository) InsertGoods(ctx context.Context, data entity.Goods) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var uid string
	err := r.db.QueryRowContext(ctx, query.InsertGoods,
		data.DateProcess,
		data.TransactionUID,
		data.CategoryID,
		data.Status,
		data.ExpiredDate,
		r.serviceName,
	).Scan(&uid)
	if err != nil {
		return "", fmt.Errorf("repo.InsertGoods (Scan): %w", err)
	}

	return uid, nil
}

func (r *goodsRepository) GetAllGoods(ctx context.Context) ([]entity.Goods, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllGoods)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllGoods (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.Goods
	for rows.Next() {
		var data entity.Goods
		if err := scanGoods(rows, &data); err != nil {
			return nil, fmt.Errorf("repo.GetAllGoods (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllGoods (Rows): %w", err)
	}

	return list, nil
}

func (r *goodsRepository) GetGoodsByUID(ctx context.Context, uid string) (entity.Goods, error) {
	var data entity.Goods

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetGoodsByUID, uid).Scan(
		&data.UID,
		&data.DateProcess,
		&data.TransactionUID,
		&data.CategoryID,
		&data.Status,
		&data.ExpiredDate,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.Goods{}, fmt.Errorf("repo.GetGoodsByUID (Scan): %w", err)
	}

	return data, nil
}

func (r *goodsRepository) UpdateGoodsByUID(ctx context.Context, data entity.Goods) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateGoodsByUID,
		data.DateProcess,
		data.TransactionUID,
		data.CategoryID,
		data.Status,
		data.ExpiredDate,
		r.serviceName,
		data.UID,
	)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateGoodsByUID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateGoodsByUID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *goodsRepository) DeleteGoodsByUID(ctx context.Context, uid string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteGoodsByUID, uid)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteGoodsByUID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteGoodsByUID (RowsAffected): %w", err)
	}

	return ra, nil
}

func scanGoods(rows *sql.Rows, data *entity.Goods) error {
	return rows.Scan(
		&data.UID,
		&data.DateProcess,
		&data.TransactionUID,
		&data.CategoryID,
		&data.Status,
		&data.ExpiredDate,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
}
