package pg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AhmedSelimYildirim/ecommerce/config"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}

	log.Println("Postgres connected")
	return db
}
