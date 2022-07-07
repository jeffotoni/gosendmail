package gosendmail

import (
	"os"
)

var (
	host         = os.Getenv("EMAIL_HOST")
	username     = os.Getenv("EMAIL_USERNAME")
	password     = os.Getenv("EMAIL_PASSWORD")
	port         = os.Getenv("EMAIL_PORT")
	insecureSkip = os.Getenv("EMAIL_INSECURE")
	from         = os.Getenv("EMAIL_FROM")
)
