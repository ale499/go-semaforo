package models

import "gorm.io/gorm"

type Semaforo struct {
	gorm.Model
	Estado string

	Ubicacion string
}
