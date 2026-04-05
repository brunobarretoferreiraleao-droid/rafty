package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	log.Println("Banco conectado com sucesso.")
	return db
}
