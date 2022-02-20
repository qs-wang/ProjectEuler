package projecteuler

import (
	"fmt"
	"math"
	"math/big"
	"os"
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

// ToBinary convert a number to binary
func ToBinary(num int) string {
	r := make([]int, 32)

	for i := 0; i < 32; i++ {
		if num&(1<<i) != 0 {
			r[31-i] = 1
		} else {
			r[31-i] = 0
		}
	}

	p := 0
	for i := 0; i < 32; i++ {
		if r[i] == 1 {
			p = i
			break
		}
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r[p:])), ""), "[]")
}

func isPalindromesStr(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeNum(num int) bool {
	s := strconv.Itoa(num)
	return isPalindromesStr(s)
}

// DoubleBasePalindromes project euler problem 36
func DoubleBasePalindromes() int {
	sum := 0

	for i := 1; i < 1000000; i++ {
		if isPalindrome(i) && isPalindromesStr(ToBinary(i)) {
			sum += i
		}
	}
	return sum
}

// RemoveLeadingDecimalNum remove the leading decimal number
func RemoveLeadingDecimalNum(num int) int {
	l := (int)(math.Log10(float64(num)) + 1)

	d := 1
	for i := 0; i < l-1; i++ {
		d = d * 10
	}

	return num % d
}

// RemoveTrailingDecimalNum remove the trailing decimal number
func RemoveTrailingDecimalNum(num int) int {
	return num / 10
}

// TruncatablePrimes project euler problem 37
func TruncatablePrimes() int {
	sum := 0
	c := 0
	for i := 11; ; i++ {
		if Prime(i) && PrimeLefRight(i) {
			sum += i
			c++
		}

		if c == 11 {
			break
		}
	}

	return sum
}

// PrimeLefRight  check if the number is a prime number
// e.g.  197, 971, 719
func PrimeLefRight(num int) bool {
	t := num
	l, r := false, false
	for {
		t = RemoveLeadingDecimalNum(t)
		if t == 0 {
			l = true
			break
		}

		if !Prime(t) {
			break
		}
	}
	t = num
	for {
		t = RemoveTrailingDecimalNum(t)
		if t == 0 {
			r = true
			break
		}

		if !Prime(t) {
			break
		}
	}
	return l && r
}

// PandigitalMultiples project euler problem 38
func PandigitalMultiples() int {
	max := 0
	for i := 1; i < 9999; i++ {
		m := make(map[int]bool)
		sum := strconv.Itoa(i)
		if !checkValidAndPutInMap(i, &m) {
			for j := 2; ; j++ {
				sum += strconv.Itoa(i * j)
				if checkValidAndPutInMap(i*j, &m) {
					break
				}
				if len(m) == 9 {
					s, _ := strconv.Atoi(sum)
					if max < s {
						max = s
					}
				}
			}
		}
	}

	return max
}

func checkValidAndPutInMap(num int, m *map[int]bool) bool {
	r := num
	for {
		if r%10 == 0 {
			return true
		}
		if _, ok := (*m)[r%10]; ok {
			return true
		}
		(*m)[r%10] = true
		r = r / 10
		if r == 0 {
			break
		}
	}
	return false
}

// IntegerRightRriangles project euler problem 39
func IntegerRightRriangles() int {
	m := make(map[int]int)
	// i,j are short edges, as the max perimeter is 1000, so the short edge cannot exceed 500
	for i := 1; i < 500; i++ {
		for j := i; j < 500; j++ {
			k := math.Sqrt(float64(i*i) + float64(j*j))
			if k == math.Trunc(k) {
				l := i + j + int(k)
				if l > 1000 {
					break
				}
				if _, ok := m[l]; ok {
					m[l] = m[l] + 1
				} else {
					m[l] = 1
				}
			}
		}
	}

	max := 0
	ret := 0
	for i, v := range m {
		if max < v {
			max = v
			ret = i
		}
	}

	return ret
}

func isPanDigital(num int) bool {
	l := len(strconv.Itoa(num))

	a := make([]int, l+1)
	n := num

	for {
		r := n % 10
		if r == 0 {
			return false
		}

		if r > l {
			return false
		}

		a[r] = 1
		n = n / 10
		if n == 0 {
			break
		}
	}

	for i := 1; i < len(a); i++ {
		if a[i] == 0 {
			return false
		}
	}

	return true

}

// PandigitalPrime project euler problem 41
func PandigitalPrime() int {
	max := 0
	for i := 9; i > 0; i-- {
		nums := []int{}
		for j := 1; j <= i; j++ {
			nums = append(nums, j)
		}

		r := GenPermutation(nums)

		for _, val := range r {
			if Prime(val) {
				if max < val {
					max = val
				}
			}
		}

		// try from big to small
		// if found one, quit loop
		if max != 0 {
			break
		}
	}
	return max
}

// CodedTriangleNumbers project euler problem 42
func CodedTriangleNumbers() (int, error) {
	count := 0

	r := genTriangleNums(1000)

	f, err := os.Open("./p042_words.txt")
	if err != nil {
		return count, err
	}
	defer f.Close()

	scanner := ReadNames(f)

	for scanner.Scan() {
		word := scanner.Text()

		c := 0
		for _, v := range word {
			if v == '"' {
				continue
			}
			c += int(v) - 'A' + 1
		}

		if r[c] {
			count++
		}

	}

	return count, nil

}

func genTriangleNums(n int) map[int]bool {
	r := make(map[int]bool)

	for i := 1; i <= n; i++ {
		r[i*(i+1)/2] = true
	}

	return r
}

// SubStringDivisibility project euler problem 43
func SubStringDivisibility() int {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	n := GenPermutation(nums)
	sum := 0
	for i := 0; i < len(n); i++ {
		if n[i] < 1000000000 {
			continue
		}
		if isProperty(n[i]) {
			// fmt.Println(n[i])
			sum += n[i]
		}
	}

	return sum
}

func isProperty(num int) bool {
	divider := []int{2, 3, 5, 7, 11, 13, 17}
	i := len(divider) - 1
	r := num
	for {
		if (r%1000)%divider[i] != 0 {
			return false
		}
		i--
		r = r / 10

		if i < 0 {
			break
		}
	}

	return true
}

// ChampernowneConst project euler problem 40
func ChampernowneConst(n int) int {
	r := 0
	s := 0
	k := 0
	for {
		if s >= n {
			break
		}
		r = s

		s += 9 * int(math.Pow10(k)) * (k + 1)
		k++
	}
	h := n - r - 1
	t := int(math.Pow10(k-1)) + h/k
	p := h % k

	ret, _ := strconv.Atoi(strconv.Itoa(t)[p : p+1])

	return ret
}

// PentagonNumbers project euler problem 44
func PentagonNumbers() int {
	min := math.MaxInt64

	for i := 1; i < 5000; i++ {
		p1 := calPentagonNumber(i)
		for j := i + 1; j < 5000; j++ {
			p2 := calPentagonNumber(j)
			if isPentagonNumber(p1 + p2) {
				d := p2 - p1
				if isPentagonNumber(d) {
					if d < min {
						min = d
					}
				}
			}
		}
	}
	return min

}

func calPentagonNumber(n int) (num int) {
	return (n * (3*n - 1)) / 2
}

func isPentagonNumber(num int) bool {
	r := math.Sqrt(float64(1 + 24*num))
	return isWholeNum(r) && (int(r)%6 == 5)
}

func isWholeNum(num float64) bool {
	return num == math.Trunc(num)
}
