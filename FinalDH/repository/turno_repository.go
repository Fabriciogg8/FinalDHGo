package repository

import (
	"database/sql"
	"FinalDH/model"
)

// TurnoRepository handles CRUD operations for turnos
type TurnoRepository struct {
	DB *sql.DB
}

// CreateTurno godoc
// @Summary Create a new turno
// @Description Adds a new turno to the database
// @Tags turnos
// @Param turno body models.Turno true "Turno data"
// @Success 201 
// @Failure 400 
// @Failure 500 
// @Router /turnos [post]
func (r *TurnoRepository) CreateTurno(turno *models.Turno) error {
	_, err := r.DB.Exec("INSERT INTO turnos (fecha_hora, descripcion, dentista_id, paciente_id) VALUES (?, ?, ?, ?)",
		turno.FechaHora, turno.Descripcion, turno.DentistaID, turno.PacienteID)
	return err
}

// GetTurnoByID godoc
// @Summary Get a turno by ID
// @Description Returns a single turno by their ID
// @Tags turnos
// @Param id path int true "Turno ID"
// @Success 200 
// @Failure 404  
// @Failure 500 
// @Router /turnos/{id} [get]
func (r *TurnoRepository) GetTurnoByID(id int) (*models.Turno, error) {
	var turno models.Turno
	err := r.DB.QueryRow("SELECT id, fecha_hora, descripcion, dentista_id, paciente_id FROM turnos WHERE id = ?", id).
		Scan(&turno.ID, &turno.FechaHora, &turno.Descripcion, &turno.DentistaID, &turno.PacienteID)
	if err != nil {
		return nil, err
	}
	return &turno, nil
}

// GetAllTurnos godoc
// @Summary Get all turnos
// @Description Retrieves a list of all turnos
// @Tags turnos
// @Success 200 {array} models.Turno "List of turnos"
// @Failure 500 
// @Router /turnos [get]
func (r *TurnoRepository) GetAllTurnos() ([]models.Turno, error) {
    rows, err := r.DB.Query("SELECT id, fecha_hora, descripcion, dentista_id, paciente_id FROM turnos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var turnos []models.Turno
    for rows.Next() {
        var t models.Turno
        err := rows.Scan(&t.ID, &t.FechaHora, &t.Descripcion, &t.DentistaID, &t.PacienteID)
        if err != nil {
            return nil, err
        }
        turnos = append(turnos, t)
    }
    return turnos, nil
}

// UpdateTurno godoc
// @Summary Update a turno by ID
// @Description Updates an existing turno's details by ID
// @Tags turnos
// @Param id path int true "Turno ID"
// @Param turno body models.Turno true "Turno data"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500 
// @Router /turnos/{id} [put]
func (r *TurnoRepository) UpdateTurno(turno *models.Turno) error {
	_, err := r.DB.Exec("UPDATE turnos SET fecha_hora = ?, descripcion = ?, dentista_id = ?, paciente_id = ? WHERE id = ?",
		turno.FechaHora, turno.Descripcion, turno.DentistaID, turno.PacienteID, turno.ID)
	return err
}

// PatchTurno godoc
// @Summary Partially update a turno by ID
// @Description Updates one or more fields of a turno's information by ID
// @Tags turnos
// @Param id path int true "Turno ID"
// @Param fields body map[string]interface{} true "Fields to update"
// @Success 200 
// @Failure 400 
// @Failure 404 
// @Failure 500  
// @Router /turnos/{id} [patch]
func (r *TurnoRepository) PatchTurno(id int, fields map[string]interface{}) error {
	query := "UPDATE turnos SET "
	args := []interface{}{}
// Changed the query to properly concatenate fields
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

// DeleteTurno godoc
// @Summary Delete a turno by ID
// @Description Removes a turno from the database by their ID
// @Tags turnos
// @Param id path int true "Turno ID"
// @Success 204 
// @Failure 404 
// @Failure 500 
// @Router /turnos/{id} [delete]
func (r *TurnoRepository) DeleteTurno(id int) error {
	_, err := r.DB.Exec("DELETE FROM turnos WHERE id = ?", id)
	return err
}

// CreateTurnoByDNIAndMatricula godoc
// @Summary Create a new turno by patient's DNI and dentist's matricula
// @Description Adds a new turno to the database using DNI and matricula
// @Tags turnos
// @Param dni query string true "Paciente DNI"
// @Param matricula query string true "Dentista Matricula"
// @Param turno body models.Turno true "Turno data"
// @Success 201 
// @Failure 400 
// @Failure 500 
// @Router /turnos/by-dni-matricula [post]
func (r *TurnoRepository) CreateTurnoByDNIAndMatricula(dni, matricula string, turno *models.Turno) error {
	// First, get the patient ID by DNI
	var pacienteID int
	err := r.DB.QueryRow("SELECT id FROM pacientes WHERE dni = ?", dni).Scan(&pacienteID)
	if err != nil {
		return err
	}

	// Then, get the dentist ID by matricula
	var dentistaID int
	err = r.DB.QueryRow("SELECT id FROM dentistas WHERE matricula = ?", matricula).Scan(&dentistaID)
	if err != nil {
		return err
	}

	// Insert the turno
	_, err = r.DB.Exec("INSERT INTO turnos (fecha_hora, descripcion, dentista_id, paciente_id) VALUES (?, ?, ?, ?)",
		turno.FechaHora, turno.Descripcion, dentistaID, pacienteID)
	return err
}

// GetTurnoByDNI godoc
// @Summary Get a turno by patient's DNI
// @Description Retrieves a turno's details by patient's DNI, including patient and dentist information
// @Tags turnos
// @Param dni query string true "Paciente DNI"
// @Success 200 
// @Failure 404 
// @Failure 500 
// @Router /turnos/by-dni [get]
func (r *TurnoRepository) GetTurnoByDNI(dni string) (*models.Turno, error) {
    var turno models.Turno
    query := `
        SELECT t.id, t.fecha_hora, t.descripcion, t.dentista_id, t.paciente_id
        FROM turnos t
        JOIN pacientes p ON t.paciente_id = p.id
        WHERE p.dni = ?`

    err := r.DB.QueryRow(query, dni).
        Scan(&turno.ID, &turno.FechaHora, &turno.Descripcion, &turno.DentistaID, &turno.PacienteID)
    if err != nil {
        return nil, err
    }
    return &turno, nil
}