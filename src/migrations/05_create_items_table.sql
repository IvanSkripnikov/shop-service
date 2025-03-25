CREATE TABLE IF NOT EXISTS items (
    id INT auto_increment PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    category_id INT NOT NULL,
    user_category_id INT DEFAULT 1,
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
    active TINYINT DEFAULT 1
);