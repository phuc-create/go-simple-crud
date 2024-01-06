CREATE TABLE IF NOT EXISTS public.users
(
  id TEXT NOT NULL,
  username TEXT,
  password TEXT,
  email TEXT,
  address TEXT,
  phone TEXT,
  country TEXT,
  company TEXT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  CONSTRAINT "Users_pkey" PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.users OWNER to postgres;
