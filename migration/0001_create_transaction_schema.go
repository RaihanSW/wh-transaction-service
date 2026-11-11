package migration

import "database/sql"

func init() {
	Register(Migration{
		Version: "202611110001",
		Name:    "create_transaction_schema_tables",
		DBKey:   "warehouse_db",
		Up: func(tx *sql.Tx) error {
			stmts := []string{
				`CREATE SCHEMA IF NOT EXISTS "Transaction"`,
				`CREATE TABLE IF NOT EXISTS "Transaction".transaction_status (
					id         INT PRIMARY KEY,
					name       VARCHAR(32) NOT NULL,
					created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					created_by VARCHAR(64),
					updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_by VARCHAR(64)
				)`,
				`CREATE TABLE IF NOT EXISTS "Transaction".goods_category (
					id         INT PRIMARY KEY,
					name       VARCHAR(32) NOT NULL,
					created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					created_by VARCHAR(64),
					updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_by VARCHAR(64)
				)`,
				`CREATE TABLE IF NOT EXISTS "Transaction".goods_status (
					id         INT PRIMARY KEY,
					name       VARCHAR(32) NOT NULL,
					created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					created_by VARCHAR(64),
					updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_by VARCHAR(64)
				)`,
				`CREATE TABLE IF NOT EXISTS "Transaction".transaction (
					uid         VARCHAR(16) PRIMARY KEY DEFAULT upper(substr(md5(random()::text), 1, 16)),
					code        VARCHAR(16) NOT NULL,
					dateprocess TIMESTAMP,
					ttl_weight  INT,
					status      INT         NOT NULL REFERENCES "Transaction".transaction_status (id),
					end_date    TIMESTAMP,
					ttl_volume  INT,
					created_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					created_by  VARCHAR(64),
					updated_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_by  VARCHAR(64)
				)`,
				`CREATE TABLE IF NOT EXISTS "Transaction".goods (
					uid             VARCHAR(16) PRIMARY KEY DEFAULT upper(substr(md5(random()::text), 1, 16)),
					dateprocess     TIMESTAMP,
					transaction_uid VARCHAR(16) NOT NULL REFERENCES "Transaction".transaction (uid),
					category_id     INT         NOT NULL REFERENCES "Transaction".goods_category (id),
					status          INT         NOT NULL REFERENCES "Transaction".goods_status (id),
					expired_date    TIMESTAMP,
					created_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					created_by      VARCHAR(64),
					updated_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_by      VARCHAR(64)
				)`,
			}
			for _, stmt := range stmts {
				if _, err := tx.Exec(stmt); err != nil {
					return err
				}
			}
			return nil
		},
		Down: func(tx *sql.Tx) error {
			stmts := []string{
				`DROP TABLE IF EXISTS "Transaction".goods`,
				`DROP TABLE IF EXISTS "Transaction".transaction`,
				`DROP TABLE IF EXISTS "Transaction".goods_status`,
				`DROP TABLE IF EXISTS "Transaction".goods_category`,
				`DROP TABLE IF EXISTS "Transaction".transaction_status`,
				`DROP SCHEMA IF EXISTS "Transaction"`,
			}
			for _, stmt := range stmts {
				if _, err := tx.Exec(stmt); err != nil {
					return err
				}
			}
			return nil
		},
	})
}
