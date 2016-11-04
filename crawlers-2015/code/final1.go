package main

import (
	"sync"
	"time"

	"gopkg.in/inconshreveable/log15.v2"
)

// STARTDEF OMIT
type Consumer struct {
	sync.WaitGroup
	// config, queue conn, stats, channels: todo/done, resume/pause, exit...
}

// ENDDEF OMIT

// STARTRUN OMIT
func (c *Consumer) Run() error {
	c.Add(c.NumWorkers) // HL
	for idx := 0; idx < c.NumWorkers; idx++ {
		w := NewWorker(c.todo)
		go eat(idx) // HL
		c.workers[idx] = w
	}
	go c.feed() // HL

	go func() {
		c.Wait() // HL
		close(todo)
	}()

	c.loop() // HL
}

// ENDRUN OMIT

// STARTEAT OMIT
func (c *Consumer) eat(idx int) {
	defer c.Done() // HL
	results := c.workers[idx].Run()
	for result := range results {
		c.consume(result) // HL
	}
}

// ENDEAT OMIT

// STARTCONSUME OMIT
func (c *Consumer) consume(result Result) {
	if c.checkErrors(result) {
		return
	}

	if !c.saveToDatabase(result) {
		return
	}

	c.Stats.Processed++
	c.showConsumeStats(result)
}

// ENDCONSUME OMIT

// STARTFEED OMIT
func (c *Consumer) feed() {
	var stop bool
	for !stop {
		var tick <-chan time.Time
		if !c.IsPaused() { // HL
			tick = c.feedTick
		}

		select {
		case <-tick:
			c.todo <- c.queuePop()
		case <-time.After(c.Options.FeedTimeout):
			if tick != nil {
				stop = true
			}
		}
	}

	log15.Warn("Stopped reading from queue.")

	c.Stop()
	// Consumer.loop now waits for all workers to finish their work
}

// ENDFEED OMIT

// STARTISPAUSED OMIT
func (c *Consumer) IsPaused() {
	resp := make(chan bool)
	c.pausedQuery <- resp
	return <-resp
}

// ENDISPAUSED OMIT

// STARTDL OMIT
func download(todo chan string) {
	var paused bool
	for {
		var todo chan string // HL
		if !paused {         // HL
			todo = c.todo // HL
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

// ENDDL OMIT

// STARTLOOP OMIT
// loop ensures there are no data races
// all mutable data is queried through channels
func (c *Consumer) loop() {
	var stop bool
	for !stop {
		select {
		case <-c.status: // ...
		case <-c.pause: // ...
		case <-c.resume: // ...
		case <-c.stop: // ...
		case resp := <-c.pausedQuery: // ...
		case <-c.exit:
			stop = true
		}
	}
}

// ENDLOOP OMIT
