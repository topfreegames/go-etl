package processors

import (
	"encoding/json"
	"os"

	"github.com/dailyburn/ratchet/data"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendgridEmailSender is a DataProcessor that sends email via Sendgrid
type SendgridEmailSender struct {
	FromName  string
	FromEmail string
	ToName    string
	ToEmail   string
	Subject   string
}

// NewSendgridEmailSender returns a new *SendgridEmailSender
func NewSendgridEmailSender(
	fromName, fromEmail, toName, toEmail, subject string,
) *SendgridEmailSender {
	return &SendgridEmailSender{
		FromName:  fromName,
		FromEmail: fromEmail,
		ToName:    toName,
		ToEmail:   toEmail,
		Subject:   subject,
	}
}

// ProcessData implementation
func (s *SendgridEmailSender) ProcessData(
	d data.JSON, outputChan chan data.JSON,
	killChan chan error,
) {
	if len(d) == 0 {
		return
	}

	var (
		from             = mail.NewEmail(s.FromName, s.FromEmail)
		subject          = s.Subject
		to               = mail.NewEmail(s.ToName, s.ToEmail)
		plainTextContent = string(d)
		htmlContent      = string(d)

		message = mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client  = sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	)

	response, err := client.Send(message)
	if err != nil {
		killChan <- err
		return
	}

	bts, err := json.Marshal(response)
	if err != nil {
		killChan <- err
		return
	}

	outputChan <- data.JSON(bts)
}

// Finish implementation
func (s *SendgridEmailSender) Finish(outputChan chan data.JSON, killChan chan error) {}
