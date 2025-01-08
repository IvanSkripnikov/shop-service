CREATE TABLE IF NOT EXISTS users (
    id INT auto_increment PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
    active TINYINT DEFAULT 1
);