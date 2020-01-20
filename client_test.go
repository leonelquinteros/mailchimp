package mailchimp

import (
	"log"
	"os"
	"testing"
)

func getTestClient() Client {
	return New(os.Getenv("MAILCHIMP_API_KEY"))
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
