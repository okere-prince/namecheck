package github_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jub0bs/namecheck"
	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/stub"
)

var _ namecheck.Checker = (*github.GitHub)(nil)

func TestUsernameTooLong(t *testing.T) {
	var gh = github.GitHub{}
	username := "obviously-longer-than-39-chars-skjdhsdkhfkshkfshdkjfhksdjhf"
	want := false
	got := gh.IsValid(username)
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
			label:     "notfound",
			username:  "dummy",
			client:    stub.ClientWithStatusCode(http.StatusNotFound),
			available: true,
		}, {
			label:    "ok",
			username: "dummy",
			client:   stub.ClientWithStatusCode(http.StatusOK),
		}, {
			label:    "other", // other than 200, 404
			username: "dummy",
			client:   stub.ClientWithStatusCode(999),
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
			gh := github.GitHub{
				Client: c.client,
			}
			available, err := gh.IsAvailable(c.username)
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
