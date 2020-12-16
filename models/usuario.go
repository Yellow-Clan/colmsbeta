package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Usuario struct {
	Id       string `gorm:"primary_key;"`
	Usuario  string
	Semestre string
	RolId    string `gorm:"size:191"`
	Rol      Rol    //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //`gorm:"embedded"` crea el compo nombres y codigo de alumnos
}

func (tab Usuario) ToString() string {
	return fmt.Sprintf("id: %d\nSemestre: %s", tab.Id, tab.Semestre)
}

func (tab *Usuario) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

// Alumno   Alumno //para crear el FK `gorm:"foreignkey:AlumnoId"`
