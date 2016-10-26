package twilio

import (
	"net/url"
	"testing"
)

func TestGetConferencePage(t *testing.T) {
	t.Parallel()
	client, s := getServer(conferencePage)
	defer s.Close()
	data := url.Values{"PageSize": []string{"3"}}
	conferences, err := client.Conferences.GetPage(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(conferences.Conferences) != 3 {
		t.Errorf("expected to get 3 conferences, got %d", len(conferences.Conferences))
	}
}

func TestGetConference(t *testing.T) {
	t.Parallel()
	client, s := getServer(conferenceInstance)
	defer s.Close()
	conference, err := client.Conferences.Get(conferenceInstanceSid)
	if err != nil {
		t.Fatal(err)
	}
	if conference.Sid != conferenceInstanceSid {
		t.Errorf("expected Sid to be %s, got %s", conferenceInstanceSid, conference.Sid)
	}
	if conference.FriendlyName != "testConference" {
		t.Errorf("expected FriendlyName to be 'testConference', got %s", conference.FriendlyName)
	}
}
