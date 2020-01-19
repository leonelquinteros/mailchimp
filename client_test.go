package mailchimp

import (
	"log"
	"os"
	"testing"
)

func getTestClient() Client {
	cc := ClientConfig{
		APIHost: os.Getenv("MAILCHIMP_API_HOST"),
		APIKey:  os.Getenv("MAILCHIMP_API_KEY"),
		Debug:   false,
	}
	return NewClient(cc)
}

func TestClient(t *testing.T) {
	c := getTestClient()
	r, err := c.Index()
	if err != nil {
		t.Fatal(err)
	}

	if c.config.Debug {
		log.Printf("%+v", r)
	}
}
