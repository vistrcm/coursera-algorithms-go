package sort

import "sort"

//Insertion implements Insertion sort
func Insertion(a sort.Interface) {
	n := a.Len()

	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if a.Less(j, j-1) {
				a.Swap(j, j-1)
			} else {
				break
			}
		}
	}

}
