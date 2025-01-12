package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var concurrency int = 15
var currentNum int32 = 2
var totalPrimeNumbers int32 = 1
var mu sync.Mutex
var INT_MAX int = 100000000

func isPrime(n int) bool {
	// fmt.Println("Checking if", n, "is prime")
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

func doWork(threadSerialNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	for {

		x := atomic.AddInt32(&currentNum, 1)
		if x > int32(INT_MAX) {
			break
		}

		if isPrime(int(x)) {
			atomic.AddInt32(&totalPrimeNumbers, 1)
		}
	}

	fmt.Println("Thread", threadSerialNumber, "done in ", time.Since(startTime).Seconds())
}

func main() {
	// fmt.Println("This is a concurrent optimized version of prime_numbers.go")
	fmt.Println("Counting prime numbers between 1 and 100000000")
	startTime := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go doWork(i, &wg)
	}

	wg.Wait()
	fmt.Println("Total prime numbers found:", totalPrimeNumbers, " in ", time.Since(startTime))
}
