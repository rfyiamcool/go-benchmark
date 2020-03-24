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

func dispatchBenchmark(c, buf, sender, receiver int) {
	start := time.Now()

	balanceNum := 2
	chpool := make([]chan bool, 0, balanceNum)
	for i := 0; i < balanceNum; i++ {
		q := make(chan bool, buf)
		chpool = append(chpool, q)
	}

	for index := 0; index < receiver; index++ {
		idx := index
		go func() {
			n := idx % balanceNum
			for {
				<-chpool[n]
			}
		}()
	}

	wg := sync.WaitGroup{}
	step := c / sender

	for i := 0; i < sender; i++ {
		idx := i
		wg.Add(1)
		go func() {
			n := idx % balanceNum
			for j := 0; j < step; j++ {
				chpool[n] <- true
			}
			wg.Done()
		}()
	}

	wg.Wait()

	cost := time.Since(start)
	fmt.Printf("dispatch count: %d, chan buf: %d, sender: %d, recevier: %d, time cost: %s, qps: %.2f \n",
		c, buf, sender, receiver, time.Since(start), float64(count)/cost.Seconds(),
	)
}

func directBenchmark(c, buf, sender, receiver int) {
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

	cost := time.Since(start)
	fmt.Printf("direct  count: %d, chan buf: %d, sender: %d, recevier: %d, time cost: %s, qps: %.2f \n",
		c, buf, sender, receiver, time.Since(start), float64(count)/cost.Seconds(),
	)
}

func main() {
	directBenchmark(count, 0, 10, 10)
	dispatchBenchmark(count, 0, 10, 10)

	directBenchmark(count, 50, 10, 10)
	dispatchBenchmark(count, 50, 10, 10)

	directBenchmark(count, 100, 10, 10)
	dispatchBenchmark(count, 100, 10, 10)

	directBenchmark(count, 1000, 10, 10)
	dispatchBenchmark(count, 1000, 10, 10)
	time.Sleep(1 * time.Second)
	fmt.Println()

	directBenchmark(count, 0, 50, 50)
	dispatchBenchmark(count, 0, 50, 50)
	directBenchmark(count, 100, 50, 50)
	dispatchBenchmark(count, 100, 50, 50)
	directBenchmark(count, 1000, 50, 50)
	dispatchBenchmark(count, 1000, 50, 50)
	time.Sleep(1 * time.Second)
	fmt.Println()

	directBenchmark(count, 0, 100, 100)
	dispatchBenchmark(count, 0, 100, 100)

	directBenchmark(count, 100, 100, 100)
	dispatchBenchmark(count, 100, 100, 100)

	directBenchmark(count, 1000, 100, 100)
	dispatchBenchmark(count, 1000, 100, 100)

	time.Sleep(1 * time.Second)
	fmt.Println()

	directBenchmark(count, 0, 200, 200)
	dispatchBenchmark(count, 0, 200, 200)

	directBenchmark(count, 100, 200, 200)
	dispatchBenchmark(count, 100, 200, 200)

	directBenchmark(count, 1000, 200, 200)
	dispatchBenchmark(count, 1000, 200, 200)
	time.Sleep(1 * time.Second)
	fmt.Println()

	directBenchmark(count, 0, 300, 300)
	dispatchBenchmark(count, 0, 300, 300)

	directBenchmark(count, 100, 300, 300)
	dispatchBenchmark(count, 100, 300, 300)

	directBenchmark(count, 1000, 300, 300)
	dispatchBenchmark(count, 1000, 300, 300)

	time.Sleep(1 * time.Second)
	fmt.Println()
}
