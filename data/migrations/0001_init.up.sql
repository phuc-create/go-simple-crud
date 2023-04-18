CREATE TABLE public."user"
(
    id text NOT NULL,
    username text,
    password text,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public."user"
    OWNER to postgres;