package main

import (
	"fmt"
	"sort"
)

// sort.Sort() uses:
// 	•	Quicksort for large arrays
// 	•	Insertion sort for small slices
// 	•	Heapsort fallback if quicksort is unbalanced (to avoid worst-case O(n²))

type Node struct {
	High, Low int
}

func main() {
	fmt.Println("Hello, World!")
	nums := []int{1, 2, 4, 3, 5}
	fmt.Println("Original slice:", nums)

	sort.Ints(nums)
	fmt.Println("Sorted slice:", nums)

	// sort.Strings([]string)
	// sort.Float64s([]float64)

	str := "ram"
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	sortedStr := string(bytes)
	fmt.Println("Sorted string:", sortedStr)

	strList := []string{"ram", "shyam", "mohan"}
	sort.Strings(strList)
	fmt.Println("Sorted string slice:", strList)

	nodes := make([]Node, 0)
	for _, num := range nums {
		nodes = append(nodes, Node{High: num * 10, Low: num})
	}
	fmt.Println("Nodes slice:", nodes)

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].High < nodes[j].High
	})
	fmt.Println("Sorted nodes slice:", nodes)
}