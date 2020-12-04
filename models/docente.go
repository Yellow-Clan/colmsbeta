package models

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

func (Docente) TableName() string {
	return "Docente"
}

func (tab *Docente) BeforeCreate(*gorm.DB) error {
	tab.Id_docente = uuid.NewV4().String()
	return nil
}

type Docente struct {
	Id_docente   string `gorm:"primary_key;column:Id_docente"`
	Nombre       string `gorm:"column:Nombre"`
	Curso_acargo string `gorm:"column:Curso_acargo"`
}
