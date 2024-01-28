package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person_info struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid"`
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationalize string
}
