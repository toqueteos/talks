package main

import (
	"fmt"
	"time"
)

// STARTMAIN OMIT
func main() {
	ch := make(chan string, 3) // HL
	ch <- "teapot"
	ch <- "pan"
	ch <- "hoven"
	go shouter(ch)                     // HL
	time.Sleep(250 * time.Millisecond) // We need this
	close(ch)
}

func shouter(ch chan string) {
	for thing := range ch {
		fmt.Printf("I'm a %s!\n", thing)
	}
}

// ENDMAIN OMIT
