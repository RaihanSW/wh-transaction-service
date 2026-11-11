package query

const InsertTransaction = `
INSERT INTO "Transaction".transaction
	(code, dateprocess, ttl_weight, status, end_date, ttl_volume, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $7)
RETURNING uid;
`

const GetAllTransactions = `
SELECT uid, code, dateprocess, ttl_weight, status, end_date, ttl_volume,
	created_at, created_by, updated_at, updated_by
FROM "Transaction".transaction
ORDER BY created_at;
`

const GetTransactionByUID = `
SELECT uid, code, dateprocess, ttl_weight, status, end_date, ttl_volume,
	created_at, created_by, updated_at, updated_by
FROM "Transaction".transaction
WHERE uid = $1;
`

const UpdateTransactionByUID = `
UPDATE "Transaction".transaction
SET code = $1,
	dateprocess = $2,
	ttl_weight = $3,
	status = $4,
	end_date = $5,
	ttl_volume = $6,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $7
WHERE uid = $8;
`

const DeleteTransactionByUID = `
DELETE FROM "Transaction".transaction
WHERE uid = $1;
`
