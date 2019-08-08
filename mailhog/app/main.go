package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("hahahaha")
}

func send(body string) {
	from := "...@gmail.com"
	//pass := "..."
	to := "foobarbazz@mailinator.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("localhost:1025", smtp.PlainAuth("", from, "", "localhost"), from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

}
