CREATE TABLE IF NOT EXISTS cart (
                                    id SERIAL PRIMARY KEY,
                                    user_id INT NOT NULL REFERENCES users(id),
    product_id INT NOT NULL REFERENCES products(id),
    quantity INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
    );