package projecteuler

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
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
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
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

// SumSquareDifference1 https://projecteuler.net/problem=6
func SumSquareDifference1(limit int) int {
	sum := 0

	for i := 1; i <= limit; i++ {
		for j := i + 1; j <= limit; j++ {
			sum += i * j
		}
	}

	return 2 * sum
}

// SumSquareDifference https://projecteuler.net/problem=6
func SumSquareDifference(limit int) int {
	sum := 0
	square := 0
	for i := 1; i <= limit; i++ {
		sum += i
		square += i * i
	}

	return sum*sum - square

}

// sum_square = (n*(n+1)*(2*n+1))/6
// sum_number = (n*(n+1))/2

//NthPrime https://projecteuler.net/problem=7
func NthPrime(num int) int {
	if num == 1 {
		return 2
	}

	i := 2
	p := 3
	for {
		if i >= num {
			break
		}

		p += 2
		if Prime(p) {
			i++
		}
	}

	return p
}

//LastestProductSeries https://projecteuler.net/problem=8
func LastestProductSeries(num int) int {
	seriesNum := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"

	max := 0

	multi := doMulti(seriesNum, 0, num)

	if multi > max {
		max = multi
	}

	pos := 0

	for {
		if pos+num >= len(seriesNum) {
			break
		}

		p, _ := strconv.Atoi(string(seriesNum[pos]))
		n, _ := strconv.Atoi(string(seriesNum[pos+num]))

		if n > p {
			multi := doMulti(seriesNum, pos+1, num)
			if multi > max {
				max = multi
			}
		}

		pos++
	}

	return max
}

func doMulti(nums string, start int, len int) int {
	multi := 1
	for i := 0; i < len; i++ {
		t, _ := strconv.Atoi(string(nums[start+i]))
		multi *= t
	}

	return multi
}

// SpecialPythagoreanTriple https://projecteuler.net/problem=9
func SpecialPythagoreanTriple(num int) int {
	for i := 1; i < num; i++ {
		for j := i + 1; j < num; j++ {
			c := float64((1000000 - 2*i*j)) / float64(2*num)
			if c == math.Trunc(c) && (i+j+int(c) == num) {
				return i * j * int(c)
			}
			if c >= float64(num-i-j) {
				break
			}
		}
	}

	return -1

}

// SumOfPrime https://projecteuler.net/problem=10
func SumOfPrime(num int) int {
	sum := 2
	for i := 3; i <= num; i += 2 {
		if Prime(i) {
			sum += i
		}
	}
	return sum
}

// SumOfPrimeSieve https://projecteuler.net/problem=10
func SumOfPrimeSieve(num int) int {
	sum := 0
	sieve := make([]bool, num+1)
	for i := 2; i <= num; i++ {
		if sieve[i] {
			continue
		}
		sum += i

		for j := 2; j*i < num+1; j++ {
			sieve[j*i] = true
		}
	}

	return sum
}

