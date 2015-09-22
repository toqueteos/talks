package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("I'm a spoon!")
	go shoutLater(2*time.Second, "fork")
	go shoutLater(3*time.Second, "teapot")
	time.Sleep(4 * time.Second)
}

func shoutLater(wait time.Duration, thing string) {
	time.Sleep(wait)
	fmt.Printf("(%s later) I'm a %s!\n", wait, thing)
}
