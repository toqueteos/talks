package main

import (
	"fmt"
	"time"
)

// STARTMAIN OMIT
func main() {
	ch := make(chan string)
	go shouter(ch)
	ch <- "teapot"
	close(ch)
	// We are not waiting shouter to be done! // HL
}

func shouter(ch chan string) {
	for thing := range ch {
		prepareShout() // Takes some time // HL
		fmt.Printf("I'm a %s!\n", thing)
	}
}

// ENDMAIN OMIT

func prepareShout() {
	time.Sleep(1 * time.Second)
}