// LargestProductInGrid https://projecteuler.net/problem=11
func LargestProductInGrid() int64 {
	matrix := [20][20]int{
		{8, 2, 22, 97, 38, 15, 00, 40, 00, 75, 04, 05, 07, 78, 52, 12, 50, 77, 91, 8},
		{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 00},
		{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65},
		{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91},
		{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
		{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
		{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
		{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
		{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
		{21, 36, 23, 9, 75, 00, 76, 44, 20, 45, 35, 14, 00, 61, 33, 97, 34, 31, 33, 95},
		{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92},
		{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 00, 17, 54, 24, 36, 29, 85, 57},
		{86, 56, 00, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
		{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40},
		{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
		{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
		{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
		{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16},
		{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54},
		{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48},
	}

	max := int64(0)

	for i := 0; i < 17; i++ {
		for j := 0; j < 17; j++ {
			max = Max(max, maxProduct(matrix, i, j))
		}
	}

	return max

}

func maxProduct(matrix [20][20]int, row int, col int) int64 {
	max := int64(0)

	for i := row; i < row+4; i++ {
		max = Max(max, int64(matrix[i][col])*int64(matrix[i][col+1])*int64(matrix[i][col+2])*int64(matrix[i][col+3]))
	}

	for j := col; j < col+4; j++ {
		max = Max(max, int64(matrix[row][j])*int64(matrix[row+1][j])*int64(matrix[row+1][j])*int64(matrix[row+3][j]))
	}

	max = Max(max, int64(matrix[row][col])*int64(matrix[row+1][col+1])*int64(matrix[row+2][col+2])*int64(matrix[row+3][col+3]))

	max = Max(max, int64(matrix[row+3][col])*int64(matrix[row+2][col+1])*int64(matrix[row+1][col+2])*int64(matrix[row][col+3]))

	return max
}

// Max calculate the max value between 2 int64 values
func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// HighlyDividedTriangularNumber https://projecteuler.net/problem=12
func HighlyDividedTriangularNumber() int {

	i := 1
	amt := 1
	for {
		if FactorsAmt(amt, 500) > 500 {
			break
		}

		i++
		amt = amt + i
	}

	return amt
}

func FactorsAmt(num int, min int) int {
	if int(math.Sqrt(float64(num))) < min {
		return 0
	}

	count := 0
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		remain := num % i
		if remain == 0 {
			count += 2
		}
	}

	return count
}

func LargeSum() string {
	file, err := os.Open("./largesum.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := big.NewInt(0)
	for scanner.Scan() {
		line := scanner.Text()
		n := new(big.Int)
		n.SetString(line, 10)
		sum.Add(n, sum)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	stringSum := sum.Text(10)

	return stringSum[0:10]
}

// LongestCollatzSequence https://projecteuler.net/problem=14
func LongestCollatzSequence(num int) int {
	max := 0
	n := 0

	for i := num; i > 0; i-- {
		count := 0
		for j := i; j > 1; {
			if j%2 == 0 {
				j = j / 2
				count++
			} else {
				j = (3*j + 1) / 2
				count += 2
			}
		}
		if count > max {
			n = i
			max = count
		}
	}

	return n
}

// LatticePaths https://projecteuler.net/problem=15
func LatticePaths() int {
	matrix := [21][21]int{}
	for i := 0; i < 21; i++ {
		matrix[0][i] = 1
		matrix[i][0] = 1
	}

	for i := 1; i < 21; i++ {
		for j := 1; j < 21; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}

	return matrix[20][20]

}

//PowerDigitSum https://projecteuler.net/problem=16
func PowerDigitSum() int64 {
	// 2^10 = 1024 => 4 digits
	// 1024*1024 => 7 digits
	// 2^1000 => 3xx. 400 should be safe
	num := [400]int64{}

	num[0] = 1

	carry := int64(0)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 400; j++ {
			c := num[j]*int64(2) + carry
			if c >= 10 {
				num[j] = c - 10
				carry = 1
			} else {
				num[j] = c
				carry = 0
			}
		}
	}

	sum := int64(0)
	for i := 0; i < 400; i++ {
		sum += num[i]
	}

	return sum
}

// NumberLetterCounts https://projecteuler.net/problem=17
func NumberLetterCounts() int {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += LetterCount(i)
	}

	return sum
}

// LetterCount returns the number of letters in a number
func LetterCount(num int) int {
	countMap := map[int]int{
		1:  3,
		2:  3,
		3:  5,
		4:  4,
		5:  4,
		6:  3,
		7:  5,
		8:  5,
		9:  4,
		10: 3,
		11: 6,
		12: 6,
		13: 8,
		14: 8,
		15: 7,
		16: 7,
		17: 9,
		18: 8,
		19: 8,
		20: 6,
		30: 6,
		40: 5,
		50: 5,
		60: 5,
		70: 7,
		80: 6,
		90: 6,
		0:  0,
	}

	q := num
	count := 0

	reminder := q % 100
	if val, ok := countMap[reminder]; ok {
		count += val
	} else {
		r := reminder % 10
		p := reminder - r

		count = count + countMap[r] + countMap[p]
	}

	q = q / 100

	if q == 0 {
		return count
	}

	if reminder != 0 {
		count += 3
	}

	reminder = q % 10
	count = count + countMap[reminder]
	if reminder != 0 {
		count += 7
	}

	q = q / 10

	if q == 0 {
		return count
	}

	count = count + countMap[q] + 8

	return count
}

//MaximumPathSum1 https://projecteuler.net/problem=18
func MaximumPathSum1() int {
	arr := [][]int{
		{75, 0},
		{95, 64, 0},
		{17, 47, 82, 0},
		{18, 35, 87, 10, 0},
		{20, 4, 82, 47, 65, 0},
		{19, 1, 23, 75, 3, 34, 0},
		{88, 2, 77, 73, 7, 63, 67, 0},
		{99, 65, 4, 28, 6, 16, 70, 92, 0},
		{41, 41, 26, 56, 83, 40, 80, 70, 33, 0},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29, 0},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14, 0},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57, 0},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48, 0},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31, 0},
		{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23, 0},
	}
	max := 0

	matrix := [15][15]int{}
	matrix[0][0] = 75
	max = 75

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr[i])-1; j++ {
			if j == 0 {
				matrix[i][j] = (arr[i][j] + matrix[i-1][j])
			} else {
				matrix[i][j] = arr[i][j] + int(math.Max(float64(matrix[i-1][j]), float64(matrix[i-1][j-1])))
			}
			max = int(math.Max(float64(max), float64(matrix[i][j])))
		}
	}

	return max

}

// CountingSundays https://projecteuler.net/problem=19
func CountingSundays() int {
	startWeekday := 1
	count := 0

	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			startWeekday = getNextFirstWeekDay(year, month, startWeekday)
			if startWeekday == 0 {
				count++
			}
		}
	}
	return count
}

func getNextFirstWeekDay(year int, month int, day int) int {

	switch month {
	case 4, 6, 9, 11:
		return (day + 30/7) % 7
	case 1, 3, 5, 7, 8, 10, 12:
		return (day + 31/7) % 7
	case 2:
		if isLeapYear(year) {
			return (day + 29/7) % 7
		}

		return (day + 28/7) % 7

	}

	return 0
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// FactorialDigitNum https://projecteuler.net/problem=20
func FactorialDigitNum(num int) string {
	result := big.NewInt(1)
	for i := 1; i <= num; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}

	fmt.Println(result.Text(10))
	count := big.NewInt(0)

	for {

		m := big.NewInt(0)
		result, mod := result.DivMod(result, big.NewInt(10), m)

		fmt.Println(mod.Text(10))

		count = count.Add(count, mod)

		if result.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}

	return count.Text(10)

}
