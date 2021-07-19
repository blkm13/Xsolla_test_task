-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS items (
    id SERIAL,
    SKU VARCHAR NOT NULL,
    name VARCHAR,
    type VARCHAR,
    price varchar,
    PRIMARY key (id)
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
