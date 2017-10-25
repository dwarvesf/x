package sendgrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dwarvesf/x/emailer"
)

var (
	SendgridEndpoint = "https://api.sendgrid.com/v3/mail/send"
)

// Sendgrid ...
type Sendgrid struct {
	APIID  string
	APIKey string
}

type To struct {
	Email string `json:"email"`
}

type From struct {
	Email string `json:"email"`
}

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Personalization struct {
	Tos     []To   `json:"to"`
	Subject string `json:"subject"`
}

type Body struct {
	Personalizations []Personalization `json:"personalizations"`
	From             From              `json:"from"`
	Contents         []Content         `json:"content"`
}

// Send sends the email with format
func (s *Sendgrid) Send(message *emailer.Message) error {

	var body = &Body{
		Personalizations: []Personalization{
			{
				Tos: []To{
					{Email: message.To},
				},
				Subject: message.Subject,
			},
		},
		From: From{message.From},
		Contents: []Content{
			{
				Type:  "text/plain",
				Value: message.Text,
			},
		},
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", SendgridEndpoint, bytes.NewReader(bodyJSON))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	_, err = client.Do(req)

	return err

}
