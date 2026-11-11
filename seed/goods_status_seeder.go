package seed

import (
	"context"
	"database/sql"
)

func init() {
	Register(Seeder{
		Name:  "goods_status",
		DBKey: "warehouse_db",
		Run: func(ctx context.Context, db *sql.DB) error {
			// TODO: sesuaikan nilai lookup
			statuses := []struct {
				id   int
				name string
			}{
				{1, "Available"},
				{2, "Reserved"},
				{3, "Expired"},
			}

			for _, s := range statuses {
				_, err := db.ExecContext(ctx, `
					INSERT INTO "Transaction".goods_status (id, name, created_by, updated_by)
					VALUES ($1, $2, $3, $3)
					ON CONFLICT (id) DO NOTHING
				`, s.id, s.name, "seeder")
				if err != nil {
					return err
				}
			}

			// id eksplisit di atas tidak menggerakkan identity sequence —
			// selaraskan agar insert berikutnya tidak bentrok PK
			_, err := db.ExecContext(ctx, `
				SELECT setval(
					pg_get_serial_sequence('"Transaction".goods_status', 'id'),
					(SELECT MAX(id) FROM "Transaction".goods_status)
				)
			`)
			return err
		},
	})
}
