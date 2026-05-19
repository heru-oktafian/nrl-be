-- Migration: Create admin_users table
-- Date: 2026-05-19

CREATE TABLE IF NOT EXISTS admin_users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default admin user (password: admin123)
-- In production, use bcrypt hash
INSERT INTO admin_users (username, password, name, email) 
VALUES ('admin', 'admin123', 'Administrator', 'admin@nuruldwi.com')
ON CONFLICT (username) DO NOTHING;