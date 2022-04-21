package entities

import (
	"time"

	"gorm.io/gorm"
)

type ArquivoCSV struct {
	gorm.Model
	ArquivosImportadosID uint
	BancoOrigem          string
	AgenciaOrigem        int
	ContaOrigem          string
	BancoDestino         string
	AgenciaDestino       int
	ContaDestino         string
	Valores              float64
	DataHoraTransacao    time.Time
}

type ArquivosImportados struct {
	gorm.Model
	NomeArquivo    string
	DataTransacao  time.Time
	DataImportacao time.Time
	UsuarioResp    string
}
