package drivermandrill

import (
	"os"

	"github.com/dwarvesf/x/emailer"
	"github.com/dwarvesf/x/emailer/mandrill"
)

func init() {
	if os.Getenv("MANDRILL_KEY") == "" || os.Getenv("MANDRILL_USER") == "" {
		panic("failed to load env for mandrill")
	}

	emailer.Register("mandrill", &mandrill.Mandrill{
		os.Getenv("MANDRILL_KEY"),
		os.Getenv("MANDRILL_USER"),
	})
}
