package main

import (
	crand "crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
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

	for i := 0; i < PRODUCERS; i++ {
		go produceN(todo, N)
	}

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
func main() {
    const PRODUCERS = 2 // ideally a flag // HL
    todo := make(chan string) // HL
    for i := 0; i < PRODUCERS; i++ { // HL
        go produceN(todo, 8) // HL
    } // HL

    for url := range todo {
        total++ // global
        if err := consume(url); err != nil {
            failed++ // global
            continue
        }
        ok++ // global
    }

    PrintStats(...)
}
// ENDMAIN OMIT
*/

// STARTPRODUCE OMIT
func produceN(ch chan string, n int) {
	it := dbFetch(n)
	for it.Next() {
		ch <- it.Value()
	}
	close(ch) // HL
}

// ENDPRODUCE OMIT

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

// STARTCONSUME OMIT
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

// ENDCONSUME OMIT

func fetch(url string) ([]byte, error) {
	if rand.Float64() > 0.8 {
		return nil, errors.New("fetch: some error")
	}
	return nil, nil
}

func process(data []byte) ([]byte, error) {
	if rand.Float64() > 0.9 {
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
