package controllers

import (
	"github.com/di3go35/agencia-viajes/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type VueloController struct {
	service services.VueloService
}

func NewVueloController(service services.VueloService) *VueloController {
	return &VueloController{service: service}
}

func (c *VueloController) CreateVuelo(ctx echo.Context) error {
	var vuelo models.Vuelo
	if err := ctx.Bind(&vuelo); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.service.CreateVuelo(&vuelo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, vuelo)
}

// Implementar los demás handlers...
