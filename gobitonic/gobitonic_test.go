package gobitonic

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	arrsize := 1 << 5
	testarr := make([]int, arrsize)
	testarr2 := make([]int, arrsize)
	for i := 0; i < arrsize; i++ {
		testarr[i] = rand.Intn(100 - 0 + 1)
		testarr2[i] = testarr[i]
	}

	Sort(testarr, true)
	sort.Ints(testarr2)

	for i := 0; i < arrsize; i++ {
		if testarr[i] != testarr2[i] {
			t.Errorf("%d %d", testarr[i], testarr2[i])
		}
	}
}
