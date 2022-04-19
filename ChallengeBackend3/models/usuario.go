package models

import (
	"Alura/ChallengeBackend3/database"
	"Alura/ChallengeBackend3/entities"
	"fmt"
)

func NovoUsuarioSis(usuario *entities.Usuarios) error {

	db := database.GetDatabase()

	err := db.Create(&usuario).Error

	if err != nil {
		fmt.Println(err)
	}
	return err
}

func TodosUsuarios() *[]entities.Usuarios {

	usuarios := []entities.Usuarios{}

	db := database.GetDatabase()

	err := db.Where("nome != ?", "Admin").Find(&usuarios).Error

	if err != nil {
		fmt.Println(err)
	}
	return &usuarios
}

func BuscarUsuarioPorEmail(email string) bool {

	var contagem int64
	usuario := entities.Usuarios{}

	db := database.GetDatabase()

	err := db.Where("email = ?", email).First(&usuario).Count(&contagem).Error

	if err != nil {
		fmt.Println(err)
	}

	return contagem >= 1

}
