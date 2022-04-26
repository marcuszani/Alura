package routes

import (
	"Alura/ChallengeBackend3/controllers"

	"github.com/gin-gonic/gin"
)

func CarregarRotasUploads(r *gin.Engine) {

	cadastrar := r.Group("cadastrar")
	{
		cadastrar.Use(controllers.Autenticado)
		cadastrar.GET("/", controllers.Cadastrar)
		cadastrar.POST("novo", controllers.NovaImportacao)
	}

	relatorios := r.Group("todasimportacoes")

	{
		relatorios.Use(controllers.Autenticado)
		relatorios.GET("/:id", controllers.RelatorioDados)
	}
}
