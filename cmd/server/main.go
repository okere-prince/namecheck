package main

import (
	"fmt"
	"sync"
)

func main() {
	printTenIntsConcurrently()
}

func printTenIntsConcurrently() {
	var wg sync.WaitGroup
	const n = 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go printConcurrently(i, &wg)
	}
	wg.Wait()
}

func printConcurrently(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}
