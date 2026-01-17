-- name: CreatePoly :one
INSERT INTO polymers (id, entry_id, poldescription, poltype,polsequence,pollength,formulaweight,source,host,number_of_molecules,created_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
)
RETURNING *; 
-- name: GetPolys :many
SELECT polymers.* FROM polymers JOIN entries ON polymers.entry_id = entries.id WHERE entries.rcsb_id=$1;

-- name: CreateNonPoly :one
INSERT INTO non_polymers (id, entry_id, nonpolname, comp_id,nonpoldescription,formula_weight,number_of_molecules,created_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *; 

-- name: GetNonPolys :many
SELECT non_polymers.* FROM non_polymers JOIN entries ON non_polymers.entry_id = entries.id WHERE entries.rcsb_id=$1;
