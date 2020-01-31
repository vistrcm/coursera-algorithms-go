package sort

import "sort"

//InsertionBounded sort elements of a between lo and hi using Insertion sort.
func InsertionBounded(a sort.Interface, lo, hi int) {
	for i := lo; i < hi; i++ {
		for j := i; j > 0; j-- {
			if a.Less(j, j-1) {
				a.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}

//Insertion implements Insertion sort
func Insertion(a sort.Interface) {
	lo := 0
	hi := a.Len()

	InsertionBounded(a, lo, hi)
}
