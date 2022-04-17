package database

import (
	"Alura/ChallengeBackend3/database/migrations"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	//conexao := "user=marcus dbname=teste password=marcus host=localhost sslmode=disable TimeZone=America/Sao_Paulo"

	database, err := gorm.Open(sqlite.Open("aluraback.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro:", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

}

func GetDatabase() *gorm.DB {
	return db
}

//
