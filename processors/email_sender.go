package processors

import (
	"bytes"
	"net/smtp"

	"github.com/topfreegames/go-etl/models"
)

// EmailSender is a DataProcessor that sends email
type EmailSender struct {
	smtpServer        string
	sender, recipient string
}

// ProcessData implementation
func (e *EmailSender) ProcessData(d models.Data, outputChan chan models.Data, killChan chan error) {
	c, err := smtp.Dial(e.smtpServer)
	if err != nil {
		killChan <- err
		return
	}
	defer c.Close()

	c.Mail(e.sender)
	c.Rcpt(e.recipient)

	wc, err := c.Data()
	if err != nil {
		killChan <- err
		return
	}
	defer wc.Close()

	buf := bytes.NewBuffer(d)
	if _, err = buf.WriteTo(wc); err != nil {
		killChan <- err
		return
	}
}
