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

func TestPrimeCache(t *testing.T) {
	assert.True(t, Prime(11))
	assert.True(t, Prime(11))
}

func Test_Nth_Prime(t *testing.T) {
	assert.Equal(t, 104743, NthPrime(10001))
}

func Test_LastestProductSeries(t *testing.T) {
	assert.Equal(t, 23514624000, LastestProductSeries(13))
}

func Test_Sum_of_Prime(t *testing.T) {
	assert.Equal(t, 142913828922, SumOfPrime(2000000))
}

func Test_Sum_of_Prime_Sieve(t *testing.T) {
	assert.Equal(t, 142913828922, SumOfPrimeSieve(2000000))
}

func TestSpecialPythagoreanTriple(t *testing.T) {
	assert.Equal(t, 31875000, SpecialPythagoreanTriple(1000))
}

func TestLargestProductInGrid(t *testing.T) {
	assert.Equal(t, int64(70600674), LargestProductInGrid())
}

func TestHighlyDividedTriangularNumber(t *testing.T) {
	assert.Equal(t, 76576500, HighlyDividedTriangularNumber())
}

// func TestHighlyDividedTriangularNumber(t *testing.T) {
// 	assert.Equal(t, 0, FactorsAmt(28, 5))
// }

func TestLargeSumNumber(t *testing.T) {
	assert.Equal(t, "5537376230", LargeSum())
}

func TestLongestCollatzSequence(t *testing.T) {
	assert.Equal(t, 837799, LongestCollatzSequence(1000000))
}

func TestLatticePaths(t *testing.T) {
	assert.Equal(t, 137846528820, LatticePaths())
}

func TestPowerDigitDum(t *testing.T) {
	assert.Equal(t, int64(1366), PowerDigitSum())
}

func TestLetterCount(t *testing.T) {
	assert.Equal(t, 11, LetterCount(1000))
}

func TestLetterCount1(t *testing.T) {
	assert.Equal(t, 10, LetterCount(100))
}

func TestLetterCount2(t *testing.T) {
	assert.Equal(t, 10, LetterCount(200))
}

func TestLetterCount3(t *testing.T) {
	assert.Equal(t, 5, LetterCount(40))
}

func TestLetterCount4(t *testing.T) {
	assert.Equal(t, 8, LetterCount(42))
}

func TestNumberLetterCounts(t *testing.T) {
	assert.Equal(t, 21124, NumberLetterCounts())
}

func TestMaximumPathSum1(t *testing.T) {
	assert.Equal(t, 1074, MaximumPathSum1())
}

func TestCountingSundays(t *testing.T) {
	assert.Equal(t, 171, CountingSundays())
}

func TestFactorialDigitNum(t *testing.T) {
	assert.Equal(t, "648", FactorialDigitNum(100))
}

func TestGetFactorSum(t *testing.T) {
	assert.Equal(t, 284, GetFactorSum(220))
	assert.Equal(t, 220, GetFactorSum(284))
	assert.Equal(t, 1, GetFactorSum(2))
	assert.Equal(t, 16, GetFactorSum(12))
	assert.Equal(t, 3, GetFactorSum(4))
	assert.Equal(t, 15, GetFactorSum(16))
}

func TestAmicableNumbers(t *testing.T) {
	assert.Equal(t, 31626, AmicableNumbers(10000))
}
