package config

import (
	"fmt"
	"log"
)

// CreateTables crea las tablas necesarias en la base de datos
func CreateTables() {
	// Define las sentencias SQL para crear las tablas
	tables := []struct {
		name    string
		schema  string
	}{
		{
			name: "dentistas",
			schema: `CREATE TABLE IF NOT EXISTS dentistas (
				id INT AUTO_INCREMENT PRIMARY KEY,
				apellido VARCHAR(255) NOT NULL,
				nombre VARCHAR(255) NOT NULL,
				matricula VARCHAR(255) NOT NULL
			);`,
		},
		{
			name: "pacientes",
			schema: `CREATE TABLE IF NOT EXISTS pacientes (
				id INT AUTO_INCREMENT PRIMARY KEY,
				apellido VARCHAR(255) NOT NULL,
				nombre VARCHAR(255) NOT NULL,
				dni VARCHAR(20) NOT NULL,
				domicilio VARCHAR(255) NOT NULL,
				fecha_alta DATE NOT NULL
			);`,
		},
		{
			name: "turnos",
			schema: `CREATE TABLE IF NOT EXISTS turnos (
				id INT AUTO_INCREMENT PRIMARY KEY,
				fecha_hora DATETIME NOT NULL,
				descripcion TEXT,
				dentista_id INT,
				paciente_id INT,
				FOREIGN KEY (dentista_id) REFERENCES dentistas(id),
				FOREIGN KEY (paciente_id) REFERENCES pacientes(id)
			);`,
		},
	}

	// Ejecuta las sentencias SQL para crear las tablas
	for _, table := range tables {
		_, err := DB.Exec(table.schema)
		if err != nil {
			log.Fatalf("Error al crear la tabla %s: %v", table.name, err)
		}
		fmt.Printf("Tabla %s creada correctamente\n", table.name)
	}
}