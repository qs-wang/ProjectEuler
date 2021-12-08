package projecteuler

import (
	"math"
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
