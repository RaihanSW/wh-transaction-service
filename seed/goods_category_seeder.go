package seed

import (
	"context"
	"database/sql"
)

func init() {
	Register(Seeder{
		Name:  "goods_category",
		DBKey: "warehouse_db",
		Run: func(ctx context.Context, db *sql.DB) error {
			// TODO: sesuaikan nilai lookup
			categories := []struct {
				id   int
				name string
			}{
				{1, "General"},
				{2, "Perishable"},
				{3, "Fragile"},
			}

			for _, c := range categories {
				_, err := db.ExecContext(ctx, `
					INSERT INTO "Transaction".goods_category (id, name, created_by, updated_by)
					VALUES ($1, $2, $3, $3)
					ON CONFLICT (id) DO NOTHING
				`, c.id, c.name, "seeder")
				if err != nil {
					return err
				}
			}

			// id eksplisit di atas tidak menggerakkan identity sequence —
			// selaraskan agar insert berikutnya tidak bentrok PK
			_, err := db.ExecContext(ctx, `
				SELECT setval(
					pg_get_serial_sequence('"Transaction".goods_category', 'id'),
					(SELECT MAX(id) FROM "Transaction".goods_category)
				)
			`)
			return err
		},
	})
}
