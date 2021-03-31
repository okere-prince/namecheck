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
	valid := []string{}
	invalid := []string{}

	for _, username := range os.Args[1:] {
		if !twitter.IsValid(username) || !github.IsValid(username) {
			invalid = append(invalid, username)
			continue
		}
		valid = append(valid, username)
	}

	fmt.Println("valid:", valid)
	fmt.Println("invalid:", invalid)
}
