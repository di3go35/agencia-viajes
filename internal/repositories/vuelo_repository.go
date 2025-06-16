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

func (r *vueloRepository) FindByID(id uint) (*models.Vuelo, error) {
	var vuelo models.Vuelo
	err := r.db.Preload("Escalas").First(&vuelo, id).Error
	return &vuelo, err
}

func (r *vueloRepository) FindByOrigenDestino(origen, destino string) ([]models.Vuelo, error) {
	var vuelos []models.Vuelo

	err := r.db.Preload("Escalas").Find(&vuelos, origen, destino).Error
	return vuelos, err
}

func (r *vueloRepository) Update(vuelo *models.Vuelo) error {
	return r.db.Save(vuelo).Error
}

func (r *vueloRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.Vuelo{}).Error
}

// Implementar los demás métodos...
