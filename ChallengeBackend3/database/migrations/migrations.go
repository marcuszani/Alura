package migrations

import (
	"Alura/ChallengeBackend3/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		entities.ArquivoCSV{},
		entities.ArquivosImportados{},
		entities.Usuarios{},
	)

	var contagem int64

	usuario := entities.Usuarios{}
	db.First(&usuario).Count(&contagem)

	if contagem == 0 {
		senhaEncriptada, _ := EncriptarSenha("123999")
		novoUsuario := entities.Usuarios{
			Nome:  "Admin",
			Email: "admin@email.com.br",
			Senha: senhaEncriptada,
		}

		db.Create(&novoUsuario)
	}

}
func EncriptarSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 6)
	return string(bytes), err
}
