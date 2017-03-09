package mailgun_test

import (
	"testing"

	"github.com/dwarvesf/x/emailer"
	_ "github.com/dwarvesf/x/emailer/driver/mailgun"
)

func TestSendMailMailGun(t *testing.T) {
	mailgun, err := emailer.UseProvider("mailgun")
	if err != nil {
		t.Fatal(err)
	}
	var message = &emailer.Message{
		To:      "lnthach2110@gmail.com",
		From:    "thach@dwarvesf.com",
		Subject: "test",
		Text:    "test",
	}
	err = mailgun.Send(message)
	if err != nil {
		t.Fatal(err)
	}
}
