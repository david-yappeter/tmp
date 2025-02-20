CREATE TABLE users (
    id char(36) NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
);


CREATE INDEX ON users (email);
