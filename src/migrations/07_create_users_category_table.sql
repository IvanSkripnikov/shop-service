CREATE TABLE IF NOT EXISTS user_category (
    id INT auto_increment PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    active TINYINT DEFAULT 1
);