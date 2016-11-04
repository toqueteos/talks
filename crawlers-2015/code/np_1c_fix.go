package main

import (
	crand "crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	PRODUCERS = 20
	N         = 5
)

func main() {
	var (
		total, failed, ok int
		start             = time.Now()
		todo              = make(chan string)
	)

	wgProduce.Add(PRODUCERS)
	for i := 0; i < PRODUCERS; i++ {
		go produceN(todo, N)
	}

	go func() {
		wgProduce.Wait()
		close(todo)
	}()

	for url := range todo {
		fmt.Print(".")
		total++
		if err := consume(url); err != nil {
			failed++
			continue
		}
		ok++
	}

	fmt.Printf("\ntotal=%d ok=%d failed %d\n", total, ok, failed)
	fmt.Println("elapsed", time.Since(start))
}

/*
// STARTMAIN OMIT
var wgProduce sync.WaitGroup // HL

func main() {
    const PRODUCERS = 6
    wgProduce.Add(PRODUCERS) // HL
    todo := make(chan string)
    for i := 0; i < PRODUCERS; i++ {
        go produceN(todo, 4)
    }

	// We iterate over todo, can't do this after `for .. range todo` block
    go func() { // HL
        wgProduce.Wait() // HL
        close(todo) // HL
    }() // HL

    ...

    PrintStats(...)
}
// ENDMAIN OMIT
*/

var wgProduce sync.WaitGroup // HL

// STARTPRODUCE_FIX OMIT
func produceN(ch chan string, n int) {
	it := dbFetch(n)
	for it.Next() {
		ch <- it.Value()
	}
	wgProduce.Done() // HL
}

// ENDPRODUCE_FIX OMIT

type iter struct {
	pos   int
	limit int
}

func (i *iter) Next() bool {
	return i.pos < i.limit
}

func (i *iter) Value() string {
	pause := time.Duration(10 + rand.Intn(15))
	time.Sleep(pause * time.Millisecond)
	i.pos++
	h := sha1.New()
	buf := make([]byte, 8)
	crand.Read(buf)
	return fmt.Sprintf("http://tyba.com/%x.png", h.Sum(buf)[:2])
}

func dbFetch(limit int) *iter {
	return &iter{pos: 0, limit: limit}
}

func consume(url string) error {
	data, err := fetch(url)
	if err != nil {
		return err
	}
	results, err := process(data)
	if err != nil {
		return err
	}
	return store(results)
}

func fetch(url string) ([]byte, error) {
	if rand.Float64() > 0.95 {
		return nil, errors.New("fetch: some error")
	}
	return nil, nil
}

func process(data []byte) ([]byte, error) {
	if rand.Float64() > 0.95 {
		return nil, errors.New("process: some error")
	}
	return nil, nil
}

func store(results []byte) error {
	if rand.Float64() > 0.95 {
		return errors.New("store: some error")
	}
	return nil
}
