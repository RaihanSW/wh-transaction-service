package query

const InsertTransactionStatus = `
INSERT INTO "Transaction".transaction_status (name, created_by, updated_by)
VALUES ($1, $2, $2)
RETURNING id;
`

const GetAllTransactionStatus = `
SELECT id, name, created_at, created_by, updated_at, updated_by
FROM "Transaction".transaction_status
ORDER BY id;
`

const GetTransactionStatusByID = `
SELECT id, name, created_at, created_by, updated_at, updated_by
FROM "Transaction".transaction_status
WHERE id = $1;
`

const UpdateTransactionStatusByID = `
UPDATE "Transaction".transaction_status
SET name = $1,
	updated_at = CURRENT_TIMESTAMP,
	updated_by = $2
WHERE id = $3;
`

const DeleteTransactionStatusByID = `
DELETE FROM "Transaction".transaction_status
WHERE id = $1;
`
