package main

import (
	"fmt"
	"math"
	"time"
)

var totalPrimeNumbers = 0

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

func main() {
	// runtime.GOMAXPROCS(3)
	fmt.Println("Counting prime numbers between 1 and 100000000")
	startTime := time.Now()
	for i := 1; i <= 100000000; i++ {
		if isPrime(i) {
			totalPrimeNumbers++
		}
	}

	fmt.Printf("Total prime numbers between 1 and 10000000: %d  calculated in %v second\n", totalPrimeNumbers, time.Since(startTime))
	// fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	// fmt.Println(runtime.NumCPU())
}
