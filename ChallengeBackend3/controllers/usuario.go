package controllers

import (
	"Alura/ChallengeBackend3/entities"
	"Alura/ChallengeBackend3/models"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CriarUsuario(c *gin.Context) {

	senhaGerada := GerarSenha(6)
	senhaEncriptada, err := EncriptarSenha(senhaGerada)

	if err != nil {
		fmt.Println(err)
	}

	usuario := entities.Usuarios{
		Nome:  c.Request.FormValue("nome"),
		Email: c.Request.FormValue("email"),
		Senha: senhaEncriptada,
	}

	valido, _ := models.BuscarUsuarioPorEmail(usuario.Email)

	if valido {
		c.String(http.StatusInternalServerError, "Email já cadastrado")
		return
	}

	//fmt.Println("funciona", usuario)
	err = models.NovoUsuarioSis(&usuario)

	if err != nil {
		c.String(http.StatusInternalServerError, "Falha ao Criar usuario")
		return
	} else {
		c.Redirect(http.StatusMovedPermanently, "/usuarios/cadastrar")
		//c.String(http.StatusOK, "Usuario Cadastrado com sucesso.")

		//EnviarEmail(usuario.Email, senhaGerada)

	}

}

func FrmCadastroUsuario(c *gin.Context) {

	usuarios := models.TodosUsuarios()

	c.HTML(http.StatusOK, "frmCadastrarUsuario.html", gin.H{"Usuarios": usuarios})

}

func GerarSenha(quantidade int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	senha := make([]byte, quantidade)

	for i := range senha {
		senha[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(senha)
}

func EncriptarSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	return string(bytes), err
}

func VerificarSenhaEncriptada(senha, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}

func DeletarUsuario(c *gin.Context) {
	//id := c.Request.URL.Query().Get("id")

	id := c.Params.ByName("id")

	if c.Request.FormValue("id") == "1" {
		c.String(http.StatusInternalServerError, "Operação não permitida")
		return
	} else {
		models.DeletarUsuario(id)
		c.Redirect(http.StatusMovedPermanently, "/usuarios/cadastrar")
	}

}

func EditarUsuario(c *gin.Context) {

	usuario := entities.Usuarios{
		Nome:  c.Request.FormValue("nome"),
		Email: c.Request.FormValue("email"),
	}

	if c.Request.FormValue("id") == "1" {
		c.String(http.StatusInternalServerError, "Operação não permitida")
		return
	} else {
		models.EditarUsuario(&usuario, c.Request.FormValue("id"))
		c.Redirect(http.StatusMovedPermanently, "/usuarios/cadastrar")
	}

}

func FrmEditarUsuario(c *gin.Context) {

	if c.Params.ByName("id") == "1" {
		c.String(http.StatusInternalServerError, "Operação não permitida")
		return
	} else {
		usuario := models.BuscarUsuarioPorID(c.Params.ByName("id"))

		c.HTML(http.StatusOK, "frmAlterarUsuario.html", gin.H{"Usuario": usuario})
	}

}
