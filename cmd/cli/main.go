package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

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

	var checkers []namecheck.Checker
	for i := 0; i < 3; i++ {
		t := &twitter.Twitter{
			Client: http.DefaultClient,
		}
		g := &github.GitHub{
			Client: http.DefaultClient,
		}
		checkers = append(checkers, t, g)
	}
	var wg sync.WaitGroup
	wg.Add(len(checkers))
	for _, checker := range checkers {
		go check(checker, username, &wg)
	}
	wg.Wait()
}

func check(checker namecheck.Checker, username string, wg *sync.WaitGroup) {
	defer wg.Done()
	if !checker.IsValid(username) {
		fmt.Printf("%q is invalid on %s\n", username, checker)
		return
	}
	avail, err := checker.IsAvailable(username)
	if err != nil {
		fmt.Printf("failed to check the availability of %q on %s\n", username, checker)
		return
	}
	if !avail {
		fmt.Printf("%q is valid but unavailable on %s\n", username, checker)
		return
	}
	fmt.Printf("%q is valid and available on %s\n", username, checker)
}
