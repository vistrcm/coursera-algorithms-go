package sort

import "sort"

func Heap(a sort.IntSlice) {
	N := a.Len()
	for k := N / 2; k >= 1; k-- {
		sink(a, k, N)
	}
	for N > 1 {
		exch(a, 1, N)
		N--
		sink(a, 1, N)
	}
}

func sink(a sort.IntSlice, k, N int) {
	for 2*k <= N {
		j := 2 * k
		if j < N && less(a, j, j+1) {
			j++
		}
		if !less(a, k, j) {
			break
		}
		exch(a, k, j)
		k = j
	}
}

//less version of 1-indexed arrays
func less(a sort.IntSlice, i, j int) bool {
	return a[i-1] < a[j-1]
}

//exch version of 1-indexed arrays
func exch(a sort.IntSlice, i, j int) {
	a[i-1], a[j-1] = a[j-1], a[i-1]
}
