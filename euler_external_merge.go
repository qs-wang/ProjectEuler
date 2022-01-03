package projecteuler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"container/heap"
)

// CompString compare 2 string, returns 1 if s1>s2, return 0 otherwise
func CompString(s1 string, s2 string) int {
	r1 := []rune(strings.ToLower(s1))
	r2 := []rune(strings.ToLower(s2))

	for i := 0; ; i++ {
		if i >= len(r1) {
			return 0
		}

		if i >= len(r2) {
			return 1
		}

		if r1[i] > r2[i] {
			return 1
		}

		if r1[i] < r2[i] {
			return 0
		}
	}
}

// SortStrings sort strings in ascending order
func SortStrings(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if CompString(s[i], s[j]) == 1 {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

// ReadNames read names split
func ReadNames(f io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(f)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if atEOF && len(data) != 0 {
			return len(data), data, nil
		}

		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		return 0, nil, nil
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)

	return scanner
}

func SplitToSortedFile(file string) (err error) {
	f, err := os.Open("./p022_names.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := ReadNames(f)

	var words = make([]string, 200)
	i := 0
	j := 0
	for scanner.Scan() {
		words[i] = scanner.Text()
		if i == 199 {
			sorted := SortStrings(words[:])
			writeStringsToFile(sorted, j)
			i = 0
			j++
			words = make([]string, 200)
		} else {
			i++
		}
	}

	if i != 199 {
		sorted := SortStrings(words[:i])
		writeStringsToFile(sorted, j)
	}

	return
}

type Item struct {
	value   string // The value of the item; arbitrary.
	index   int
	sacnner *bufio.Scanner
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return CompString(pq[i].value, pq[j].value) == 0
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, value string, priority int) {
// 	item.value = value
// 	heap.Fix(pq, item.index)
// }

func MergeSort() (err error) {
	pq := make(PriorityQueue, 26)

	for i := 0; i < 26; i++ {
		f, err := os.Open(fmt.Sprintf("./temp/s%d", i))
		if err != nil {
			return err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Scan()
		pq[i] = &Item{
			value:   scanner.Text(),
			index:   i,
			sacnner: scanner,
		}
	}
	heap.Init(&pq)

	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	for {
		if len(pq) == 1 {
			break
		}

		item := heap.Pop(&pq).(*Item)
		fmt.Fprintln(w, item.value)

		if item.sacnner.Scan() {
			newItem := &Item{
				value:   item.sacnner.Text(),
				sacnner: item.sacnner,
			}
			pq.Push(newItem)
		}

	}

	// Write all remain items
	fmt.Fprintln(w, pq[0].value)
	for pq[0].sacnner.Scan() {
		fmt.Fprintln(w, pq[0].sacnner.Text())
	}

	return w.Flush()
}

func writeStringsToFile(words []string, seq int) (err error) {
	file, err := os.Create(fmt.Sprintf("temp/s%d", seq))
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range words {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func CalculateWordWorth(word string) int {
	runes := []rune(word)

	result := 0

	for i := 0; i < len(runes); i++ {
		result += (int(runes[i]) - 'A' + 1)
	}

	return result
}

// CalculateScore https://projecteuler.net/problem=22
// external merge solution
func CalculateScore(file string) (int, error) {

	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	amt := 0

	i := 1
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			worth := CalculateWordWorth(text[1 : len(text)-1])
			amt += i * worth
			i++
		}

	}

	return amt, nil
}
