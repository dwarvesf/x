package mandrill

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/dwarvesf/x/emailer"
)

var (
	MandrillEndpoint = "https://mandrillapp.com/api/1.0/messages/send.json"
)

// Mandrill ...
type Mandrill struct {
	Key  string
	User string
}

// Message represents the message payload sent to the API
type Message struct {
	// the full HTML content to be sent
	HTML      string `json:"html,omitempty"`
	Text      string `json:"text,omitempty"`
	Subject   string `json:"subject,omitempty"`
	FromEmail string `json:"from_email,omitempty"`
	FromName  string `json:"from_name,omitempty"`
	To        []To   `json:"to"`
}

// To is a single recipient's information.
type To struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
	// the header type to use for the recipient, defaults to "to" if not provided
	// oneof(to, cc, bcc)
	Type string `json:"type,omitempty"`
}

// Send sends the email with format
func (m *Mandrill) Send(message *emailer.Message) error {
	type Body struct {
		Key     string  `json:"key"`
		Message Message `json:"message"`
	}

	var body = &Body{
		Key: m.Key,
		Message: Message{
			Text:      message.Text,
			Subject:   message.Subject,
			FromEmail: message.From,
			To: []To{
				{
					Email: message.To,
				},
			},
		},
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", MandrillEndpoint, bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	_, err = client.Do(req)

	return nil
}
