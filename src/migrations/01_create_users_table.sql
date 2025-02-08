CREATE TABLE IF NOT EXISTS users (
    id INT auto_increment PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) DEFAULT '',
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
    active TINYINT DEFAULT 1,
    CONSTRAINT login_unique UNIQUE (username),
    CONSTRAINT email_unigue UNIQUE (email)
);