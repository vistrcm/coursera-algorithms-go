package sort

import (
	"math/rand"
	"sort"
)

const (
	quickCutoff = 10 // cutoff for insertion sort
)

//Quick implements quicksort algorithm
func Quick(a sort.IntSlice) {
	rand.Shuffle(a.Len(), a.Swap)
	qIntSort(a, 0, a.Len()-1)
}

func qIntSort(a sort.IntSlice, lo, hi int) {
	if hi <= (lo + quickCutoff - 1) {
		InsertionBounded(a, lo, hi+1)
		return
	}

	j := partition(a, lo, hi)
	qIntSort(a, lo, j-1)
	qIntSort(a, j+1, hi)
}

func partition(a sort.IntSlice, lo, hi int) int {
	i := lo
	j := hi + 1
	for {
		for i++; a.Less(i, lo); i++ { // find item on the left to swap
			if i == hi {
				break
			}
		}

		for j--; a.Less(lo, j); j-- { // find item on the right to swap
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}

		a.Swap(i, j)
	}
	a.Swap(lo, j)
	return j
}
