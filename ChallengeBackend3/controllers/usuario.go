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

	if models.BuscarUsuarioPorEmail(usuario.Email) {
		c.String(http.StatusInternalServerError, "Email já cadastrado")
		return
	}

	//fmt.Println("funciona", usuario)
	err = models.NovoUsuarioSis(&usuario)

	if err != nil {
		c.String(http.StatusInternalServerError, "Falha ao Criar usuario")
		return
	} else {
		c.String(http.StatusOK, "Usuario Cadastrado com sucesso.")
		c.Redirect(http.StatusMovedPermanently, "/usuarios/cadastrar")
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
	id := c.Request.URL.Query().Get("id")

	models.DeletarUsuario(id)

	c.Redirect(http.StatusMovedPermanently, "/usuarios/cadastrar")

}

func EditarUsuario(c *gin.Context) {

	usuario := entities.Usuarios{
		Nome:  c.Request.FormValue("nome"),
		Email: c.Request.FormValue("email"),
	}

	fmt.Println(usuario)

}

func FrmEditarUsuario(c *gin.Context) {
	fmt.Println(c.Param("id"))
}
