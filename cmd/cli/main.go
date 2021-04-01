package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jub0bs/namecheck"
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
	username := os.Args[1]

	checkers := []namecheck.Checker{
		&twitter.Twitter{
			Client: http.DefaultClient,
		},
		&github.GitHub{
			Client: http.DefaultClient,
		},
	}
	for _, checker := range checkers {
		if !checker.IsValid(username) {
			fmt.Printf("%q is invalid on %s\n", username, checker)
			continue
		}
		avail, err := checker.IsAvailable(username)
		if err != nil {
			fmt.Printf("failed to check the availability of %q on %s\n", username, checker)
			continue
		}
		if !avail {
			fmt.Printf("%q is valid but unavailable on %s\n", username, checker)
			continue
		}
		fmt.Printf("%q is valid and available on %s\n", username, checker)
	}
}
