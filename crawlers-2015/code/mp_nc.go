package main

import (
	crand "crypto/rand"
	"crypto/sha1"
	"fmt"
	"sync"
	"time"
)

const (
	M = 1 << 12 // 4096
	N = 1 << 11 // 2048
	W = 8
)

/*
var (
    wgProduce sync.WaitGroup
    wgConsume sync.WaitGroup
)

// STARTMAIN OMIT
func main() {
    var todo = make(chan string)

    wgProduce.Add(M)
    wgConsume.Add(N)

    for i := 0; i < M; i++ {
        go produce(todo)
    }
    for i := 0; i < N; i++ {
        go consume(todo)
    }

    wgProduce.Wait()
    close(todo)
    wgConsume.Wait()

    PrintStats(...)
}
// ENDMAIN OMIT
*/

var (
	wgProduce         sync.WaitGroup
	wgConsume         sync.WaitGroup
	total, failed, ok int
)

func main() {
	var start = time.Now()
	var todo = make(chan string)

	fmt.Printf("Starting %d producers...\n", M)
	fmt.Printf("Starting %d consumers...\n", N)
	fmt.Printf("Launching %d jobs...\n", M*W)

	wgProduce.Add(M)
	wgConsume.Add(N)

	for i := 0; i < M; i++ {
		go produce(todo)
	}
	for i := 0; i < N; i++ {
		go consume(todo)
	}

	wgProduce.Wait()
	close(todo)
	wgConsume.Wait()

	fmt.Println()

	realTotal := M * W
	fmt.Printf("total=%d (%d)\n", total, total-realTotal)
	fmt.Printf("ok=%d (%d)\n", ok, ok-realTotal)
	elapsed := time.Since(start)
	fmt.Println("elapsed", elapsed)
	fmt.Printf("throughput=%.2f work/s\n", float64(realTotal)/elapsed.Seconds())
}

// STARTPRODUCE OMIT
func produce(ch chan string) {
	it := dbFetch(W)
	for it.Next() {
		ch <- it.Value()
	}
	wgProduce.Done()
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
	// pause := time.Duration(10 + rand.Intn(15))
	// time.Sleep(pause * time.Millisecond)
	i.pos++
	h := sha1.New()
	buf := make([]byte, 8)
	crand.Read(buf)
	return fmt.Sprintf("http://tyba.com/%x.png", h.Sum(buf)[:2])
}

func dbFetch(limit int) *iter {
	return &iter{pos: 0, limit: limit}
}

/*
// STARTCONSUME OMIT
// func consumeOld(url string) error { .. }

func consume(todo chan string) {
	for url := range todo {
		total++ // HL
		if err := doConsume(url); err != nil {
			failed++ // HL
			continue
		}
		ok++ // HL
	}
	wgConsume.Done() // HL
}
// ENDCONSUME OMIT
*/

func consume(todo chan string) {
	for url := range todo {
		// fmt.Print(".")
		total++
		if err := doConsume(url); err != nil {
			failed++
			continue
		}
		ok++
	}
	wgConsume.Done()
}

func doConsume(url string) error {
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
	// if rand.Float64() > 0.8 {
	// 	return nil, errors.New("fetch: some error")
	// }
	return nil, nil
}

func process(data []byte) ([]byte, error) {
	// if rand.Float64() > 0.9 {
	// 	return nil, errors.New("process: some error")
	// }
	return nil, nil
}

func store(results []byte) error {
	// if rand.Float64() > 0.95 {
	// 	return errors.New("store: some error")
	// }
	return nil
}
