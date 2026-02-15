package config

import (
	"net/url"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "user"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5433"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "rbac_db"
	}

	userInfo := url.UserPassword(dbUser, dbPassword)
	databaseURL := "postgres://" + userInfo.String() + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "secret"
	}

	return Config{
		Port:        port,
		DatabaseURL: databaseURL,
		JWTSecret:   jwtSecret,
	}
}
