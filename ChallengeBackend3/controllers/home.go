package controllers

import (
	"Alura/ChallengeBackend3/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Home(c *gin.Context) {

	//resultados := models.BuscarTodasImportacoes()

	c.HTML(http.StatusOK, "frmHome.html", nil)
}

var keyuser = "usuario"

func FrmLogin(c *gin.Context) {

	session := sessions.Default(c)

	usuario := session.Get(keyuser)

	if usuario == nil {
		//c.AbortWithStatus(http.StatusUnauthorized)
		c.HTML(http.StatusOK, "frmAuth.html", gin.H{})
		//c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
	}

}

func Dashboard(c *gin.Context) {

	session := sessions.Default(c)

	// if session.Get(keyuser) == nil {
	// 	c.Redirect(http.StatusTemporaryRedirect, "/")
	// }

	usuario := session.Get(keyuser)

	c.HTML(http.StatusOK, "frmDashboard.html", gin.H{
		keyuser: usuario,
	})

}

func Autenticado(c *gin.Context) {

	session := sessions.Default(c)

	usuario := session.Get(keyuser)

	if usuario == nil {
		//c.AbortWithStatus(http.StatusUnauthorized)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.Next()

}

func Login(c *gin.Context) {

	username := c.PostForm("user")
	password := c.PostForm("password")

	// Validacao de entrada

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.String(http.StatusBadRequest, "Parametros Vazios")
		return
	}

	_, usuario := models.BuscarUsuarioPorEmail(username)

	if !((username == usuario.Email || username == usuario.Nome) && CompararSenha(password, usuario.Senha)) {
		c.String(http.StatusUnauthorized, "Falha no Login")

		fmt.Println(password, usuario.Senha)
		return
	}

	// if username != usuario.Email || password != senhaEncriptada {
	// 	c.String(http.StatusUnauthorized, "Falha no Login")
	// 	return
	// }

	session := sessions.Default(c)
	session.Set(keyuser, username)

	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, "Falha ao salvar a sessão")
		return
	}

	//c.String(http.StatusOK, "Autenticado com sucesso")

	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")

}

func Logout(c *gin.Context) {
	session := sessions.Default(c)

	usuario := session.Get(keyuser)

	if usuario == nil {
		c.String(http.StatusBadRequest, "Token Invalidado")
		return
	}

	// Revoke users authentication
	session.Delete(keyuser)

	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, "Falha ao salvar a sessão")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func CompararSenha(senha, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))

	return err == nil
}
