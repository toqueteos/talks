package main

import (
	crand "crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const N = 24

// STARTMAIN OMIT
func main() {
	var (
		total, failed, ok int          // HL
		start             = time.Now() // HL
	)

	for url := range produce(N) {
		fmt.Print(".") // HL
		total++        // HL
		if err := consume(url); err != nil {
			failed++ // HL
			continue
		}
		ok++ // HL
	}

	fmt.Printf("\ntotal=%d ok=%d failed %d\n", total, ok, failed) // HL
	fmt.Println("elapsed", time.Since(start))                     // HL
}

// ENDMAIN OMIT

var ErrReasons = errors.New("some error")

func produce(n int) chan string {
	ch := make(chan string)
	go func() {
		it := dbFetch(n)
		for it.Next() {
			ch <- it.Value()
		}
		close(ch)
	}()
	return ch
}

type iter struct {
	pos   int
	limit int
}

func (i *iter) Next() bool {
	return i.pos < i.limit
}

func (i *iter) Value() string {
	pause := time.Duration(50 + rand.Intn(50))
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
