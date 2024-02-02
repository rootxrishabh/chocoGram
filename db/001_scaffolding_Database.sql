-- +goose Up

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    username VARCHAR(255) PRIMARY KEY
);

DROP TABLE IF EXISTS `friendships`;
CREATE TABLE friendships (
    user_a_username VARCHAR(255) REFERENCES users(username),
    user_b_username VARCHAR(255) REFERENCES users(username),
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (user_a_username, user_b_username)
);

-- +goose Down

DROP TABLE IF EXISTS `friendships`;
DROP TABLE IF EXISTS users;
