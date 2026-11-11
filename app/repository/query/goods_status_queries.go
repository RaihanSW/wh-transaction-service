package query

const InsertGoodsStatus = `
INSERT INTO "Transaction".goods_status (name, created_by, updated_by)
VALUES ($1, $2, $2)
RETURNING id;
`

const GetAllGoodsStatus = `
SELECT id, name, created_at, created_by, updated_at, updated_by
FROM "Transaction".goods_status
ORDER BY id;
`

const GetGoodsStatusByID = `
SELECT id, name, created_at, created_by, updated_at, updated_by
FROM "Transaction".goods_status
WHERE id = $1;
`

const UpdateGoodsStatusByID = `
UPDATE "Transaction".goods_status
SET name = $1,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $2
WHERE id = $3;
`

const DeleteGoodsStatusByID = `
DELETE FROM "Transaction".goods_status
WHERE id = $1;
`
