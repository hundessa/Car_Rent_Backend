package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Cars struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	CarName           string `json:"carname" validate:"required"`
	CarModel          string `json:"carmodel" validate:"required"`
	CarProductionYear string `json:"carproductionyear" validate:"required"`
	CarMileage        string `json:"carmileage" validate:"required"`
	Description       string `json:"description" validate:"required"`
	CarImage          string `json:"carimage" validate:"" gorm:""`
	CarPrice          string `json:"carprice" validate:"required"`
	CarRating         string `json:"carrating" validate:"required"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func (u *Cars) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
