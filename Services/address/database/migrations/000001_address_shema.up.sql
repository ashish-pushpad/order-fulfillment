CREATE TABLE address (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT NOT NULL,

    full_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,

    house_no TEXT NOT NULL,
    apartment TEXT,
    colony TEXT NOT NULL,

    city TEXT NOT NULL,
    state TEXT NOT NULL,

    pin_code VARCHAR(10) NOT NULL,

    is_default BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_address_user
ON address(user_id);