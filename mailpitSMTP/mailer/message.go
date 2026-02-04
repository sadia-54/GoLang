package mailer

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"text/template"

	"mailpitSMTP/configs"
)

type Message struct {
	Sender   string
	Receiver string
	Host     string
	Port     int
	Data     string
}

type MessageBuilder struct {
	template     *template.Template
	templateData map[string]string
	content      string
	header       map[string]string
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		header: map[string]string{},
	}
}

func (b *MessageBuilder) UseTemplate(name string, data map[string]string) *MessageBuilder {
	cfg := configs.GetConfig()
	templateFile := path.Join(cfg.BasePath, "templates", name)

	emailTemplate, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	b.template = emailTemplate
	b.templateData = data
	b.header["Content-Type"] = "text/html"
	return b
}

func (b *MessageBuilder) AddHeader(key, value string) *MessageBuilder {
	b.header[key] = value
	return b
}

func (b *MessageBuilder) SetContent(rawContent string) *MessageBuilder {
	b.content = rawContent
	return b
}

func (b *MessageBuilder) Build() ([]byte, error) {
	if b.content == "" && b.template == nil {
		return nil, errors.New("no body set")
	}

	contentType := b.header["Content-Type"]
	if contentType == "" {
		contentType = "text/plain"
	}

	headers :=
		fmt.Sprintf("From: %s\r\n", b.header["from"]) +
			fmt.Sprintf("To: %s\r\n", b.header["to"]) +
			fmt.Sprintf("Subject: %s\r\n", b.header["subject"]) +
			"MIME-Version: 1.0\r\n" +
			fmt.Sprintf("Content-Type: %s; charset=UTF-8\r\n", contentType) +
			"\r\n"

	var body bytes.Buffer
	body.Write([]byte(headers))

	if b.template != nil {
		if err := b.template.Execute(&body, b.templateData); err != nil {
			return nil, err
		}
		return body.Bytes(), nil
	}

	body.Write([]byte(b.content))
	return body.Bytes(), nil
}
