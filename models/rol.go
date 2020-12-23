package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Rol struct {
	Id       string `gorm:"primaryKey;"`
	Nombre   string
	Codigo   string
	Personas []Persona
}

func (tab Rol) ToString() string {
	return tab.Nombre
}

func (tab *Rol) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

func (rol Rol) FindAll(conn *gorm.DB) ([]Rol, error) {
	var roles []Rol
	if err := conn.Preload("Personas").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (rol Rol) GetAll(conn *gorm.DB) ([]Rol, error) {
	var roles []Rol
	if err := conn.Find(&roles).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//fmt.Printf("Error: %v", err)
		//return fmt.Errorf("Error: %v", err)
		//continue
		return nil, fmt.Errorf("Error: %v", err)
	}
	return roles, nil
}
