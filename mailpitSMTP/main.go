package main

import (
	"fmt"
	"mailpitSMTP/mailer"

)

func main() {

	emailCfg := struct {
		Host   string
		Port   int
		Sender string
	}{
		Host:   "localhost",
		Port:   1025,
		Sender: "noreply@test.com",
	}

	receiver := "user@test.com"

	// PLAIN TEXT EMAIL
	textBuilder := mailer.NewMessageBuilder().
		AddHeader("from", emailCfg.Sender).
		AddHeader("to", receiver).
		AddHeader("subject", "Plain Text Email").
		AddHeader("Content-Type", "text/plain").
		SetContent("Hello!\nThis is a plain text email sent via Mailpit.")

	textData, err := textBuilder.Build()
	if err != nil {
		panic(err)
	}

	err = mailer.SendByGoSMTPPackage(&mailer.Message{
		Host:     emailCfg.Host,
		Port:     emailCfg.Port,
		Sender:   emailCfg.Sender,
		Receiver: receiver,
		Data:     string(textData),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Plain text email sent")

	// HTML EMAIL
	htmlBuilder := mailer.NewMessageBuilder().
		AddHeader("from", emailCfg.Sender).
		AddHeader("to", receiver).
		AddHeader("subject", "HTML Welcome Email").
		UseTemplate("welcome.html", map[string]string{
			"Name": "Sadia",
		})

	htmlData, err := htmlBuilder.Build()
	if err != nil {
		panic(err)
	}

	err = mailer.SendByGoSMTPPackage(&mailer.Message{
		Host:     emailCfg.Host,
		Port:     emailCfg.Port,
		Sender:   emailCfg.Sender,
		Receiver: receiver,
		Data:     string(htmlData),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("HTML email sent")
}
