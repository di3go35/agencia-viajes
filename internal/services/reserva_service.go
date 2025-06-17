package services

import (
	"github.com/di3go35/agencia-viajes/internal/models"
	"github.com/di3go35/agencia-viajes/internal/repositories"
)

type ReservaService interface {
	CreateReserva(reserva models.Reserva) error
	FindReservasAll() ([]models.Reserva, error)
	FindReservaByID(id uint) (models.Reserva, error)
	UpdateReserva(reserva models.Reserva) error
	DeleteReserva(id uint) error
}

type reservaService struct {
	repo repositories.ReservaRepository
}

func NewReservaService(repo repositories.ReservaRepository) ReservaService {
	return &reservaService{repo: repo}
}

func (s *reservaService) CreateReserva(reserva models.Reserva) error {
	return s.repo.Create(reserva)
}

func (s *reservaService) FindReservasAll() ([]models.Reserva, error) {
	return s.repo.FindAll()
}

func (s *reservaService) FindReservaByID(id uint) (models.Reserva, error) {
	return s.repo.FindByID(id)
}

func (s *reservaService) UpdateReserva(reserva models.Reserva) error {
	return s.repo.Update(reserva)
}

func (s *reservaService) DeleteReserva(id uint) error {
	return s.repo.Delete(id)
}
