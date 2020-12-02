package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Curso struct {
	Id       string `gorm:"primary_key;"`
	Semestre string
	AlumnoId string `gorm:"size:191"`
	Alumno   Alumno //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //`gorm:"embedded"` crea el compo nombres y codigo de alumnos
}

func (tab Curso) ToString() string {
	return fmt.Sprintf("id: %d\nSemestre: %s", tab.Id, tab.Semestre)
}

func (tab *Curso) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

// Alumno   Alumno //para crear el FK `gorm:"foreignkey:AlumnoId"`
