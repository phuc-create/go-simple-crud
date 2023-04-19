CREATE TABLE IF NOT EXISTS "user_account"
(
    id text NOT NULL,
    username text,
    password text,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS "user_account" OWNER to postgres;
