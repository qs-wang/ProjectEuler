package projecteuler

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
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

// QuadraticPrimes https://projecteuler.net/problem=27
func QuadraticPrimes() (int, int) {
	max := 0
	product := 0

	for a := -999; a <= 999; a++ {
		for b := -1000; b <= 1000; b++ {
			n := 0
			for {
				var t = n*n + a*n + b
				if !Prime(t) {
					break
				}
				n++
			}
			if n > max {
				max = n
				product = a * b
			}
		}
	}
	return max, product
}

// NumberSpiralSiagonals https://projecteuler.net/problem=28
func NumberSpiralSiagonals(num int) int {
	sum := 1
	next := 1
	step := 2

	for {
		for i := 0; i < 4; i++ {
			next = next + step
			sum += next
		}
		if next == num*num {
			break
		}
		step += 2
	}

	return sum
}

// DistinctPowers https://projecteuler.net/problem=29
func DistinctPowers(num int) int {
	r := make(map[string]bool)
	for a := 2; a <= num; a++ {
		for b := 2; b <= num; b++ {
			p := new(big.Int).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			if !r[p.String()] {
				r[p.String()] = true
			}
		}
	}

	return len(r)
}

// DigitFifthPowers https://projecteuler.net/problem=16
func DigitFifthPowers() int {
	sum := 0
	for i := 10; ; i++ {
		if i > 354294 {
			break
		}
		temp := SumDigitalsPow(i, 5)

		if temp == i {
			sum += temp
		}

	}

	return sum
}

// SumDigitalsPow sums the digits of a number to the power of a given power
func SumDigitalsPow(i int, p int) int {
	temp := 0
	for {
		r := i % 10
		pow := 1
		for i := 0; i < p; i++ {
			pow *= r
		}
		temp += pow

		if i/10 == 0 {
			break
		}
		i = i / 10
	}
	return temp
}

// CoinSums TBD
func CoinSums() int {
	coins := []float32{1, 2, 5, 10, 20, 50, 100}

	var cnt int = 1

	for i := 0; i < len(coins); i++ {
		findSum(&coins, 200-coins[i], &cnt)
	}

	return cnt
}

func findSum(coins *[]float32, num float32, cnt *int) {

	if num == 0 {
		(*cnt)++
		return
	}

	if num < 0 {
		return
	}

	for i := 0; i < len(*coins); i++ {
		findSum(coins, num-(*coins)[i], cnt)
	}

}

// OneThroughNine identities can be written as a 1 through 9 pandigital
func OneThroughNine(num1, num2, num3 int) bool {
	s1 := strconv.Itoa(num1)
	s2 := strconv.Itoa(num2)
	s3 := strconv.Itoa(num3)

	numSet := make(map[rune]bool)

	t := strings.Join([]string{s1, s2, s3}, "")

	if len(t) < 9 {
		return false
	}

	for _, v := range t {
		if v == '0' {
			return false
		}

		numSet[v] = true

	}

	return len(numSet) == 9
}

//PandigitalProduct https://projecteuler.net/problem=32
func PandigitalProduct() int {
	sum := 0
	for i := 1000; i < 9999; i++ {
		if findFactor(i) {
			sum += i
			continue
		}
	}
	return sum
}

func findFactor(num int) bool {
	for i := 10; i <= 9999; i++ {
		if num%i == 0 {
			x := num / i
			if OneThroughNine(i, x, num) {
				return true
			}
		}
	}

	return false
}

// DigitCancellingFractions https://projecteuler.net/problem=33
func DigitCancellingFractions() {
	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			c := float32(10*a*b) / float32(9*a+b)
			if c == float32(int32(c)) {
				if float32(10*a+b)/(float32(10*b)+c) == float32(a)/c {
					if a != b && b != int(c) {
						fmt.Println(fmt.Sprintf("number is: %d/%d", a, int(c)))
					}
				}
			}
		}
	}
}

// DigitFactorials https://projecteuler.net/problem=34
func DigitFactorials() int {
	max := 241920 // this is the possible max number which equal to 8! * 6

	s := 0
	for i := 3; i < max; i++ {
		r := i
		sum := 0
		for {
			if r == 0 {
				break
			}

			sum += Factorial(r % 10)
			r = r / 10
		}

		if sum == i {
			s += i
		}
	}

	return s
}

// Factorial get factorials of a number
func Factorial(num int) int {
	if num > 21 {
		return -1 //overflow
	}

	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	return result
}

// RotateRight rotate the number to right
func RotateRight(num int) int {
	l := (int)(math.Log10(float64(num)) + 1)
	n := num / 10
	r := num % 10

	return int(float64(r)*math.Pow10(l-1)) + n
}

// CircularPrimes project euler problem 35
func CircularPrimes() int {
	count := 5
	for i := 13; i < 1000000; i++ {
		if Prime(i) {
			if RotatePrime(i) {
				count++
			}
		}
	}
	return count

}

// RotatePrime check if the number is a rotate prime number
// e.g.  197, 971, 719
func RotatePrime(i int) bool {
	l := (int)(math.Log10(float64(i)) + 1)
	r := i
	f := false
	for j := 0; j < l; j++ {
		r = RotateRight(r)
		if !Prime(r) {
			break
		}
		if r == i {
			f = true
		}
	}
	return f
}
