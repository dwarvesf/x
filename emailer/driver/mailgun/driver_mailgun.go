package drivermailgun

import (
	"os"

	"github.com/dwarvesf/x/emailer"
	"github.com/dwarvesf/x/emailer/mailgun"
)

func init() {
	if os.Getenv("MAILGUN_DOMAIN") == "" || os.Getenv("MAILGUN_PRIVATE") == "" {
		panic("failed to load env for mailgun")
	}

	emailer.Register("mailgun", &mailgun.Mailgun{
		os.Getenv("MAILGUN_DOMAIN"),
		os.Getenv("MAILGUN_PRIVATE"),
	})
}
