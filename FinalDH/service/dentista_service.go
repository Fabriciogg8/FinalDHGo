package service

import (
	"FinalDH/models"
	"FinalDH/repository"
)

type DentistaService struct {
	Repo *repository.DentistaRepository
}

func (s *DentistaService) GetDentistaByID(id int) (*models.Dentista, error) {
	return s.Repo.GetDentistaByID(id)
}

func (s *DentistaService) CreateDentista(dentista *models.Dentista) error {
	return s.Repo.CreateDentista(dentista)
}

func (s *DentistaService) GetAllDentistas() ([]*models.Dentista, error) {
	return s.Repo.GetAllDentistas()
}

func (s *DentistaService) UpdateDentista(dentista *models.Dentista) error {
	return s.Repo.UpdateDentista(dentista)
}

func (s *DentistaService) PatchDentista(id int, fields map[string]interface{}) error {
	return s.Repo.PatchDentista(id, fields)
}

func (s *DentistaService) DeleteDentista(id int) error {
	return s.Repo.DeleteDentista(id)
}