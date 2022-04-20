package main

import (
	"Alura/ChallengeBackend3/database"
	"Alura/ChallengeBackend3/routes"
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var server embed.FS

func main() {

	database.StartDB()

	r := gin.New()

	templ := template.Must(template.New("").ParseFS(server, "templates/*html"))

	r.SetHTMLTemplate(templ)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies([]string{"localhost"})
	routes.CarregarRotas(r)
	//r.LoadHTMLGlob("templates/*")
	r.Run(":8000")
}
