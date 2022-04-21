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
func BuscarUsuarioPorID(id string) *entities.Usuarios {
	usuario := entities.Usuarios{}

	db := database.GetDatabase()

	err := db.Where("id = ?", id).First(&usuario).Error

	if err != nil {
		fmt.Println(err)
	}

	return &usuario
}
func BuscarUsuarioPorEmail(email string) (bool, *entities.Usuarios) {

	var contagem int64
	usuario := entities.Usuarios{}

	db := database.GetDatabase()

	err := db.Where("email = ? OR nome = ?", email, email).First(&usuario).Count(&contagem).Error

	if err != nil {
		fmt.Println(err)
	}

	return (contagem >= 1), &usuario

}

func DeletarUsuario(id string) {

	usuario := entities.Usuarios{}

	db := database.GetDatabase()

	err := db.Where("id = ? AND id != ?", id, "1").Delete(&usuario).Error

	if err != nil {
		fmt.Println(err)
	}

}

func EditarUsuario(usuario *entities.Usuarios, id string) {
	db := database.GetDatabase()

	db.Where("id = ? AND id != ?", id, "1").Save(&usuario)
}

func VerificarUsuarios() bool {
	db := database.GetDatabase()

	var contagem int64
	usuario := entities.Usuarios{}

	err := db.First(&usuario).Count(&contagem).Error

	if err != nil {
		fmt.Println(err)
	}

	return contagem <= 0
}
