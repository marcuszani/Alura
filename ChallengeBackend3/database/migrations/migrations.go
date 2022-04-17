package migrations

import (
	"Alura/ChallengeBackend3/entities"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		entities.ArquivoCSV{},
		entities.ArquivosImportados{},
	)
}
