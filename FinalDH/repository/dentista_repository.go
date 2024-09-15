package repository

import (
	"database/sql"
	"FinalDH/models"
)

// DentistaRepository handles CRUD operations for dentists
type DentistaRepository struct {
	DB *sql.DB
}

// GetDentistaByID godoc
// @Summary Get a dentist by ID
// @Description Returns a single dentist by their ID
// @Tags dentists
// @Param id path int true "Dentist ID"
// @Success 200 
// @Failure 404  
// @Failure 500 
// @Router /dentists/{id} [get]
func (r *DentistaRepository) GetDentistaByID(id int) (*models.Dentista, error) {
	var dentista models.Dentista
	err := r.DB.QueryRow("SELECT id, apellido, nombre, matricula FROM dentistas WHERE id = ?", id).
		Scan(&dentista.ID, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula)
	if err != nil {
		return nil, err
	}
	return &dentista, nil
}

// GetAllDentistas godoc
// @Summary Get all dentists
// @Description Retrieves a list of all dentists
// @Tags dentists
// @Produce json
// @Success 200 
// @Failure 500 
// @Router /dentists [get]
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

// CreateDentista godoc
// @Summary Create a new dentist
// @Description Adds a new dentist to the database
// @Tags dentists
// @Param dentist body models.Dentista true "Dentist data"
// @Success 201 
// @Failure 400 
// @Failure 500 
// @Router /dentists [post]
func (r *DentistaRepository) CreateDentista(dentista *models.Dentista) error {
	_, err := r.DB.Exec("INSERT INTO dentistas (apellido, nombre, matricula) VALUES (?, ?, ?)",
		dentista.Apellido, dentista.Nombre, dentista.Matricula)
	return err
}

// UpdateDentista godoc
// @Summary Update a dentist by ID
// @Description Updates an existing dentist's details by ID
// @Tags dentists
// @Param id path int true "Dentist ID"
// @Param dentist body models.Dentista true "Dentist data"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500 
// @Router /dentists/{id} [put]
func (r *DentistaRepository) UpdateDentista(dentista *models.Dentista) error {
	_, err := r.DB.Exec("UPDATE dentistas SET apellido = ?, nombre = ?, matricula = ? WHERE id = ?",
		dentista.Apellido, dentista.Nombre, dentista.Matricula, dentista.ID)
	return err
}

// PatchDentista godoc
// @Summary Partially update a dentist by ID
// @Description Updates one or more fields of a dentist's information by ID
// @Tags dentists
// @Param id path int true "Dentist ID"
// @Param fields body map[string]interface{} true "Fields to update"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500  
// @Router /dentists/{id} [patch]
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

// DeleteDentista godoc
// @Summary Delete a dentist by ID
// @Description Removes a dentist from the database by their ID
// @Tags dentists
// @Param id path int true "Dentist ID"
// @Success 204 
// @Failure 404 
// @Failure 500 
// @Router /dentists/{id} [delete]
func (r *DentistaRepository) DeleteDentista(id int) error {
	_, err := r.DB.Exec("DELETE FROM dentistas WHERE id = ?", id)
	return err
}