package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Administrador struct {
	Id           string `gorm:"primaryKey;"`
	Nombres      string
	Codigo       string
	Ingresoadmis []Ingresoadmi
}

func (tab Administrador) ToString() string {
	return tab.Nombres
}

func (tab *Administrador) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

func (administrador Administrador) FindAll(conn *gorm.DB) ([]Administrador, error) {
	var administradores []Administrador
	if err := conn.Preload("Ingresoadmis").Find(&administradores).Error; err != nil {
		return nil, err
	}
	return administradores, nil
}

func (administrador Administrador) GetAll(conn *gorm.DB) ([]Administrador, error) {
	var administradores []Administrador
	if err := conn.Find(&administradores).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//fmt.Printf("Error: %v", err)
		//return fmt.Errorf("Error: %v", err)
		//continue
		return nil, fmt.Errorf("Error: %v", err)
	}
	return administradores, nil
}
