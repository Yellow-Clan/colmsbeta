package models

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)


func (tab *Usuario) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

type Usuario struct {
	
	Id 	 	string `gorm:"primary_key;"` //;default:UUID()
	Nombres string
	Rol 	string `gorm:"column:Rol_usuario"`
}


