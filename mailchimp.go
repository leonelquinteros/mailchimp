package mailchimp

import "os"

var (
	apiHost = "https://example.com/3.0/"
	apiKey  string
)

func init() {
	// Set environment variables configuration.
	if os.Getenv("MAILCHIMP_API_HOST") != "" {
		apiHost = os.Getenv("MAILCHIMP_API_HOST")
	}
	if os.Getenv("MAILCHIMP_API_KEY") != "" {
		apiKey = os.Getenv("MAILCHIMP_API_KEY")
	}
}
