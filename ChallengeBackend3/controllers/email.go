package controllers

import (
	"log"
	"net/smtp"
)

func EnviarEmail(email, senha string) {
	// Configuration
	from := "teste@gmail.com"
	password := "super_secret_password"
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Sua senha é: " + senha)

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}
