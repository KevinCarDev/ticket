package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Cliente string `gorm:"not null"`
	Origen  string
	Destino string
	Price   float64 `gorm:"check:price>= 0"`
	Date    time.Time
}
