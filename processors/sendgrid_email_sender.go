package processors

import (
	"encoding/json"
	"os"

	"github.com/dailyburn/ratchet/data"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendgridEmailSender is a DataProcessor that sends email via Sendgrid
type SendgridEmailSender struct{}

// NewSendgridEmailSender returns a new *SendgridEmailSender
func NewSendgridEmailSender() *SendgridEmailSender {
	return &SendgridEmailSender{}
}

// ProcessData implementation
func (s *SendgridEmailSender) ProcessData(
	d data.JSON, outputChan chan data.JSON,
	killChan chan error,
) {
	var (
		from             = mail.NewEmail("Example", "test@example.com")
		subject          = "Daily Report"
		to               = mail.NewEmail("Example", "test@example.com")
		plainTextContent = "easy to do anywhere, even with Go"
		htmlContent      = "<strong>easy to do anywhere, even with Go</strong>"

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
