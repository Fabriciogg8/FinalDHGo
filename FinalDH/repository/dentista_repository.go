package repository

import (
	"database/sql"
	//"errors"
	"FinalDH/models"
)

type DentistaRepository struct {
	DB *sql.DB
}

func (r *DentistaRepository) GetDentistaByID(id int) (*models.Dentista, error) {
	var dentista models.Dentista
	err := r.DB.QueryRow("SELECT id, apellido, nombre, matricula FROM dentistas WHERE id = ?", id).
		Scan(&dentista.ID, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula)
	if err != nil {
		return nil, err
	}
	return &dentista, nil
}

func (r *DentistaRepository) CreateDentista(dentista *models.Dentista) error {
	_, err := r.DB.Exec("INSERT INTO dentistas (apellido, nombre, matricula) VALUES (?, ?, ?)",
		dentista.Apellido, dentista.Nombre, dentista.Matricula)
	return err
}

func (r *DentistaRepository) GetAllDentistas() ([]*models.Dentista, error) {
	rows, err := r.DB.Query("SELECT id, apellido, nombre, matricula FROM dentistas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dentistas []*models.Dentista
	for rows.Next() {
		var dentista models.Dentista
		if err := rows.Scan(&dentista.ID, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula); err != nil {
			return nil, err
		}
		dentistas = append(dentistas, &dentista)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dentistas, nil
}

func (r *DentistaRepository) UpdateDentista(dentista *models.Dentista) error {
	_, err := r.DB.Exec("UPDATE dentistas SET apellido = ?, nombre = ?, matricula = ? WHERE id = ?",
		dentista.Apellido, dentista.Nombre, dentista.Matricula, dentista.ID)
	return err
}

func (r *DentistaRepository) PatchDentista(id int, fields map[string]interface{}) error {
	query := "UPDATE dentistas SET "
	args := []interface{}{}
	i := 1
	for field, value := range fields {
		if i > 1 {
			query += ", "
		}
		query += field + " = ?"
		args = append(args, value)
		i++
	}
	query += " WHERE id = ?"
	args = append(args, id)

	_, err := r.DB.Exec(query, args...)
	return err
}

func (r *DentistaRepository) DeleteDentista(id int) error {
	_, err := r.DB.Exec("DELETE FROM dentistas WHERE id = ?", id)
	return err
}