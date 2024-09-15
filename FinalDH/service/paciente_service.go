package service

import (
	"FinalDH/model"
	"FinalDH/repository"
)

type PacienteService struct {
	Repo *repository.PacienteRepository
}

func (s *PacienteService) GetPacienteByID(id int) (*models.Paciente, error) {
	return s.Repo.GetPacienteByID(id)
}

func (s *PacienteService) CreatePaciente(paciente *models.Paciente) error {
	return s.Repo.CreatePaciente(paciente)
}

func (s *PacienteService) GetAllPacientes() ([]*models.Paciente, error) {
	return s.Repo.GetAllPacientes()
}

func (s *PacienteService) UpdatePaciente(paciente *models.Paciente) error {
	return s.Repo.UpdatePaciente(paciente)
}

func (s *PacienteService) PatchPaciente(id int, fields map[string]interface{}) error {
	return s.Repo.PatchPaciente(id, fields)
}

func (s *PacienteService) DeletePaciente(id int) error {
	return s.Repo.DeletePaciente(id)
}