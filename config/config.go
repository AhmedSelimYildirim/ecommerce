package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	Env     string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	JWTSecret      string
	JWTExpireHours int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env yüklenemedi, devam ediliyor...")
	}

	expHours, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS"))

	return &Config{
		AppPort:        os.Getenv("APP_PORT"),
		Env:            os.Getenv("ENV"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPass:         os.Getenv("DB_PASS"),
		DBName:         os.Getenv("DB_NAME"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		JWTExpireHours: expHours,
	}
}
