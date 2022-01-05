package projecteuler

import (
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
