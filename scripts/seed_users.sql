-- scripts/seed_users.sql
-- Test kullanıcılarını ekler

INSERT INTO users (name, email, password, role, created_at)
VALUES
    ('Admin User', 'admin@example.com', 'hashedpassword1', 'admin', NOW()),
    ('John Doe', 'john@example.com', 'hashedpassword2', 'user', NOW()),
    ('Jane Smith', 'jane@example.com', 'hashedpassword3', 'user', NOW());
