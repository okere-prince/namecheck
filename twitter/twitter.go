package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

func (*Twitter) IsAvailable(username string) (bool, error) {
	const tmpl = "https://europe-west6-namechecker-api.cloudfunctions.net/userlookup?username=%s"
	endpoint := fmt.Sprintf(tmpl, url.QueryEscape(username))
	resp, err := http.Get(endpoint)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("unexpected response from API")
	}
	var dto struct {
		Data interface{} `json:"data"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&dto); err != nil {
		return false, err
	}
	// the absence of a data field in the response body indicates the username's availability
	return dto.Data == nil, nil
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
