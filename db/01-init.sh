#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    CREATE TABLE IF NOT EXISTS public.book
(
    id integer NOT NULL DEFAULT nextval('book_id_seq'::regclass),
    name text COLLATE pg_catalog."default",
    book_type text COLLATE pg_catalog."default",
    price double precision,
    owner_name text COLLATE pg_catalog."default",
    CONSTRAINT book_pkey PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS public.orders
(
    id bigint NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    code text COLLATE pg_catalog."default",
    price text COLLATE pg_catalog."default",
    CONSTRAINT orders_pkey PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    age integer,
    first_name text COLLATE pg_catalog."default",
    last_name text COLLATE pg_catalog."default",
    address1 text COLLATE pg_catalog."default",
    address2 text COLLATE pg_catalog."default",
    username text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email)
);

  COMMIT;
EOSQL