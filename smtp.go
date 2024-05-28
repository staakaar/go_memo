//go:build ignore

package main

import (
	"log"
	"net/smtp"

	"github.com/jhillyerd/enmime"
)

func main() {
	smtpHost := "my-main-server:25"
	smtpAuth := smtp.PlainAuth(
		"example.com",
		"example-user",
		"example-password",
		"auth.example.com",
	)

	sender := enmime.NewSMTP(smtpHost, smtpAuth)

	master := enmime.Builder().From("Go太郎", "gotaro@example.com").Subject("main title").Text([]byte("テキストメール本文")).HTML([]byte("<p>HTMLメール<b>本文</b></p>")).AddFileAttachment("document.pdf")

	msg := master.To("宛先", "hanako@example.com")
	err := msg.Send(sender)
	if err != nil {
		log.Fatal(err)
	}
}
