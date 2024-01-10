CREATE TABLE users (
                         id bigserial PRIMARY KEY,
                         username varchar UNIQUE NOT NULL,
                         email varchar UNIQUE NOT NULL,
                         hash_password varchar NOT NULL,
                         avatar_url varchar DEFAULT '',
                         bio text,
                         created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
                         id bigserial PRIMARY KEY,
                         user_id bigint NOT NULL,
                         video_url varchar NOT NULL,
                         caption text,
                         created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
                            id bigserial PRIMARY KEY,
                            user_id bigint NOT NULL,
                            post_id bigint NOT NULL,
                            text varchar NOT NULL,
                            created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE likes (
                         id bigserial PRIMARY KEY,
                         user_id bigint NOT NULL,
                         post_id bigint NOT NULL,
                         created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON users (username);

CREATE INDEX ON posts (user_id);

CREATE INDEX ON comments (post_id);

CREATE INDEX ON likes (user_id);

CREATE INDEX ON likes (post_id);

ALTER TABLE posts ADD FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

ALTER TABLE comments ADD FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

ALTER TABLE comments ADD FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE;

ALTER TABLE likes ADD FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

ALTER TABLE likes ADD FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE;
