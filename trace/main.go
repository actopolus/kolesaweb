package main

import (
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	wg := sync.WaitGroup{}

	for i := 0; i < 120; i++ {
		wg.Add(1)

		go func() {
			a := make([]byte, 1000)
			a = append(a, 255)

			time.Sleep(2)

			wg.Done()
		}()
	}

	wg.Wait()
}