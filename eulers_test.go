package projecteuler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestM3_5(t *testing.T) {
	assert.Equal(t, 233168, M3_5(1000))
}

func Test_fibo(t *testing.T) {
	assert.Equal(t, 13, Fibonacci(7))
}

func Test_fibo_evnt(t *testing.T) {
	assert.Equal(t, 4613732, FibonacciEvent(4000000))
}

func Test_Large_prime_factor(t *testing.T) {
	assert.Equal(t, 6857, LargePrimeFactor(600851475143))
}

func Test_Largest_palindom(t *testing.T) {
	assert.Equal(t, 906609, LargestPalindrome(999, 99))
}

func Test_Smallest_multiples(t *testing.T) {
	assert.Equal(t, 232792560, SmallestMultiple(20))
}

func Test_Sum_square_differences(t *testing.T) {
	assert.Equal(t, 25164150, SumSquareDifference(100))
}

func Test_Sum_square_differences1(t *testing.T) {
	assert.Equal(t, 25164150, SumSquareDifference1(100))
}

func TestPrime(t *testing.T) {
	assert.False(t, Prime(9))
}

func Test_Nth_Prime(t *testing.T) {
	assert.Equal(t, 104743, NthPrime(10001))
}

func Test_LastestProductSeries(t *testing.T) {
	assert.Equal(t, 23514624000, LastestProductSeries(13))
}
