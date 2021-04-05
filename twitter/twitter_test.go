package twitter_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jub0bs/namecheck"
	"github.com/jub0bs/namecheck/stub"
	"github.com/jub0bs/namecheck/twitter"
)

var _ namecheck.Checker = (*twitter.Twitter)(nil)

func TestUsernameTooLong(t *testing.T) {
	tw := twitter.Twitter{}
	username := "obviously_longer_than_15_chars"
	want := false
	got := tw.IsValid(username)
	if got != want {
		t.Errorf(
			"IsValid(%s) = %t; want %t",
			username,
			got,
			want,
		)
	}
}

func TestIsAvailable(t *testing.T) {
	cases := []struct {
		label         string
		username      string
		client        namecheck.Client
		available     bool
		errorOccurred bool
	}{
		{
			label:     "200 OK with empty JSON",
			username:  "dummy",
			client:    stub.ClientWithStatusCodeAndBody(http.StatusOK, "{}"),
			available: true,
		}, {
			label:    "200 OK with JSON body containing data field",
			username: "dummy",
			client:   stub.ClientWithStatusCodeAndBody(http.StatusOK, `{"data": "whatever"}`),
		}, {
			label:         "other than 200 OK",
			username:      "dummy",
			client:        stub.ClientWithStatusCodeAndBody(999, ""),
			available:     false,
			errorOccurred: true,
		}, {
			label:         "clienterror",
			username:      "dummy",
			client:        stub.ClientWithError(errors.New("some network error")),
			available:     false,
			errorOccurred: true,
		},
	}

	const tmpl = "IsAvailable(%q): got %t (and %s error); want %t (and %s error)"
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			tw := twitter.Twitter{
				Client: c.client,
			}
			available, err := tw.IsAvailable(c.username)
			if available != c.available || (err != nil != c.errorOccurred) {
				t.Errorf(
					tmpl,
					c.username,
					available,
					errorMsgHelper(err != nil),
					c.available,
					errorMsgHelper(c.errorOccurred))
			}
		})
	}
}

func errorMsgHelper(errorOccurred bool) string {
	if errorOccurred {
		return "some"
	}
	return "no"
}
