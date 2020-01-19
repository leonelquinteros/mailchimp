package mailchimp

import "testing"

func TestGetLists(t *testing.T) {
	c := getTestClient()
	r, err := c.Lists().All(nil)
	if err != nil {
		t.Fatal(err)
	}

	// Debug
	if c.config.Debug {
		t.Logf("%+v", r)
	}
}
