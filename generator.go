package main

import (
	"math/rand"
	"time"
)

func generate() ([]int, []int) {
	rand.Seed(time.Now().UnixMicro())
	parents := make([]int, rand.Intn(20)+2)
	inflows := make([]int, len(parents))
	parents[0] = -1
	inflows[0] = rand.Intn(9) + 1
	parents[1] = 0
	inflows[1] = rand.Intn(9) + 1
	for i := 2; i < len(parents); i++ {

		parents[i] = rand.Intn(i - 1)
		inflows[i] = rand.Intn(9) + 1
	}
	return parents, inflows
}
