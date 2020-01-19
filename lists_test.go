package mailchimp

import (
	"testing"
)

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

	// Test first list
	if len(r.Lists) > 0 {
		listID := r.Lists[0].ID

		m := ListMember{
			EmailAddress: "leonelquinteros@gmail.com",
			Status:       ListMemberStatusSubscribed,
		}

		// Create member
		_, err := c.Lists().CreateMember(listID, m)
		if err != nil {
			t.Error(err)
		}

		// CreateOrUpdate
		m.StatusIfNew = ListMemberStatusSubscribed
		m.EmailType = "html"
		_, err = c.Lists().CreateOrUpdateMember(listID, m)
		if err != nil {
			t.Error(err)
		}

		// Update
		_, err = c.Lists().UpdateMember(listID, m)
		if err != nil {
			t.Error(err)
		}

		// Delete
		err = c.Lists().DeleteMember(listID, m)
		if err != nil {
			t.Error(err)
		}

	}
}
