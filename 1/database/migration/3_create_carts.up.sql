CREATE TABLE carts (
    id char(36) NOT NULL,
    user_id char(36) NOT NULL,
    product_id char(36) NOT NULL,
    qty int NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT carts_pk PRIMARY KEY (id),
    CONSTRAINT carts_users_fk FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT carts_products_fk FOREIGN KEY (product_id) REFERENCES products (id),
    CONSTRAINT carts_uk_1 UNIQUE (user_id, product_id)
);

CREATE INDEX on carts (user_id);
CREATE INDEX on carts (product_id);
