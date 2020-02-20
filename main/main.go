package main

import (
	"fmt"
	"math"
)

var primes = make([]int, 0)
var notPrimes = make([]int, 0)

func Contains(s []int, n int) bool {
	for _, val := range s {
		if val == n {
			return true
		}
	}

	return false
}

func ReverseNumber(n int) int {
	r := 0

	for n > 0 {
		r *= 10
		r += n % 10
		n /= 10
	}

	return r
}

func IsPrime(n int) bool {
	if Contains(primes, n) {
		return true
	} else if Contains(notPrimes, n) {
		return false
	}

	for d := 2; d <= int(math.Sqrt(float64(n))); d++ {
		if n%d == 0 {
			notPrimes = append(notPrimes, n)
			return false
		}
	}

	if n > 1 {
		primes = append(primes, n)
	}

	return n > 1
}

func IsBackwardsPrime(n int) bool {
	rn := ReverseNumber(n)
	correct := n > 10 && n != rn && IsPrime(n) && IsPrime(rn)

	return correct
}

func BackwardsPrime(start int, stop int) []int {
	ans := make([]int, 0)

	for ; start <= stop; start++ {
		if IsBackwardsPrime(start) {
			ans = append(ans, start)
		}
	}

	if len(ans) == 0 {
		return nil
	}

	return ans
}

func main() {
	fmt.Println(BackwardsPrime(15100, 18133))
	fmt.Println(BackwardsPrime(9900, 10000))
	fmt.Println(BackwardsPrime(501, 599))
}
