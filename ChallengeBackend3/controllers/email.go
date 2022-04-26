package controllers

import (
	"Alura/ChallengeBackend3/config"
	"fmt"
	"log"
	"net/smtp"
)

func EnviarEmail(email, senhaSys string) error {
	// Configuration
	//from := "bawoweb712@wowcg.com"
	usuario := config.Cfg.Email["Usuario"]

	fmt.Println(usuario)

	senha := config.Cfg.Email["Senha"]

	fmt.Println(senha)

	destinatario := []string{email}
	smtpHost := config.Cfg.Email["SMTP Server"]
	smtpPort := "2525"
	assunto := "Senha do Sistema"
	corpo := "Sua senha é: " + senhaSys

	msgString := "From: " + usuario + "\r\n" +
		"To: " + destinatario[0] + "\r\n" +
		"Subject: " + assunto + "\r\n\r\n" +
		corpo + "\r\n"

	msg := []byte(msgString)

	// Create authentication
	auth := smtp.PlainAuth("Teste Sistema Financeiro", usuario, senha, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, usuario, destinatario, msg)
	if err != nil {

		log.Fatal(err)
	}

	return err
}
