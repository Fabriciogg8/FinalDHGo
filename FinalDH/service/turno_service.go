package service

import (
	"FinalDH/model"
	"FinalDH/repository"
)

type TurnoService struct {
    Repo *repository.TurnoRepository
}

func (s *TurnoService) GetAllTurnos() ([]models.Turno, error) {
    return s.Repo.GetAllTurnos()
}

func (s *TurnoService) CreateTurno(turno *models.Turno) error {
	return s.Repo.CreateTurno(turno)
}

func (s *TurnoService) GetTurnoByID(id int) (*models.Turno, error) {
	return s.Repo.GetTurnoByID(id)
}

func (s *TurnoService) UpdateTurno(turno *models.Turno) error {
	return s.Repo.UpdateTurno(turno)
}

func (s *TurnoService) PatchTurno(id int, fields map[string]interface{}) error {
	return s.Repo.PatchTurno(id, fields)
}

func (s *TurnoService) DeleteTurno(id int) error {
	return s.Repo.DeleteTurno(id)
}

func (s *TurnoService) CreateTurnoByDNIAndMatricula(dni, matricula string, turno *models.Turno) error {
	return s.Repo.CreateTurnoByDNIAndMatricula(dni, matricula, turno)
}

func (s *TurnoService) GetTurnoByDNI(dni string) (*models.Turno, error) {
	return s.Repo.GetTurnoByDNI(dni)
}