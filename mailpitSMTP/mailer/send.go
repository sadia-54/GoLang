package mailer

import (
	"fmt"

	"github.com/emersion/go-smtp"
)

func SendByGoSMTPPackage(m *Message) error {

	c, err := smtp.Dial(fmt.Sprintf("%s:%d", m.Host, m.Port))
	if err != nil {
		return err
	}

	if err := c.Mail(m.Sender, nil); err != nil {
		return err
	}

	if err := c.Rcpt(m.Receiver, nil); err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(wc, m.Data)
	if err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return c.Quit()
}
