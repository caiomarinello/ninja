DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;

CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    product_description VARCHAR(300) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(100) NOT NULL,
    stock INT NOT NULL
);

INSERT INTO products (product_name, product_description, price, category, stock) VALUES
('Travel Backpack', 'A durable and lightweight travel backpack with multiple compartments.', 79.99, 'Gear', 50),
('Waterproof Tent', 'A two-person waterproof tent suitable for all weather conditions.', 129.99, 'Camping', 25),
('Sleeping Bag', 'Compact and warm sleeping bag, ideal for camping and hiking.', 49.99, 'Camping', 100),
('Portable Stove', 'A lightweight portable stove with a compact design for easy transportation.', 39.99, 'Gear', 40),
('Hiking Boots', 'Comfortable and sturdy hiking boots for long trails.', 89.99, 'Clothing', 30),
('Travel Mug', 'Insulated travel mug to keep beverages hot or cold.', 19.99, 'Accessories', 200),
('First Aid Kit', 'Compact first aid kit with essential medical supplies.', 24.99, 'Accessories', 150),
('Travel Pillow', 'Soft and ergonomic travel pillow for long trips.', 14.99, 'Accessories', 120),
('Solar Charger', 'Portable solar charger for charging devices on the go.', 59.99, 'Electronics', 35),
('Compression Bags', 'Set of 5 compression bags to save space in your backpack.', 29.99, 'Gear', 80);

CREATE TABLE orders (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
