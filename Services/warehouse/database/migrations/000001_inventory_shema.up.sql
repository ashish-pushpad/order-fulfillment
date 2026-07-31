CREATE TABLE inventories (
    id BIGSERIAL PRIMARY KEY,

    warehouse_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,

    quantity INTEGER NOT NULL
        CHECK (quantity >= 0),

    reserved_quantity INTEGER NOT NULL DEFAULT 0
        CHECK (reserved_quantity >= 0),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT uq_inventory
        UNIQUE (warehouse_id, product_id)
);