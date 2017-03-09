package drivermailgun

import (
	"os"

	"github.com/dwarvesf/x/emailer"
	"github.com/dwarvesf/x/emailer/sendgrid"
)

func init() {
	if os.Getenv("SENDGRID_API_ID") == "" || os.Getenv("SENDGRID_API_KEY") == "" {
		panic("failed to load env for mailgun")
	}

	emailer.Register("sendgrid", &sendgrid.Sendgrid{
		os.Getenv("SENDGRID_API_ID"),
		os.Getenv("SENDGRID_API_KEY"),
	})
}
