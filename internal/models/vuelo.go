package models

import "gorm.io/gorm"

type Vuelo struct {
	gorm.Model
	Aerolinea           string   `gorm:"not null"`
	NumeroVuelo         string   `gorm:"not null;unique"`
	Origen              string   `gorm:"not null"`
	Destino             string   `gorm:"not null"`
	FechaSalida         string   `gorm:"not null"`
	FechaLlegada        string   `gorm:"not null"`
	AsientosTotales     int      `gorm:"not null"`
	AsientosDisponibles int      `gorm:"not null"`
	Precio              float64  `gorm:"not null"`
	Escalas             []Escala `gorm:"foreignKey:VueloID"`
}

type Escala struct {
	gorm.Model
	VueloID      uint   `gorm:"not null"`
	Ciudad       string `gorm:"not null"`
	TiempoEspera string `gorm:"not null"`
}
