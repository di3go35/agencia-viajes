package repositories

import (
	"github.com/di3go35/agencia-viajes/internal/models"
	"gorm.io/gorm"
)

type ReservaRepository interface {
	Create(reserva models.Reserva) error
	FindAll() ([]models.Reserva, error)
	FindByID(id uint) (models.Reserva, error)
	Update(reserva models.Reserva) error
	Delete(id uint) error
}

type reservaRepository struct {
	db *gorm.DB
}

func NewReservaRepository(db *gorm.DB) ReservaRepository {
	return &reservaRepository{db: db}
}

func (r *reservaRepository) Create(reserva models.Reserva) error {
	return r.db.Create(reserva).Error
}

func (r *reservaRepository) FindAll() ([]models.Reserva, error) {
	var reservas []models.Reserva
	err := r.db.Find(&reservas).Error
	return reservas, err
}

func (r *reservaRepository) FindByID(id uint) (models.Reserva, error) {
	var reserva models.Reserva
	err := r.db.First(&reserva, id).Error
	return reserva, err
}

func (r *reservaRepository) Update(reserva models.Reserva) error {
	return r.db.Save(reserva).Error
}

func (r *reservaRepository) Delete(id uint) error {
	return r.db.Delete(&models.Reserva{}, id).Error
}
