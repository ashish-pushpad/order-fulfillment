CREATE TABLE warehouse (
    id BIGSERIAL PRIMARY KEY,

    name VARCHAR(100) NOT NULL,
    city TEXT NOT NULL,
    address TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_warehouse_city
ON warehouse(city);
