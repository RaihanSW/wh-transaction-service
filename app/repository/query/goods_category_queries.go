package query

const InsertGoodsCategory = `
INSERT INTO "Transaction".goods_category (name, created_by, updated_by)
VALUES ($1, $2, $2)
RETURNING id;
`

const GetAllGoodsCategory = `
SELECT id, name, created_at, created_by, updated_at, updated_by
FROM "Transaction".goods_category
ORDER BY id;
`

const GetGoodsCategoryByID = `
SELECT id, name, created_at, created_by, updated_at, updated_by
FROM "Transaction".goods_category
WHERE id = $1;
`

const UpdateGoodsCategoryByID = `
UPDATE "Transaction".goods_category
SET name = $1,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $2
WHERE id = $3;
`

const DeleteGoodsCategoryByID = `
DELETE FROM "Transaction".goods_category
WHERE id = $1;
`
