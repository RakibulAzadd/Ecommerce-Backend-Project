CREATE TABLE products (
    id            BIGSERIAL PRIMARY KEY,
    title         VARCHAR(100) NOT NULL,
    description   TEXT ,
    price         DOUBLE PRECISION NOT NULL,
    img_url       TEXT,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE  DEFAULT CURRENT_TIMESTAMP
);
