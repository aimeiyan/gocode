package main

import (
	"fmt"
	// "os"
	// "math/rand"
	"runtime"
	"syscall"
	"time"
)

// Guess, there are a scheduler, schedule goroutine to run in real threads.

func main() {
	fmt.Println("--------", runtime.GOMAXPROCS(runtime.NumCPU()), runtime.NumCPU())

	for i := 0; i < 10; i++ {
		go func(id int) {
			ch := time.Tick(time.Second * 4)
			for now := range ch {
				// bounce bettween many threads => a goroutine can be run by any real thread
				fmt.Printf("%v pid: %d:%d\n", now, syscall.Gettid(), id)
			}

			// for j := 0; j < 10; j++ {
			// 	fmt.Println("pid is:", )
			// 	time.Sleep(time.Second * 3)
			// }
		}(i)
	}

	time.Sleep(time.Second * 300)
}
