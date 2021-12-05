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
