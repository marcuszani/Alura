package entities

import "time"

type ArquivoCSV struct {
	BancoOrigem       string
	AgenciaOrigem     int
	ContaOrigem       string
	BancoDestino      string
	AgenciaDestino    int
	ContaDestino      string
	Valores           float64
	DataHoraTransacao time.Time
}

type ArquivosImportados struct {
	NomeArquivo    string
	DataTransacao  time.Time
	DataImportacao time.Time
}
