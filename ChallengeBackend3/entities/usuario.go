package entities

import "gorm.io/gorm"

type Usuarios struct {
	gorm.Model
	Nome  string
	Email string
	Senha string
}
