package models

import (
    "github.com/go-playground/validator/v10" 
)

// User represents a user in the system
type User struct {
    FirstName string `json:"firstname" validate:"required"`
    LastName  string `json:"lastname" validate:"required"`
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required"`
}

// Validate validates the User fields
func (u *User) Validate() error {
    validate := validator.New()
    return validate.Struct(u)
}
