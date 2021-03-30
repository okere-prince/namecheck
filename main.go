package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Status int

const (
	Unknown Status = iota
	Active
	Suspended
	Available
)

const (
	minLen         = 1
	maxLen         = 15
	illegalPattern = "twitter"
)

func main() {
	username := "jub0bs"
	fmt.Println(isLongEnough(username))
	fmt.Println(isShortEnough(username))
	fmt.Println(containsNoIllegalPattern(username))
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
