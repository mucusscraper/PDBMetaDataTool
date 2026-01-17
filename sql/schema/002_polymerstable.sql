-- +goose Up
CREATE TABLE polymers (
    id UUID PRIMARY KEY,
    entry_id UUID NOT NULL,
    poldescription TEXT NOT NULL,
    poltype TEXT NOT NULL,
    polsequence TEXT NOT NULL,
    pollength INTEGER NOT NULL,
    formulaweight REAL NOT NULL,
    source TEXT NOT NULL,
    host TEXT NOT NULL,
    number_of_molecules INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT con_polymers_entry
        FOREIGN KEY(entry_id)
        REFERENCES entries(id)
        ON DELETE CASCADE
);
CREATE TABLE non_polymers (
    id UUID PRIMARY KEY,
    entry_id UUID NOT NULL,
    nonpolname TEXT NOT NULL,
    comp_id TEXT NOT NULL,
    nonpoldescription TEXT NOT NULL,
    formula_weight REAL NOT NULL,
    number_of_molecules INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT con_nonpolymers_entry
        FOREIGN KEY (entry_id)
        REFERENCES entries(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE polymers;
DROP TABLE non_polymers;
