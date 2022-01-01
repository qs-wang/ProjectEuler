package projecteuler

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompString(t *testing.T) {
	assert.Equal(t, 1, CompString("abc", "ab"))
	assert.Equal(t, 0, CompString("abc", "abc"))
	assert.Equal(t, 0, CompString("ab", "abc"))
	assert.Equal(t, 1, CompString("d", "abc"))
	assert.Equal(t, 0, CompString("d", "D"))
	assert.Equal(t, 1, CompString("D", "a"))
	assert.Equal(t, 1, CompString("d", "Beeee"))
	assert.Equal(t, 1, CompString("ZONIA", "ZONA"))
}

func TestSortStrings(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"c", "b", "a"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"c", "a", "b"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"b", "a", "c"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"b", "c", "a"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"c", "a", "b"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"c", "b", "a"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"a", "c", "b"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"a", "b", "c"}))
}

func TestSortStrings1(t *testing.T) {
	assert.Equal(t, []string{"a", "ab", "ac"}, SortStrings([]string{"ac", "ab", "a"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"c", "a", "b"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"b", "a", "c"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"b", "c", "a"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"c", "a", "b"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"c", "b", "a"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"a", "c", "b"}))
	assert.Equal(t, []string{"a", "b", "c"}, SortStrings([]string{"a", "b", "c"}))
}

func TestSortStrings2(t *testing.T) {
	assert.Equal(t, []string{"ZONA", "ZONIA"}, SortStrings([]string{"ZONIA", "ZONA"}))
}

func TestReadNames(t *testing.T) {
	assert.Equal(t, []string{"c", "b", "a"}, ReadNames(strings.NewReader("c,b,a")))
	assert.Equal(t, []string{"abc"}, ReadNames(strings.NewReader("abc")))
	assert.Equal(t, []string([]string(nil)), ReadNames(strings.NewReader("")))
}

func TestReadNamesFromFile(t *testing.T) {
	fmt.Println("hello")
	f, err := os.Open("./p022_names.txt")
	defer f.Close()

	assert.Nil(t, err)
	scanner := ReadNames(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func TestSplitToSortedFile(t *testing.T) {
	err := SplitToSortedFile("p022_names.txt")

	if err != nil {
		assert.Fail(t, "Failed")
	}
}

func TestMergeSort(t *testing.T) {
	err := MergeSort()

	if err != nil {
		assert.Fail(t, "Failed")
	}
}

func TestWordWorth(t *testing.T) {
	assert.Equal(t, 53, CalculateWordWorth("COLIN"))
}

func TestCalculateScore(t *testing.T) {
	score, err := CalculateScore("./output.txt")

	if err != nil {
		assert.Fail(t, "failed")
	}

	assert.Equal(t, 871198282, score)
}
