package twitter_test

import (
	"testing"

	"github.com/jub0bs/namecheck/twitter"
)

var tw = twitter.Twitter{}

func TestUsernameTooLong(t *testing.T) {
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
