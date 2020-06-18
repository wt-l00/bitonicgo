package gobitonic

func BitonicSort(x []int, up bool) bool {
	if !isPowerOfTwo(len(x)) {
		return false
	}

	if len(x) > 1 {
		mid := len(x) / 2
		BitonicSort(x[:mid], true)
		BitonicSort(x[mid:], false)
		SubSort(x, up)
	}
	return true
}

func isPowerOfTwo(num int) bool {
	if num&(num-1) == 0 {
		return true
	}
	return false
}

func SubSort(x []int, up bool) {
	if len(x) > 1 {
		CompareAndSwap(x, up)
		mid := len(x) / 2
		SubSort(x[:mid], up)
		SubSort(x[mid:], up)
	}
}

func CompareAndSwap(x []int, up bool) {
	mid := len(x) / 2
	for i := 0; i < mid; i++ {
		if x[i] > x[mid+i] == up {
			x[i], x[mid+i] = x[mid+i], x[i]
		}
	}
}
