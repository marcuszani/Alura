package routes

import (
	"Alura/ChallengeBackend3/controllers"

	"github.com/gin-gonic/gin"
)

func CarregarRotas(r *gin.Engine) {

	r.GET("/", controllers.Home)
	r.GET("/cadastrar", controllers.Cadastrar)
	r.POST("/cadastrar/novo", controllers.NovaImportacao)
	r.GET("/todasimportacoes", controllers.RelatorioDados)

	r.GET("/usuarios/cadastrar", controllers.FrmCadastroUsuario)
	r.POST("/usuarios/cadastrar/novo", controllers.CriarUsuario)

}
