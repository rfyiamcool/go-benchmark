package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	wg        = sync.WaitGroup{}
	taskqueue chan func()
)

func init() {
	taskqueue = make(chan func(), 100000)
	handlePool()
}

func main() {
	if len(os.Args) == 0 {
		panic("input 0 1 2 mode")
	}

	for index := 0; index < 3; index++ {
		wg.Add(1)
		mode, _ := strconv.Atoi(os.Args[1])
		worker(mode)
	}

	wg.Wait()
}

func handlePool() {
	for index := 0; index < 100; index++ {
		go func() {
			for {
				select {
				case task := <-taskqueue:
					task()
				}
			}
		}()
	}
}

func worker(mode int) {
	nolock := sync.RWMutex{}
	qs := []chan bool{}
	for index := 0; index < 50000; index++ {
		q := make(chan bool, 500)
		qs = append(qs, q)

		go func() {
			for {
				select {
				case <-q:
				}
			}
		}()
	}

	go func() {
		defer wg.Done()
		incr := 0
		for incr < 10 {
			incr++
			nolock.RLock()
			now := time.Now()
			switch mode {
			case 0:
				for _, q := range qs {
					q <- true
				}
			case 1:
				inwg := sync.WaitGroup{}
				chunks := splitChunks(qs, 500)
				for _, chunk := range chunks {
					inwg.Add(1)
					go func(list []chan bool) {
						for _, q := range list {
							q <- true
						}
						inwg.Done()
					}(chunk)
				}

				inwg.Wait()

			case 2:
				inwg := sync.WaitGroup{}
				chunks := splitChunks(qs, 500)

				for _, chunk := range chunks {
					list := chunk
					inwg.Add(1)
					taskqueue <- func() {
						for _, q := range list {
							q <- true
						}
						inwg.Done()
					}
				}
				inwg.Wait()
			}
			end := time.Since(now)

			switch mode {
			case 0:
				fmt.Println("pipeline cost", end.String())
			case 1:
				fmt.Println("concurrent cost", end.String())
			case 2:
				fmt.Println("go pool cost", end.String())
			}

			time.Sleep(200 * time.Millisecond)
			nolock.RUnlock()
		}

	}()
}

func splitChunks(taskList []chan bool, chunkSize int) [][]chan bool {
	var taskChunk [][]chan bool
	for i := 0; i < len(taskList); i += chunkSize {
		end := i + chunkSize

		if end > len(taskList) {
			end = len(taskList)
		}

		taskChunk = append(taskChunk, taskList[i:end])
	}

	return taskChunk
}
