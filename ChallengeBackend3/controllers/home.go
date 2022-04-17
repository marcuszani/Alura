package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	//resultados := models.BuscarTodasImportacoes()

	c.HTML(http.StatusOK, "frmHome.html", nil)
}
