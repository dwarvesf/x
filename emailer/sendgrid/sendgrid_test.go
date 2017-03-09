package sendgrid_test

import (
	"testing"

	"github.com/dwarvesf/x/emailer"
	_ "github.com/dwarvesf/x/emailer/driver/sendgrid"
)

func TestSendMailSendgrid(t *testing.T) {
	sendgrid, err := emailer.UseProvider("sendgrid")
	if err != nil {
		t.Fatal(err)
	}
	var message = &emailer.Message{
		To:      "lnthach2110@gmail.com",
		From:    "thach@dwarvesf.com",
		Subject: "test",
		Text:    "test",
	}
	err = sendgrid.Send(message)
	if err != nil {
		t.Fatal(err)
	}
}
