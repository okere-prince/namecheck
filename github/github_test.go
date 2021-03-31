package github_test

import (
	"testing"

	"github.com/jub0bs/namecheck/github"
)

func TestUsernameTooLong(t *testing.T) {
	username := "obviously-longer-than-39-chars-skjdhsdkhfkshkfshdkjfhksdjhf"
	want := false
	got := github.IsValid(username)
	if got != want {
		t.Errorf(
			"IsValid(%s) = %t; want %t",
			username,
			got,
			want,
		)
	}
}
