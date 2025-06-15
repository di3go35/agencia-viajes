package main

import (
	"github.com/di3go35/agencia-viajes/internal/config"
	"github.com/di3go35/agencia-viajes/internal/routes"

	//"github.com/di3go35/agencia-viajes/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Cargar configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Inicializar Echo
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Conectar a la base de datos

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Configurar rutas
	routes.SetupRoutes(e, db)

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":" + cfg.ServerPort))
}
