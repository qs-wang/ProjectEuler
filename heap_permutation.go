package projecteuler

import (
	"fmt"
	"strconv"
	"strings"
)

func GenPermutation(nums []int) []int {
	var r []int

	generate(nums, len(nums), &r)

	return r
}

func generate(nums []int, k int, r *[]int) {
	if k == 1 {
		n, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums[0:])), ""), "[]"))
		*r = append((*r), n)
	} else {
		generate(nums, k-1, r)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				nums[i], nums[k-1] = nums[k-1], nums[i]
			} else {
				nums[0], nums[k-1] = nums[k-1], nums[0]
			}
			generate(nums, k-1, r)
		}
	}
}
