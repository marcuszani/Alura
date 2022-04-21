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

func NovaImportacao(dadosDoArquivo *entities.ArquivosImportados) (bool, uint) {

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

	return ArquivoJaExiste, dadosDoArquivo.ID

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

func BuscarImportacoesPorID(id string) *entities.ArquivosImportados {
	db := database.GetDatabase()

	importado := entities.ArquivosImportados{}

	err := db.Where("id = ?", id).First(&importado).Error

	if err != nil {
		fmt.Println(err)
	}
	return &importado

}

func BuscarTodosDadosPorID(id string) *[]entities.ArquivoCSV {
	db := database.GetDatabase()

	dados := []entities.ArquivoCSV{}

	err := db.Where("arquivos_importados_id = ?", id).Order("data_hora_transacao desc").Find(&dados).Error

	if err != nil {
		fmt.Println(err)
	}

	return &dados

}
