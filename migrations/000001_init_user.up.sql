CREATE TABLE users (
                       id bigserial PRIMARY KEY,
                       username varchar UNIQUE NOT NULL,
                       email varchar UNIQUE NOT NULL,
                       hash_password varchar NOT NULL,
                       avatar_url varchar DEFAULT '',
                       bio text,
                       created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

