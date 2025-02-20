CREATE TABLE transactions (
    id char(36) NOT NULL,
    user_id char(36) NOT NULL,
    total_price decimal(16,2) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT transactions_pk PRIMARY KEY (id),
    CONSTRAINT transactions_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX ON transactions (user_id);
