package github

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/jub0bs/namecheck"
)

type GitHub struct {
	Client namecheck.Client
}

const (
	minLen         = 1
	maxLen         = 39
	illegalPrefix  = "-"
	illegalSuffix  = "-"
	illegalPattern = "--"
)

func New() *GitHub {
	return &GitHub{
		Client: http.DefaultClient,
	}
}

var legalPattern = regexp.MustCompile("^[-0-9A-Za-z]*$")

func (*GitHub) String() string {
	return "GitHub"
}

func (*GitHub) IsValid(username string) bool {
	return isLongEnough(username) &&
		isShortEnough(username) &&
		containsNoIllegalPattern(username) &&
		containsOnlyLegalChars(username) &&
		containsNoIllegalPrefix(username) &&
		containsNoIllegalSuffix(username)
}

func (gh *GitHub) IsAvailable(username string) (bool, error) {
	resp, err := gh.Client.Get("https://github.com/" + url.PathEscape(username))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusNotFound, nil
}

func isLongEnough(username string) bool {
	return utf8.RuneCountInString(username) >= minLen
}

func isShortEnough(username string) bool {
	return utf8.RuneCountInString(username) <= maxLen
}

func containsNoIllegalPattern(username string) bool {
	return !strings.Contains(username, illegalPattern)
}

func containsOnlyLegalChars(username string) bool {
	return legalPattern.MatchString(username)
}

func containsNoIllegalPrefix(username string) bool {
	return !strings.HasPrefix(username, illegalPrefix)
}

func containsNoIllegalSuffix(username string) bool {
	return !strings.HasSuffix(username, illegalSuffix)
}
