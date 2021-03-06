\c postgres;
DROP DATABASE mygrpc;
CREATE DATABASE mygrpc;
GRANT ALL PRIVILEGES ON DATABASE "mygrpc" to postgres;

\c mygrpc;

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE);
