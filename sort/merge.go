package sort

import (
	"math"
	"sort"
)

const (
	cutoff = 7 // cutoff for insertion sort
)

//Merge merge sort for integers
func Merge(a sort.IntSlice) {
	aux := make(sort.IntSlice, a.Len())

	intSort(a, aux, 0, a.Len()-1, func(a, b int) bool { return a < b })
}

func intSort(a sort.IntSlice, aux sort.IntSlice, lo, hi int, less func(a, b int) bool) {

	if hi <= (lo + cutoff - 1) {
		InsertionBounded(a, lo, hi+1)
		return
	}
	mid := lo + (hi-lo)/2
	intSort(a, aux, lo, mid, less)
	intSort(a, aux, mid+1, hi, less)
	if !less(a[mid+1], a[mid]) { // already sorted no need to continue
		return
	}
	merge(a, aux, lo, mid, hi, less)
}

func merge(a sort.IntSlice, aux sort.IntSlice, lo, mid, hi int, less func(a, b int) bool) {

	// copy a to aux
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	var i = lo
	var j = mid + 1

	for k := lo; k <= hi; k++ {
		if i > mid {
			// nothing left on the left part. Take from right
			a[k] = aux[j]
			j++
		} else {
			if j > hi {
				// nothing left on the right part. Take from left.
				a[k] = aux[i]
				i++
			} else {
				if less(aux[j], aux[i]) {
					a[k] = aux[j]
					j++
				} else {
					a[k] = aux[i]
					i++
				}
			}
		}
	}
}

//MergeBU bottom up version of merge sort
func MergeBU(a sort.IntSlice) {
	N := a.Len()
	aux := make(sort.IntSlice, a.Len())
	for sz := 1; sz < N; sz = sz + sz {
		for lo := 0; lo < N-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz-1, int(math.Min(float64(lo+sz+sz-1), float64(N-1))),
				func(a, b int) bool { return a < b })
		}
	}
}
