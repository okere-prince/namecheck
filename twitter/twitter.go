package twitter

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Twitter struct{}

const (
	minLen         = 1
	maxLen         = 15
	illegalPattern = "twitter"
)

var legalPattern = regexp.MustCompile("^[0-9A-Z_a-z]*$")

func (*Twitter) String() string {
	return "Twitter"
}

func (*Twitter) IsValid(username string) bool {
	return isLongEnough(username) &&
		isShortEnough(username) &&
		containsNoIllegalPattern(username) &&
		containsOnlyLegalChars(username)
}

func isLongEnough(username string) bool {
	return utf8.RuneCountInString(username) >= minLen
}

func isShortEnough(username string) bool {
	return utf8.RuneCountInString(username) <= maxLen
}

func containsNoIllegalPattern(username string) bool {
	return !strings.Contains(username, strings.ToLower(illegalPattern))
}

func containsOnlyLegalChars(username string) bool {
	return legalPattern.MatchString(username)
}
