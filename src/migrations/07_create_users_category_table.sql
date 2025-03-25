CREATE TABLE IF NOT EXISTS user_category (
    id INT auto_increment PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created DATETIME(3) DEFAULT NOW(),
    active TINYINT DEFAULT 1
);