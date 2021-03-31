package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jub0bs/namecheck/github"
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
	if len(os.Args[1:]) == 0 {
		log.Fatal("username args is required")
	}
	for _, username := range os.Args[1:] {
		const tmpl = "%q is valid on %s: %t\n"
		fmt.Printf(tmpl, username, "Twitter", twitter.IsValid(username))
		fmt.Printf(tmpl, username, "GitHub", github.IsValid(username))
	}
}
