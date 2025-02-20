CREATE TABLE transaction_items (
    id char(36) NOT NULL,
    product_id char(36) NOT NULL,
    price_per_unit decimal(16,2) NOT NULL,
    qty int NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT transaction_items_pk PRIMARY KEY (id),
    CONSTRAINT transaction_items_products_fk FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE INDEX on transaction_items (product_id);
