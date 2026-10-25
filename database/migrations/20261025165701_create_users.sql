-- +goose Up
SELECT 'up SQL query';

CREATE TABLE users (
    id UUID PRIMARY KEY,

    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    
   bullet_elo INTEGER NOT NULL DEFAULT 800 CHECK (bullet_elo >= 0),
blitz_elo INTEGER NOT NULL DEFAULT 800 CHECK (blitz_elo >= 0),
rapid_elo INTEGER NOT NULL DEFAULT 800 CHECK (rapid_elo >= 0),

    image_id UUID,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    last_login_at TIMESTAMP
);

-- +goose Down
SELECT 'down SQL query';
