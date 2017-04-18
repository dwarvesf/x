package mailgun

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dwarvesf/x/emailer"
)

var (
	MailgunEndpoint = "https://api.mailgun.net/v3/%s/messages"
)

// Maingun ...
type Mailgun struct {
	Domain     string
	PrivateKey string
}

func (m *Mailgun) Send(message *emailer.Message) error {
	req, err := http.NewRequest("POST", fmt.Sprintf(MailgunEndpoint, m.Domain), nil)
	req.Form = url.Values{"to": {message.To}, "from": {message.From}, "subject": {message.Subject}, "text": {message.Text}}
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "multipart/form-data")

	req.SetBasicAuth("api", m.PrivateKey)

	client := &http.Client{}

	_, err = client.Do(req)

	return err
}
