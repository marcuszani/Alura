package routes

import (
	"Alura/ChallengeBackend3/controllers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	key = []byte("fu5q]WDJ{N8dz2HW7kG{Dlt09wp3O1-PzBO)5DREF4(NeFRV{2aP97v7Qao>8f44os-UFV2vy>iOPe8KKqrdF2Nmi{M_b2b{HpBXtMrjkry(cVJ0JNr5S1Qzrvem1uVPj>92qtm2h1wC1hilZsl)TC")
)

const cookieSessao = "ImportadorFinanceiro"

func CarregarRotasAuth(r *gin.Engine) {
	store := cookie.NewStore(key)
	store.Options(sessions.Options{
		Secure:   true,
		MaxAge:   36000,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	})

	r.Use(sessions.Sessions(cookieSessao, store))
	r.GET("/", controllers.FrmLogin)
	r.Any("/login", controllers.Login)

	r.GET("/logout", controllers.Logout)

	dashboard := r.Group("/")

	dashboard.Use(controllers.Autenticado)
	{
		dashboard.Any("/dashboard", controllers.Dashboard)

	}
}
