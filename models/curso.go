package models

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

func (Curso) TableName() string {
	return "curso"
}

func (tab *Curso) BeforeCreate(*gorm.DB) error {
	tab.Id_curso = uuid.NewV4().String()
	return nil
}

type Curso struct {
	Id_curso string `gorm:"primary_key;column:id"`

	Nombrecurso string
}
