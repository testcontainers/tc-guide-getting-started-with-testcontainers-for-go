#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS customers (id serial, name varchar(255), email varchar(255));
    INSERT INTO customers(name, email) VALUES ('John', 'john@gmail.com');
EOSQL
