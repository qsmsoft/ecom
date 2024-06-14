package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var Envs = initConfig()

func initConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     port,
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "postgres"),
		SSLMode:  getEnv("SSLMode", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
