CREATE TABLE IF NOT EXISTS items (
    id INT auto_increment PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    category_id INT NOT NULL,
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
    active TINYINT DEFAULT 1
);