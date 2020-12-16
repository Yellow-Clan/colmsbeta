package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Rol struct {
	Id       string `gorm:"primaryKey;"`
	Nombres  string
	Codigo   string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Usuarios []Usuario
}

func (tab Rol) ToString() string {
	return tab.Nombres
}

func (tab *Rol) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

func (rol Rol) FindAll(conn *gorm.DB) ([]Rol, error) {
	var roles []Rol
	if err := conn.Preload("Usuarios").Find(&roles).Error; err != nil {
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
