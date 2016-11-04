package main

import "fmt"

// STARTMAIN OMIT
func main() {
	ch := make(chan string) // HL
	go shouter(ch)
	ch <- "teapot"
	ch <- "pan"
	ch <- "hoven"
	close(ch)
}

func shouter(ch chan string) {
	for thing := range ch {
		fmt.Printf("I'm a %s!\n", thing)
	}
}

// ENDMAIN OMIT
