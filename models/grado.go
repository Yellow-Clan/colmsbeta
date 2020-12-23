package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Grado struct {
	Id_grado string `gorm:"primary_key;"`
	//Fecha  string
	Grados  string
	NivelId string `gorm:"size:191"`
	Nivel   Nivel  //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //`gorm:"embedded"` crea el compo nombres y codigo de alumnos
}

func (tab Grado) ToString() string {
	return fmt.Sprintf("id_grado: %d\nGrados: %s", tab.Id_grado, tab.Grados)
}

func (tab *Grado) BeforeCreate(*gorm.DB) error {
	tab.Id_grado = uuid.NewV4().String()
	return nil
}

// Alumno   Alumno //para crear el FK `gorm:"foreignkey:AlumnoId"`
