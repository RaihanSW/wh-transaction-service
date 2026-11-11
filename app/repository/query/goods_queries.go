package query

const InsertGoods = `
INSERT INTO "Transaction".goods
	(dateprocess, transaction_uid, category_id, status, expired_date, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $6)
RETURNING uid;
`

const GetAllGoods = `
SELECT uid, dateprocess, transaction_uid, category_id, status, expired_date,
	created_at, created_by, updated_at, updated_by
FROM "Transaction".goods
ORDER BY created_at;
`

const GetGoodsByUID = `
SELECT uid, dateprocess, transaction_uid, category_id, status, expired_date,
	created_at, created_by, updated_at, updated_by
FROM "Transaction".goods
WHERE uid = $1;
`

const UpdateGoodsByUID = `
UPDATE "Transaction".goods
SET dateprocess = $1,
	transaction_uid = $2,
	category_id = $3,
	status = $4,
	expired_date = $5,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $6
WHERE uid = $7;
`

const DeleteGoodsByUID = `
DELETE FROM "Transaction".goods
WHERE uid = $1;
`
