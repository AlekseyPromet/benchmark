package main

import (
	"flag"
	"strconv"
	"sync"
	"time"
)

func main() {

	var n int
	var err error

	if n, err = strconv.Atoi(flag.Arg(1)); err != nil {
		n = 1_000_000
	}

	Run(n)

}

func Run(n int) {
	wg := &sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			time.Sleep(time.Second * 10)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
