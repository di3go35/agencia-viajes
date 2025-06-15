package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Inicializar repositorios
	vueloRepo := repositories.NewVueloRepository(db)
	reservaRepo := repositories.NewReservaRepository(db)

	// Inicializar servicios
	vueloService := services.NewVueloService(vueloRepo)
	reservaService := services.NewReservaService(reservaRepo)

	// Inicializar controladores
	vueloCtrl := controllers.NewVueloController(vueloService)
	reservaCtrl := controllers.NewReservaController(reservaService)

	// Grupo de rutas API
	api := e.Group("/api/v1")

	// Rutas de vuelos
	vuelos := api.Group("/vuelos")
	vuelos.POST("", vueloCtrl.CreateVuelo)
	vuelos.GET("", vueloCtrl.GetAllVuelos)
	vuelos.GET("/:id", vueloCtrl.GetVueloByID)
	vuelos.GET("/buscar", vueloCtrl.SearchVuelos)
	vuelos.PUT("/:id", vueloCtrl.UpdateVuelo)
	vuelos.DELETE("/:id", vueloCtrl.DeleteVuelo)

	// Rutas de reservas
	reservas := api.Group("/reservas")
	reservas.POST("", reservaCtrl.CreateReserva)
	reservas.GET("", reservaCtrl.GetAllReservas)
	// ... otras rutas de reservas
}
