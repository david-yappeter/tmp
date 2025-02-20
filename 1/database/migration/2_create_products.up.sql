CREATE TABLE products (
    id char(36) NOT NULL,
    name varchar(255) NOT NULL,
    price decimal(16,2) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT products_pk PRIMARY KEY (id)
);

