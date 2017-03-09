package mandrill_test

import (
	"testing"

	"github.com/dwarvesf/x/emailer"
	_ "github.com/dwarvesf/x/emailer/driver/mandrill"
)

func TestSendMailMandrill(t *testing.T) {
	mandrill, err := emailer.UseProvider("mandrill")
	if err != nil {
		t.Fatal(err)
	}
	var message = &emailer.Message{
		To:      "lnthach2110@gmail.com",
		From:    "thach@dwarvesf.com",
		Subject: "test",
		Text:    "test",
	}
	err = mandrill.Send(message)
	if err != nil {
		t.Fatal(err)
	}
}
