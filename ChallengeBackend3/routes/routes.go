package routes

import (
	"Alura/ChallengeBackend3/controllers"

	"github.com/gin-gonic/gin"
)

func CarregarRotas(r *gin.Engine) {

	raiz := r.Group("cadastrar")
	{
		raiz.Use(controllers.Autenticado)
		raiz.GET("/", controllers.Cadastrar)
		raiz.POST("novo", controllers.NovaImportacao)
	}

	usuarios := r.Group("usuarios")
	{
		usuarios.Use(controllers.Autenticado)
		usuarios.GET("cadastrar", controllers.FrmCadastroUsuario)
		usuarios.POST("cadastrar/novo", controllers.CriarUsuario)
		usuarios.GET("deletar/:id", controllers.DeletarUsuario)
		usuarios.GET("editar/:id", controllers.FrmEditarUsuario)
		usuarios.POST("editar/:id/salvar", controllers.EditarUsuario)
	}

	relatorios := r.Group("todasimportacoes")
	{
		relatorios.Use(controllers.Autenticado)
		relatorios.GET("/:id", controllers.RelatorioDados)
	}
	//r.GET("/", controllers.Home)
	//r.GET("/cadastrar", controllers.Cadastrar)
	//r.POST("/cadastrar/novo", controllers.NovaImportacao)
	//r.GET("/todasimportacoes/:id", controllers.RelatorioDados)

	//r.GET("/usuarios/cadastrar", controllers.FrmCadastroUsuario)
	//r.POST("/usuarios/cadastrar/novo", controllers.CriarUsuario)

	//r.GET("/usuarios/deletar/:id", controllers.DeletarUsuario)
	//r.GET("usuarios/editar/:id", controllers.FrmEditarUsuario)
	//r.POST("/usuarios/editar/:id/salvar", controllers.EditarUsuario)
	//r.GET("/usuarios/editar/:id", controllers.EditarUsuario)
}
