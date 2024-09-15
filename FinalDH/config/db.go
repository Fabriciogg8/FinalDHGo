package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Abre la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la conexión a la base de datos: %v", err)
	}

	// Para probar la conexión
	if err = db.Ping(); err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	DB = db
	log.Println("Conexión a la base de datos exitosa")

	// Crear tablas
	CreateTables()
}