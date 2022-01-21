package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(300)
	fmt.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	close(done)
}
