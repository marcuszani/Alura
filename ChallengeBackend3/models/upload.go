package models

import (
	"Alura/ChallengeBackend3/database"
	"Alura/ChallengeBackend3/entities"
	"fmt"
)

func ImportarCSV(itens *[]entities.ArquivoCSV) error {
	db := database.GetDatabase()

	err := db.Create(&itens).Error

	if err != nil {
		fmt.Println(err)
	}

	return err

}

func NovaImportacao(dadosDoArquivo *entities.ArquivosImportados) bool {

	db := database.GetDatabase()

	resultado := db.Where("nome_arquivo = ?", dadosDoArquivo.NomeArquivo).First(&dadosDoArquivo)

	var ArquivoJaExiste bool

	if resultado.RowsAffected <= int64(0) {

		err := db.Create(&dadosDoArquivo).Error

		if err != nil {
			fmt.Println(err)
		}

		ArquivoJaExiste = true

	} else {

		fmt.Println("Arquivo ja Importado")

		ArquivoJaExiste = false

	}

	return ArquivoJaExiste

}

func BuscarTodasImportacoes() *[]entities.ArquivosImportados {
	db := database.GetDatabase()

	importacoes := []entities.ArquivosImportados{}

	err := db.Order("data_transacao desc").Find(&importacoes).Error

	if err != nil {
		fmt.Println(err)
	}

	return &importacoes

}

func BuscarTodosDados() *[]entities.ArquivoCSV {
	db := database.GetDatabase()

	dados := []entities.ArquivoCSV{}

	err := db.Order("data_hora_transacao desc").Find(&dados).Error

	if err != nil {
		fmt.Println(err)
	}

	return &dados

}
