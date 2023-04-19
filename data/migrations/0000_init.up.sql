CREATE TABLE IF NOT EXISTS "user_account"
(
    id text NOT NULL,
    username text,
    password text,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS "user_account" OWNER to postgres;
