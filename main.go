package main

import (
	"fmt"

	"github.com/jub0bs/namecheck/twitter"
)

type Status int

const (
	Unknown Status = iota
	Active
	Suspended
	Available
)

func main() {
	username := "jub0bs"
	fmt.Println(twitter.IsValid(username))
}
