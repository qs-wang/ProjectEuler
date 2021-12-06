package projecteuler

import (
	"math"
	"strconv"
)

// M3_5 for https://projecteuler.net/problem=1
func M3_5(num int) int {
	amt := 0
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			amt += i
		}
	}

	return amt
}

// Fibonacci sequence
func Fibonacci(num int) int {
	a, b := 0, 1
	for i := 0; i < num; i++ {
		a, b = b, a+b
	}
	return a
}

//FibonacciEvent even number for https://projecteuler.net/problem=2
func FibonacciEvent(num int) int {
	amt := 0
	a, b, c := 1, 1, 2
	for c <= num {
		amt += c
		a = b + c
		b = c + a
		c = a + b

	}

	return amt
}

// Prime number
func Prime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// LargePrimeFactor for https://projecteuler.net/problem=3
func LargePrimeFactor(num int) int {
	max := 1
	for num%2 == 0 {
		num = num / 2
		max = 2
	}

	for num%3 == 0 {
		num = num / 3
		max = 3
	}

	// only calculate to square root of num
	// if the bigger factor is bigger than square root
	// it should be the remaining part which is handling
	// by the num > 4 case after the loop
	// or we can loop by
	// for i := 5; i*i <= num; i = i + 6
	// and remove the following num>4 if statement
	for i := 5; i*i <= num; i = i + 6 {
		for num%i == 0 {
			num = num / i
			max = i
		}
		for num%(i+2) == 0 {
			num = num / (i + 2)
			max = i + 2
		}
	}

	if num > 4 {
		max = num
	}

	return max
}

// LargestPalindrome https://projecteuler.net/problem=4
func LargestPalindrome(high int, low int) int {
	max := 0
	for i := high; i > low; i-- {
		for j := i; j > low; j-- {
			if isPalindrome(i*j) && i*j > max {
				max = i * j
			}
		}
	}
	return max
}

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

// SmallestMultiple https://projecteuler.net/problem=5
func SmallestMultiple(limit int) int {
	smallest := 1

	for num := 2; num <= limit; num++ {
		square := num
		for square <= limit {
			if smallest%square != 0 {
				smallest = smallest * num
			}
			square = square * num
		}
	}

	return smallest
}
