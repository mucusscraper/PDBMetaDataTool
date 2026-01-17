-- name: CreateEntry :one
INSERT INTO entries (id, created_at, updated_at, rcsb_id,deposit_date,doi,paper_title,method,user_group)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries WHERE rcsb_id=$1;

-- name: GetEntryByUserGroup :many
SELECT * FROM entries WHERE user_group=$1;

-- name: InsertGroup :one
UPDATE entries SET user_group=$1 WHERE rcsb_id=$2 RETURNING *;

-- name: RemoveGroup :one
UPDATE entries SET user_group='' WHERE rcsb_id=$1 RETURNING *;
