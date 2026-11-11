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

type goodsCategoryRepository struct {
	db             *sql.DB
	dbReplica      *sql.DB
	timeout        time.Duration
	timeoutReplica time.Duration
	serviceName    string
}

func NewGoodsCategoryRepositoryImpl(serviceName string, cfg config.DatabaseConfig, db *sql.DB, cfgReplica config.DatabaseConfig, dbReplica *sql.DB) GoodsCategoryRepositoryInterface {
	return &goodsCategoryRepository{
		db:             db,
		dbReplica:      dbReplica,
		timeout:        cfg.Timeout * time.Second,
		timeoutReplica: cfgReplica.Timeout * time.Second,
		serviceName:    serviceName,
	}
}

func (r *goodsCategoryRepository) InsertGoodsCategory(ctx context.Context, data entity.GoodsCategory) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var id int
	err := r.db.QueryRowContext(ctx, query.InsertGoodsCategory, data.Name, r.serviceName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo.InsertGoodsCategory (Scan): %w", err)
	}

	return id, nil
}

func (r *goodsCategoryRepository) GetAllGoodsCategory(ctx context.Context) ([]entity.GoodsCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	rows, err := r.dbReplica.QueryContext(ctx, query.GetAllGoodsCategory)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllGoodsCategory (QueryContext): %w", err)
	}
	defer rows.Close()

	var list []entity.GoodsCategory
	for rows.Next() {
		var data entity.GoodsCategory
		if err := rows.Scan(&data.ID, &data.Name, &data.CreatedAt, &data.CreatedBy, &data.UpdatedAt, &data.UpdatedBy); err != nil {
			return nil, fmt.Errorf("repo.GetAllGoodsCategory (Scan): %w", err)
		}
		list = append(list, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.GetAllGoodsCategory (Rows): %w", err)
	}

	return list, nil
}

func (r *goodsCategoryRepository) GetGoodsCategoryByID(ctx context.Context, id int) (entity.GoodsCategory, error) {
	var data entity.GoodsCategory

	ctx, cancel := context.WithTimeout(ctx, r.timeoutReplica)
	defer cancel()

	err := r.dbReplica.QueryRowContext(ctx, query.GetGoodsCategoryByID, id).Scan(
		&data.ID,
		&data.Name,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return entity.GoodsCategory{}, fmt.Errorf("repo.GetGoodsCategoryByID (Scan): %w", err)
	}

	return data, nil
}

func (r *goodsCategoryRepository) UpdateGoodsCategoryByID(ctx context.Context, data entity.GoodsCategory) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.UpdateGoodsCategoryByID, data.Name, r.serviceName, data.ID)
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateGoodsCategoryByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.UpdateGoodsCategoryByID (RowsAffected): %w", err)
	}

	return ra, nil
}

func (r *goodsCategoryRepository) DeleteGoodsCategoryByID(ctx context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query.DeleteGoodsCategoryByID, id)
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteGoodsCategoryByID (ExecContext): %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("repo.DeleteGoodsCategoryByID (RowsAffected): %w", err)
	}

	return ra, nil
}
