package routes

import (
	"Alura/ChallengeBackend3/controllers"

	"github.com/gin-gonic/gin"
)

func CarregarRotas(r *gin.Engine) {

	usuarios := r.Group("/usuarios")
	{
		usuarios.Use(controllers.Autenticado)
		usuarios.GET("cadastrar", controllers.FrmCadastroUsuario)
		usuarios.POST("cadastrar/novo", controllers.CriarUsuario)
		usuarios.GET("remover/:id", controllers.DeletarUsuario)
		usuarios.GET("editar/:id", controllers.FrmEditarUsuario)
		usuarios.POST("editar/:id/salvar", controllers.EditarUsuario)
	}

}
