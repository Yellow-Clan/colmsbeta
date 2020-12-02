package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Docente struct {
	Id      string `gorm:"primaryKey;"`
	Nombres string
	Codigo  string
	Cursos  []Curso
}

func (tab Docente) ToString() string {
	return tab.Nombres
}

func (tab *Docente) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

func (docente Docente) FindAll(conn *gorm.DB) ([]Docente, error) {
	var docentes []Docente
	if err := conn.Preload("Cursos").Find(&docentes).Error; err != nil {
		return nil, err
	}
	return docentes, nil
}

func (docente Docente) GetAll(conn *gorm.DB) ([]Docente, error) {
	var docentes []Docente
	if err := conn.Find(&docentes).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//fmt.Printf("Error: %v", err)
		//return fmt.Errorf("Error: %v", err)
		//continue
		return nil, fmt.Errorf("Error: %v", err)
	}
	return docentes, nil
}
