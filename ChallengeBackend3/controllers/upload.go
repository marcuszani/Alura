package controllers

import (
	"Alura/ChallengeBackend3/entities"
	"Alura/ChallengeBackend3/models"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Cadastrar(c *gin.Context) {

	resultados := models.BuscarTodasImportacoes()
	usuarios := models.TodosUsuarios(true)

	c.HTML(http.StatusOK, "frmCadastrar.html", gin.H{
		"objetosImportados": resultados,
		"Usuarios":          usuarios,
	})

}

func NovaImportacao(c *gin.Context) {

	session := sessions.Default(c)

	err := c.Request.ParseMultipartForm(10 << 20)

	if err != nil {
		log.Println("Tamanho do Arquivo", err)
	}

	object, handler, err := c.Request.FormFile("meuArquivo")

	if err != nil {
		log.Println("Erro ao Coletar Arquivo", err)
	}
	defer object.Close()
	// Verificar se o arquivo esta em branco antes de adicionar no banco

	csvBuf := bytes.NewBuffer(nil)

	if _, err := io.Copy(csvBuf, object); err != nil {
		log.Println("Problema pra abrir arquivo ", err)
	}

	csvLines, err := csv.NewReader(csvBuf).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	if csvLines == nil {
		fmt.Println("Arquivo Vazio")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//var dataDaTransacao time.Time

	csvTodasLinhas := []entities.ArquivoCSV{}
	// valoresVazios := false

	dados, dataDaTransacao := converterCsvParaStruct(csvLines)

	idUsuario := fmt.Sprintf("%v", session.Get("usuarioID"))

	idUsuarioConvertido, _ := strconv.ParseInt(idUsuario, 10, 32)

	dadosDoArquivo := entities.ArquivosImportados{
		NomeArquivo:    handler.Filename,
		DataTransacao:  dataDaTransacao,
		DataImportacao: time.Now().Local(),
		UsuarioResp:    uint(idUsuarioConvertido),
	}

	_, idTransacao := models.NovaImportacao(&dadosDoArquivo)

	if true {

		fmt.Println("Nome do Arquivo: ", handler.Filename)
		fmt.Println("O id da transação é:", idTransacao)

		dadosCorrigidos := []entities.ArquivoCSV{}

		for _, itens := range *dados {
			linhas := entities.ArquivoCSV{
				ArquivosImportadosID: idTransacao,
				BancoOrigem:          itens.BancoOrigem,
				AgenciaOrigem:        itens.AgenciaOrigem,
				ContaOrigem:          itens.ContaOrigem,
				BancoDestino:         itens.BancoDestino,
				AgenciaDestino:       itens.AgenciaDestino,
				ContaDestino:         itens.ContaDestino,
				Valores:              itens.Valores,
				DataHoraTransacao:    itens.DataHoraTransacao,
			}

			dadosCorrigidos = append(dadosCorrigidos, linhas)
		}

		err = models.ImportarCSV(&dadosCorrigidos)

		if err != nil {
			fmt.Println("Erro ao cadastrar no banco", err)
		} else {
			fmt.Println(csvTodasLinhas)
		}
		fmt.Println(handler.Size)
		c.Redirect(http.StatusMovedPermanently, "/cadastrar")
		return

	} else {
		c.String(http.StatusInternalServerError, "Arquivo Ja Importado")
		return
	}

}

func converterData(data string) time.Time {
	timezone, _ := time.LoadLocation("America/Sao_Paulo")
	converterData, err := time.ParseInLocation("2006-01-02T15:04:05", data, timezone)

	if err != nil {
		fmt.Println(err)
	}

	return converterData
}

func verificarPresencaEspacoString(coluna []string) bool {

	valoresVazios := false

	for _, colunaVazia := range coluna {

		colunaVazia := strings.ReplaceAll(colunaVazia, " ", "")

		if colunaVazia == "" || colunaVazia == " " {
			valoresVazios = true
		}
	}

	return valoresVazios

}

func converterCsvParaStruct(csvLines [][]string) (dados *[]entities.ArquivoCSV, data time.Time) {

	var dataDaTransacao time.Time

	csvTodasLinhas := []entities.ArquivoCSV{}
	// valoresVazios := false

	for contador, coluna := range csvLines {

		agenciaInt, _ := strconv.Atoi(coluna[1])
		agenciaDestInt, _ := strconv.Atoi(coluna[4])
		valoresFloat, _ := strconv.ParseFloat(coluna[6], 64)
		// valoresVazios = false

		if contador == 0 {
			dataDaTransacao = converterData(coluna[7])
		}

		dados := entities.ArquivoCSV{

			BancoOrigem:       coluna[0],
			AgenciaOrigem:     agenciaInt,
			ContaOrigem:       coluna[2],
			BancoDestino:      coluna[3],
			AgenciaDestino:    agenciaDestInt,
			ContaDestino:      coluna[5],
			Valores:           valoresFloat,
			DataHoraTransacao: converterData(coluna[7]),
		}

		if dataDaTransacao.Format("2006-01-02") == dados.DataHoraTransacao.Format("2006-01-02") && !verificarPresencaEspacoString(coluna) {
			csvTodasLinhas = append(csvTodasLinhas, dados)
		} else {
			fmt.Println("Arquivo possui datas diferentes")
		}
	}
	return &csvTodasLinhas, dataDaTransacao
}

func RelatorioDados(c *gin.Context) {

	transacao := models.BuscarImportacoesPorID(c.Params.ByName("id"))
	detalhes := models.BuscarTodosDadosPorID(c.Params.ByName("id"))

	c.HTML(http.StatusOK, "frmTodosDados.html", gin.H{
		"objetosImportados": detalhes,
		"dadosTransacao":    transacao,
	})
}
