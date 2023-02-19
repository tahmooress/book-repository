BEGIN;

CREATE TABLE users
(
    id         varchar(26) PRIMARY KEY                                  NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULl,
    created_at TIMESTAMP(6) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
    CONSTRAINT name_uq UNIQUE (name),
    CONSTRAINT email_uq UNIQUE (email)
);

COMMIT;