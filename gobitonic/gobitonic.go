package gobitonic

func Sort(x []int, up bool) {
	if len(x) > 1 {
		mid := len(x) / 2
		Sort(x[:mid], true)
		Sort(x[mid:], false)
		SubSort(x, up)
	}
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
