package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Ingresoadmi struct {
	Id string `gorm:"primary_key;"`
	//Fecha  string
	Semestre        string
	AdministradorId string        `gorm:"size:191"`
	Administrador   Administrador //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //`gorm:"embedded"` crea el compo nombres y codigo de alumnos
}

func (tab Ingresoadmi) ToString() string {
	return fmt.Sprintf("id: %d\nSemestre: %s", tab.Id, tab.Semestre)
}

func (tab *Ingresoadmi) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

// Alumno   Alumno //para crear el FK `gorm:"foreignkey:AlumnoId"`
