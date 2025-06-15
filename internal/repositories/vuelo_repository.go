package repositories

import (
	"github.com/di3go35/agencia-viajes/internal/models"
	"gorm.io/gorm"
)

type VueloRepository interface {
	Create(vuelo *models.Vuelo) error
	FindAll() ([]models.Vuelo, error)
	FindByID(id uint) (*models.Vuelo, error)
	FindByOrigenDestino(origen, destino string) ([]models.Vuelo, error)
	Update(vuelo *models.Vuelo) error
	Delete(id uint) error
}

type vueloRepository struct {
	db *gorm.DB
}

func NewVueloRepository(db *gorm.DB) VueloRepository {
	return &vueloRepository{db: db}
}

func (r *vueloRepository) Create(vuelo *models.Vuelo) error {
	return r.db.Create(vuelo).Error
}

func (r *vueloRepository) FindAll() ([]models.Vuelo, error) {
	var vuelos []models.Vuelo
	err := r.db.Preload("Escalas").Find(&vuelos).Error
	return vuelos, err
}

// Implementar los demás métodos...
