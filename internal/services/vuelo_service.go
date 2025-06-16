package services

import (
	"github.com/di3go35/agencia-viajes/internal/models"
	"github.com/di3go35/agencia-viajes/internal/repositories"
)

type VueloService interface {
	CreateVuelo(vuelo *models.Vuelo) error
	GetAllVuelos() ([]models.Vuelo, error)
	GetVueloByID(id uint) (*models.Vuelo, error)
	SearchVuelos(origen, destino string) ([]models.Vuelo, error)
	UpdateVuelo(vuelo *models.Vuelo) error
	DeleteVuelo(id uint) error
}

type vueloService struct {
	repo repositories.VueloRepository
}

func NewVueloService(repo repositories.VueloRepository) VueloService {
	return &vueloService{repo: repo}
}

func (s *vueloService) CreateVuelo(vuelo *models.Vuelo) error {
	return s.repo.Create(vuelo)
}

func (s *vueloService) GetAllVuelos() ([]models.Vuelo, error) {
	return s.repo.FindAll()
}

func (s *vueloService) GetVueloByID(id uint) (*models.Vuelo, error) {
	return s.repo.FindByID(id)
}

func (s *vueloService) SearchVuelos(origen, destino string) ([]models.Vuelo, error) {
	return s.repo.FindByOrigenDestino(origen, destino)
}

func (s *vueloService) UpdateVuelo(vuelo *models.Vuelo) error {
	return s.repo.Update(vuelo)
}

func (s *vueloService) DeleteVuelo(id uint) error {
	return s.repo.Delete(id)
}

// Implementar los demás métodos...
