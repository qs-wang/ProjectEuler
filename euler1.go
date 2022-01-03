package projecteuler

import (
	"fmt"
	"sort"
	"strings"
)

// NonAbundantSums https://projecteuler.net/problem=23
func NonAbundantSums() int {
	abundantNums := []int{}
	for i := 1; i <= 28123; i++ {
		if GetFactorSum(i) > i {
			abundantNums = append(abundantNums, i)
		}
	}

	set := make([]int, 28124)
	for i := 0; i < len(abundantNums); i++ {
		for j := i; j < len(abundantNums); j++ {
			x := abundantNums[i] + abundantNums[j]
			if x <= 28123 {
				set[x] = x
			}
		}
	}

	sum := 0
	for i := 1; i <= 28123; i++ {
		if set[i] == 0 {
			sum += i
		}
	}

	return sum
}

// Find1MillionthLexicographicPermutation https://projecteuler.net/problem=24
func Find1MillionthLexicographicPermutation(nums []int) string {
	count := 1
	next := nums

	for {
		if count == 1000000 {
			break
		}
		next = NextLexicographicPermutations(next)
		count++
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(next)), ""), "[]")
}

// NextLexicographicPermutations finds the next lexicographic permutation of the given slice of ints.
func NextLexicographicPermutations(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return nums
	}

	var k int
	for i := len(nums) - 1; i >= 0; i-- {
		if i == 1 && nums[1] < nums[0] {
			// nums[len(nums)-1], nums[len(nums)-2] = nums[len(nums)-2], nums[len(nums)-1]
			return nums
		}

		if nums[i-1] < nums[i] {
			k = i - 1
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[k] {
			nums[i], nums[k] = nums[k], nums[i]

			sort.Ints(nums[k+1:])
			break
		}
	}

	return nums
}
