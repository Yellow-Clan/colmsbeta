package models

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

// TableName overrides the table name used by Nivel to `employee`
func (Nivel) TableName() string {
	return "nivel"
}

func (tab *Nivel) BeforeCreate(*gorm.DB) error {
	tab.Id_nivel = uuid.NewV4().String()
	return nil
}

type Nivel struct {
	Id_nivel string `gorm:"primary_key;column:id"`
	Niveles  string
}
