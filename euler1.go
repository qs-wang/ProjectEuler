package projecteuler

func NonAbundantSums() int {
	abundantNums := make([]int, 28124)
	k := 0
	for i := 1; i <= 28123; i++ {
		if GetFactorSum(i) > i {
			abundantNums[k] = i
			k++
		}
	}

	set := make([]int, 28124)
	for i := 0; i < k; i++ {
		for j := i; j < k; j++ {
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
