package main

import (
	"os"
)

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func getTotalInflow(inflows []int, subInflows []int, children map[int][]int, i int) int {
	if subInflows[i] == 0 {
		subInflows[i] = inflows[i]
		for _, j := range children[i] {
			subInflows[i] += getTotalInflow(inflows, subInflows, children, j)
		}
	}
	return subInflows[i]
}

func getSubInflows(parents []int, inflows []int) []int {
	subInflows := make([]int, len(parents))
	children := make(map[int][]int, len(parents))
	for i := 0; i < len(parents); i++ {
		if parents[i] < 0 {
			continue
		}
		children[parents[i]] = append(children[parents[i]], i)
	}
	getTotalInflow(inflows, subInflows, children, 0)
	return subInflows
}

func main() {
	// parents := []int{-1, 0, 0, 1, 1, 2}
	// inflows := []int{1, 2, 2, 1, 1, 1}
	parents, inflows := generate()
	subInflows := getSubInflows(parents, inflows)

	totalInflow := subInflows[0]
	avgInflow := totalInflow / 2
	cutIx := -1
	cutScore := 1<<63 - 1
	for i := 0; i < len(subInflows); i++ {
		score := abs(subInflows[i] - avgInflow)
		if score < cutScore {
			cutIx = i
			cutScore = score
		}
	}

	printSlice(os.Stderr, "parents", parents)
	printSlice(os.Stderr, "inputs", inflows)
	printDot(os.Stdout, parents, inflows, subInflows, cutIx)
}
