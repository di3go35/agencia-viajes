package models

import "gorm.io/gorm"

type Reserva struct {
	gorm.Model
	UsuarioID    uint       `gorm:"not null"`
	VueloID      uint       `gorm:"not null"`
	FechaReserva string     `gorm:"not null"`
	Estado       string     `gorm:"not null"` // Confirmada, Pendiente, Cancelada
	Total        float64    `gorm:"not null"`
	Pasajeros    []Pasajero `gorm:"foreignKey:ReservaID"`
}

type Pasajero struct {
	gorm.Model
	ReservaID uint   `gorm:"not null"`
	Nombre    string `gorm:"not null"`
	Apellido  string `gorm:"not null"`
	Pasaporte string `gorm:"not null"`
	Edad      int    `gorm:"not null"`
}
