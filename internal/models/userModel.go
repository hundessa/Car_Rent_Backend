package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
    gorm.Model
    FirstName string `json:"firstname" validate:"required"`
    LastName  string `json:"lastname" validate:"required"`
    Email     string `json:"email" validate:"required,email" gorm:"uniqueIndex"`
    Password  string `json:"password" validate:"required"`
    Role string `json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    if u.Role == "" {
        u.Role = "user"
    }
    return nil
}

// Validate validates the User fields
func (u *User) Validate() error {
    validate := validator.New()
    return validate.Struct(u)
}
