package main

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

func init() {
	debug.SetGCPercent(-1)
}

const count = 10000000

func benchmark(c, buf, sender, receiver int) {
	start := time.Now()
	ch := make(chan bool, buf)

	for index := 0; index < receiver; index++ {
		go func() {
			for {
				<-ch
			}
		}()
	}

	wg := sync.WaitGroup{}
	step := c / sender

	for i := 0; i < sender; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < step; j++ {
				ch <- true
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(ch)

	fmt.Printf("count: %d, chan buf: %d, sender: %d, recevier: %d, time cost: %s \n",
		c, buf, sender, receiver, time.Now().Sub(start),
	)
}

func main() {
	benchmark(count, 0, 10, 10)
	benchmark(count, 100, 10, 10)
	time.Sleep(1 * time.Second)
	fmt.Println()

	benchmark(count, 0, 50, 50)
	benchmark(count, 100, 50, 50)
	time.Sleep(1 * time.Second)
	fmt.Println()

	benchmark(count, 0, 100, 100)
	benchmark(count, 100, 100, 100)
	time.Sleep(1 * time.Second)
	fmt.Println()

	benchmark(count, 0, 200, 200)
	benchmark(count, 100, 200, 200)
	time.Sleep(1 * time.Second)
	fmt.Println()

	benchmark(count, 0, 300, 300)
	benchmark(count, 100, 300, 300)
	time.Sleep(1 * time.Second)
	fmt.Println()
}
