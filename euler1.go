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

// OneThousandsDigitFibonacciNumber https://projecteuler.net/problem=25
func OneThousandsDigitFibonacciNumber() int {
	f1 := "1"
	f2 := "1"
	index := 2
	for {
		if len(fmt.Sprint(f2)) == 1000 {
			return index
		}

		f1, f2 = f2, AddTwoBigNumber(f1, f2)
		index++
	}
}

// AddTwoBigNumber supports add big number as string
func AddTwoBigNumber(a, b string) string {
	carry := 0
	l, r := a, b

	result := ""
	if len(b) > len(a) {
		l, r = b, a
	}

	for i := 0; i < len(l)-len(r); i++ {
		r = "0" + r
	}

	for i := len(l) - 1; i >= 0; i-- {
		sum := int(l[i]-'0') + int(r[i]-'0') + carry
		carry = sum / 10

		result = fmt.Sprintf("%d%s", sum%10, result)
	}

	if carry > 0 {
		result = fmt.Sprintf("%d%s", carry, result)
	}

	return result
}

// ReciprocalCycles https://projecteuler.net/problem=26
func ReciprocalCycles(limit int) int {
	max := 0
	value := 0
	for i := 2; i < limit; i++ {
		len := FindCycle(i)
		if max < len {
			max = len
			value = i
		}
	}

	return value

}

// FindCycle finds the length of the recurring cycle of a decimal fraction
func FindCycle(i int) int {
	if i == 1 {
		return 0
	}

	divider := 1

	for {
		if divider > i {
			break
		}
		divider *= 10
	}

	set := map[int]int{}
	count := 0
	for {
		reminder := divider % i
		count++

		if reminder == 0 {
			return 0
		}

		for {
			if reminder > i {
				break
			}
			reminder *= 10
		}

		if value, ok := set[reminder]; ok {
			return count - value
		}

		set[reminder] = count

		divider = reminder
	}
}
