package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DB   DatabaseConfig
	App  AppConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type AppConfig struct {
	Port string
}

func Load() *Config {
	return &Config{
		DB: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USERNAME", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "nrl_be"),
		},
		App: AppConfig{
			Port: getEnv("APP_PORT", "3001"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func InitDB(cfg *Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	poolConfig.MaxConns = 10

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}

func RunMigrations(db *pgxpool.Pool) error {
	ctx := context.Background()

	// Create tables if they don't exist
	migrations := []string{
		// Profiles table
		`CREATE TABLE IF NOT EXISTS profiles (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			title VARCHAR(255),
			bio TEXT,
			email VARCHAR(255),
			phone VARCHAR(50),
			location VARCHAR(255),
			image_url TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Experiences table
		`CREATE TABLE IF NOT EXISTS experiences (
			id SERIAL PRIMARY KEY,
			profile_id INTEGER DEFAULT 1,
			title VARCHAR(255) NOT NULL,
			company VARCHAR(255),
			location VARCHAR(255),
			start_date DATE,
			end_date DATE,
			current BOOLEAN DEFAULT FALSE,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Skills table
		`CREATE TABLE IF NOT EXISTS skills (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			level INTEGER DEFAULT 50 CHECK (level >= 0 AND level <= 100),
			category VARCHAR(50),
			icon VARCHAR(100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Projects table
		`CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			tech_stack TEXT,
			image_url TEXT,
			demo_url TEXT,
			repo_url TEXT,
			featured BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Social Links table
		`CREATE TABLE IF NOT EXISTS social_links (
			id SERIAL PRIMARY KEY,
			platform VARCHAR(50) NOT NULL,
			url TEXT NOT NULL,
			icon VARCHAR(100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Tools table
		`CREATE TABLE IF NOT EXISTS tools (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			icon VARCHAR(100),
			url TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Contact Messages table
		`CREATE TABLE IF NOT EXISTS contact_messages (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			subject VARCHAR(255),
			message TEXT NOT NULL,
			read BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Insert default profile if not exists
		`INSERT INTO profiles (id, name, title, bio, email, location)
		 SELECT 1, 'Nurul Deere', 'Full Stack Developer', 'Passionate developer crafting digital experiences.', 'hello@nuruldeere.com', 'Indonesia'
		 WHERE NOT EXISTS (SELECT 1 FROM profiles WHERE id = 1)`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(ctx, migration); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	return nil
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}