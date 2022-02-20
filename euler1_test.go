package projecteuler

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonAbundantSums(t *testing.T) {
	assert.Equal(t, 4179871, NonAbundantSums())
}

func TestLexicographicPermutations(t *testing.T) {
	assert.Equal(t, []int([]int{2, 1}), NextLexicographicPermutations([]int{1, 2}))
}

func TestLexicographicPermutations1(t *testing.T) {
	assert.Equal(t, []int([]int{0, 2, 1}), NextLexicographicPermutations([]int{0, 1, 2}))
}

func TestLexicographicPermutations2(t *testing.T) {
	assert.Equal(t, []int([]int{0, 1, 2, 3, 4, 5, 6, 7, 9, 8}), NextLexicographicPermutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func TestLexicographicPermutations3(t *testing.T) {
	assert.Equal(t, []int([]int{0, 1, 3, 0, 2, 3, 5}), NextLexicographicPermutations([]int{0, 1, 2, 5, 3, 3, 0}))
}

func TestLexicographicPermutations4(t *testing.T) {
	assert.Equal(t, []int([]int{1, 0, 2, 3, 4, 5, 6, 7, 8, 9}), NextLexicographicPermutations([]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func TestLexicographicPermutations5(t *testing.T) {
	assert.Equal(t, []int([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), NextLexicographicPermutations([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}

func TestFind1MillionthLexicographicPermutation(t *testing.T) {
	assert.Equal(t, "2783915460", Find1MillionthLexicographicPermutation([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func TestAdd2BigNumber(t *testing.T) {
	assert.Equal(t, "579", AddTwoBigNumber("123", "456"))
}

func TestAdd2BigNumber1(t *testing.T) {
	assert.Equal(t, "4691", AddTwoBigNumber("123", "4568"))
}

func TestAdd2BigNumber2(t *testing.T) {
	assert.Equal(t, "10001", AddTwoBigNumber("433", "9568"))
}

func TestOneThousandsDigitFibonacciNumber(t *testing.T) {
	assert.Equal(t, 4782, OneThousandsDigitFibonacciNumber())
}

func TestFindCycle(t *testing.T) {
	assert.Equal(t, 0, FindCycle(2))
}

func TestFindCycle1(t *testing.T) {
	assert.Equal(t, 1, FindCycle(3))
}

func TestFindCycle2(t *testing.T) {
	assert.Equal(t, 1, FindCycle(6))
}

func TestFindCycle3(t *testing.T) {
	assert.Equal(t, 6, FindCycle(7))
}

func TestReciprocalCycles(t *testing.T) {
	assert.Equal(t, 983, int(ReciprocalCycles(1000)))
}

func TestQuadraticPrimes(t *testing.T) {
	_, p := QuadraticPrimes()

	assert.Equal(t, -59231, p)
}

func TestNumberSpiralSiagonals(t *testing.T) {
	assert.Equal(t, 669171001, NumberSpiralSiagonals(1001))
}

func TestDistinctPowers(t *testing.T) {
	assert.Equal(t, 9183, DistinctPowers(100))
}

func TestDigitFifthPowers(t *testing.T) {
	assert.Equal(t, 443839, DigitFifthPowers())
}

func TestSumDigitalsPow(t *testing.T) {
	assert.Equal(t, 1634, SumDigitalsPow(1634, 4))
}

// func TestCoinSums(t *testing.T) {
// 	assert.Equal(t, 1634, CoinSums())
// }

func TestOneThroughNine(t *testing.T) {
	assert.True(t, OneThroughNine(12, 34, 56789))
}

func TestOneThroughNine1(t *testing.T) {
	assert.False(t, OneThroughNine(12, 30, 56789))
}

func TestOneThroughNine2(t *testing.T) {
	assert.False(t, OneThroughNine(12, 3, 56789))
}

func TestPandigitalProduct(t *testing.T) {
	assert.Equal(t, 45228, PandigitalProduct())
}

func TestDigitCancellingFractions(t *testing.T) {
	DigitCancellingFractions()
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, 120, Factorial(5))
}

func TestDigitFactorials(t *testing.T) {
	assert.Equal(t, 40730, DigitFactorials())
}

func TestRotate(t *testing.T) {
	assert.Equal(t, 121, RotateRight(211))
}

func TestCircularPrime(t *testing.T) {
	assert.Equal(t, 55, CircularPrimes())
}

func TestRotatePrime(t *testing.T) {
	assert.True(t, RotatePrime(719))
}

func TestRotatePrime1(t *testing.T) {
	assert.True(t, RotatePrime(97))
}

func TestToBinary(t *testing.T) {
	assert.Equal(t, "1001001001", ToBinary(585))
}

func TestDoubleBasePalindromes(t *testing.T) {
	assert.Equal(t, 872187, DoubleBasePalindromes())
}

func TestRemoveLeadingDecimalNum(t *testing.T) {
	assert.Equal(t, 797, RemoveLeadingDecimalNum(3797))
}

func TestRemoveLeadingDecimalNum1(t *testing.T) {
	assert.Equal(t, 0, RemoveLeadingDecimalNum(3))
}

func TestRemoveTrailingDecimalNum(t *testing.T) {
	assert.Equal(t, 379, RemoveTrailingDecimalNum(3797))
}

func TestTruncatablePrimes(t *testing.T) {
	assert.Equal(t, 748317, TruncatablePrimes())
}

func TestPrimeLefRight(t *testing.T) {
	assert.True(t, PrimeLefRight(3797))
}

func TestPandigitalMultiples(t *testing.T) {
	assert.Equal(t, 932718654, PandigitalMultiples())
}

func TestIntegerRightRriangles(t *testing.T) {
	assert.Equal(t, 840, IntegerRightRriangles())
}

func TestPandigitalPrime(t *testing.T) {
	assert.Equal(t, 7652413, PandigitalPrime())
}

func TestIsPanDigital(t *testing.T) {
	assert.True(t, isPanDigital(1234))
}

func TestCodedTriangleNumbers(t *testing.T) {
	r, err := CodedTriangleNumbers()
	if err != nil {
		assert.Fail(t, "failed")
	}
	assert.Equal(t, 162, r)
}

func TestIsProperty(t *testing.T) {
	assert.True(t, isProperty(1406357289))
}

func TestIsNotProperty(t *testing.T) {
	assert.False(t, isProperty(1506357299))
}

func TestSubStringSivisibility(t *testing.T) {
	assert.Equal(t, 16695334890, SubStringDivisibility())
}

func TestChampernowneConst(t *testing.T) {
	i := ChampernowneConst(1)

	assert.Equal(t, 1, i)
}

func TestChampernowneConst12(t *testing.T) {
	i := ChampernowneConst(12)

	assert.Equal(t, 1, i)
}

func TestChampernowneConst15(t *testing.T) {
	i := ChampernowneConst(15)

	assert.Equal(t, 2, i)
}

func TestChampernowneConst10(t *testing.T) {
	i := ChampernowneConst(10)

	assert.Equal(t, 1, i)
}

func TestChampernowneConst100(t *testing.T) {
	i := ChampernowneConst(100)

	assert.Equal(t, 5, i)
}

func TestChampernowneConst1000(t *testing.T) {
	i := ChampernowneConst(1000)

	assert.Equal(t, 3, i)
}

func TestChampernowneConstP40(t *testing.T) {
	c := 1
	for i := 0; i < 7; i++ {
		c *= ChampernowneConst(int(math.Pow10(i)))
	}

	assert.Equal(t, 210, c)

}

func TestPentagonNumbers(t *testing.T) {
	assert.Equal(t, 5482660, PentagonNumbers())

}

func TestIsPentagonNumber(t *testing.T) {
	assert.True(t, isPentagonNumber(92))
}

func TestIsPentagonNumberFalse(t *testing.T) {
	assert.False(t, isPentagonNumber(48))
}

func TestTriangularPentagonalHexagonal(t *testing.T) {
	assert.Equal(t, 1533776805, TriangularPentagonalHexagonal(286))
}
