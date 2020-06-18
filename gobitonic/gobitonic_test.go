package gobitonic

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		arrsize    int
		max        int
		min        int
		shouldsort bool
	}{
		{8, 10, 1, true},
		{1024, 1000, 1, true},
		{3, 100, 1, false},
	}

	for _, tt := range tests {
		testarr := make([]int, tt.arrsize)
		testarr2 := make([]int, tt.arrsize)
		for i := 0; i < tt.arrsize; i++ {
			testarr[i] = rand.Intn(tt.max - tt.min + 1)
			testarr2[i] = testarr[i]
		}

		if tt.shouldsort {
			BitonicSort(testarr, true)
			sort.Ints(testarr2)
			for i := 0; i < tt.arrsize; i++ {
				if testarr[i] != testarr2[i] {
					t.Errorf("%d %d", testarr[i], testarr2[i])
				}
			}
		} else {
			if BitonicSort(testarr, true) != false {
				t.Errorf("shuold not sort")
			}
		}
	}
}

func TestCompareAndSwap(t *testing.T) {
	tests := []struct {
		arrsize int
		max     int
		min     int
	}{
		{16, 10, 1},
		{32, 1000, 1},
	}

	for _, tt := range tests {
		testarr := make([]int, tt.arrsize)
		testarr2 := make([]int, tt.arrsize)
		for i := 0; i < tt.arrsize; i++ {
			testarr[i] = rand.Intn(tt.max - tt.min + 1)
			testarr2[i] = testarr[i]
		}
		BitonicSort(testarr, true)
		sort.Ints(testarr2)

	}
}
