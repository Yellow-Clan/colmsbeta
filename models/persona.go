package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Persona struct {
	Id string `gorm:"primary_key;"`
	//Fecha  string
	Semestre string
	Email    string `gorm:"type:varchar(100);unique_index"`
	RolId    string `gorm:"size:191"`
	Rol      Rol
}

func (tab Persona) ToString() string {
	return fmt.Sprintf("id: %d\nSemestre: %s", tab.Id, tab.Semestre)
}

func (tab *Persona) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

// Alumno Alumno //para crear el FK `gorm:"foreignkey:AlumnoId"`
//`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//`gorm:"embedded"` crea el compo nombres y codigo de alumnos
