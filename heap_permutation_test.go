package projecteuler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenPurmutation(t *testing.T) {
	r := GenPermutation([]int{1, 2, 3})
	assert.Equal(t, []int{123, 213, 312, 132, 231, 321}, r)
}

// func TestGenPurmutation1(t *testing.T) {
// 	r := GenPermutation([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

// 	// fmt.Println(len(r))
// 	assert.Equal(t, []int{123, 213, 312, 132, 231, 321}, r)
// }
