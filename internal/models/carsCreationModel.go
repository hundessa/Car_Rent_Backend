package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)


type Cars struct {
	gorm.Model
	CarName string `json:"carname" validate:"required" gorm:"uniqueIndex"`
	CarModel string `json:"carmodel" validate:"required" gorm:"uniqueIndex"`
	CarProductionYear string `json:"carproductionyear" validate:"required" gorm:"uniqueIndex"`
	CarMileage string `json:"carmileage" validate:"required" gorm:"uniqueIndex"`
	Description string `json:"description" validate:"required" gorm:"uniqueIndex"`
	CarImage string `json:"carimage" validate:"required" gorm:"uniqueIndex"`
	CarPrice string `json:"carprice" validate:"required" gorm:"uniqueIndex"`
	CarRating string `json:"carrating" validate:"required" gorm:"uniqueIndex"`
}


func (u *Cars) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}