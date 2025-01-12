package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

const concurrency = 10

var totalPrimeNumbers int32 = 0

func isPrime(n int) bool {
	if n&1 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func doBatch(wg *sync.WaitGroup, start int, end int) {
	defer wg.Done()
	startTime := time.Now()

	for i := start; i < end; i++ {
		if isPrime(i) {
			// totalPrimeNumbers++ // this is not thread safe operation and will cause deadlock
			atomic.AddInt32(&totalPrimeNumbers, 1)
		}
	}

	fmt.Println("Counting prime numbers between", start, "and", end-1, "took", time.Since(startTime).Seconds())
}

func main() {
	fmt.Println("Counting prime numbers between 1 and 100000000")
	// runtime.GOMAXPROCS(3) this makes sure only 3 out of 12 available cores are used. I say 12 because runtime.NumCPU() returns 12 on my machine
	startTime := time.Now()
	batchSize := 100000000 / concurrency
	nstart := 1
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		// countPrimes(i)
		wg.Add(1)
		go doBatch(&wg, nstart, nstart+batchSize)
		nstart += batchSize
	}

	wg.Wait()
	fmt.Println("Total prime numbers found:", totalPrimeNumbers, " in ", time.Since(startTime))
}
