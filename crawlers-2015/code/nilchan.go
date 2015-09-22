package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	var todo = make(chan string)
	go produce(todo)
	work(todo)
}

// STARTMAIN OMIT
func (f *Foo) work() {
	var paused bool
	for {
		var todo chan string // HL
		if !paused {         // HL
			todo = f.todo // HL
		} // HL
		select {
		case item := <-todo: // HL
			go use(item) // HL
		case <-pause:
			paused = true
		case <-resume:
			paused = false
		case <-time.After(5 * time.Second):
			// Avoids waiting for something to happen for a long time
		}
	}
}

// ENDMAIN OMIT

func produce(todo chan string) {
	for {
		h := sha1.New()
		buf := make([]byte, 8)
		crand.Read(buf)
		todo <- fmt.Sprintf("http://tyba.com/%x.png", h.Sum(buf)[:2])
	}
}

func use(item string) {

}
