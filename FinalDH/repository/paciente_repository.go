package repository

import (
	"database/sql"
	"FinalDH/model"
)

// PacienteRepository handles CRUD operations for patients
type PacienteRepository struct {
	DB *sql.DB
}

// GetPacienteByID godoc
// @Summary Get a patient by ID
// @Description Returns a single patient by their ID
// @Tags pacientes
// @Param id path int true "Patient ID"
// @Success 200 
// @Failure 404  
// @Failure 500 
// @Router /pacientes/{id} [get]
func (r *PacienteRepository) GetPacienteByID(id int) (*models.Paciente, error) {
	var paciente models.Paciente
	err := r.DB.QueryRow("SELECT id, apellido, nombre, dni, domicilio, fecha_alta FROM pacientes WHERE id = ?", id).
		Scan(&paciente.ID, &paciente.Apellido, &paciente.Nombre, &paciente.DNI, &paciente.Domicilio, &paciente.FechaAlta)
	if err != nil {
		return nil, err
	}
	return &paciente, nil
}

// GetAllPacientes godoc
// @Summary Get all patients
// @Description Retrieves a list of all patients
// @Tags pacientes
// @Produce json
// @Success 200 
// @Failure 500 
// @Router /pacientes [get]
func (r *PacienteRepository) GetAllPacientes() ([]*models.Paciente, error) {
	rows, err := r.DB.Query("SELECT id, apellido, nombre, dni, domicilio, fecha_alta FROM pacientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pacientes []*models.Paciente
	for rows.Next() {
		var paciente models.Paciente
		if err := rows.Scan(&paciente.ID, &paciente.Apellido, &paciente.Nombre, &paciente.DNI, &paciente.Domicilio, &paciente.FechaAlta); err != nil {
			return nil, err
		}
		pacientes = append(pacientes, &paciente)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pacientes, nil
}

// CreatePaciente godoc
// @Summary Create a new patient
// @Description Adds a new patient to the database
// @Tags pacientes
// @Param paciente body models.Paciente true "Patient data"
// @Success 201 
// @Failure 400 
// @Failure 500 
// @Router /pacientes [post]
func (r *PacienteRepository) CreatePaciente(paciente *models.Paciente) error {
	_, err := r.DB.Exec("INSERT INTO pacientes (apellido, nombre, dni, domicilio, fecha_alta) VALUES (?, ?, ?, ?, ?)",
		paciente.Apellido, paciente.Nombre, paciente.DNI, paciente.Domicilio, paciente.FechaAlta)
	return err
}

// UpdatePaciente godoc
// @Summary Update a patient by ID
// @Description Updates an existing patient's details by ID
// @Tags pacientes
// @Param id path int true "Patient ID"
// @Param paciente body models.Paciente true "Patient data"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500 
// @Router /pacientes/{id} [put]
func (r *PacienteRepository) UpdatePaciente(paciente *models.Paciente) error {
	_, err := r.DB.Exec("UPDATE pacientes SET apellido = ?, nombre = ?, dni = ?, domicilio = ?, fecha_alta = ? WHERE id = ?",
		paciente.Apellido, paciente.Nombre, paciente.DNI, paciente.Domicilio, paciente.FechaAlta, paciente.ID)
	return err
}

// PatchPaciente godoc
// @Summary Partially update a patient by ID
// @Description Updates one or more fields of a patient's information by ID
// @Tags pacientes
// @Param id path int true "Patient ID"
// @Param fields body map[string]interface{} true "Fields to update"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500  
// @Router /pacientes/{id} [patch]
func (r *PacienteRepository) PatchPaciente(id int, fields map[string]interface{}) error {
	query := "UPDATE pacientes SET "
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

// DeletePaciente godoc
// @Summary Delete a patient by ID
// @Description Removes a patient from the database by their ID
// @Tags pacientes
// @Param id path int true "Patient ID"
// @Success 204 
// @Failure 404 
// @Failure 500 
// @Router /pacientes/{id} [delete]
func (r *PacienteRepository) DeletePaciente(id int) error {
	_, err := r.DB.Exec("DELETE FROM pacientes WHERE id = ?", id)
	return err
}