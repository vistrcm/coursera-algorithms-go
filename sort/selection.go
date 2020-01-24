package sort

import "sort"

//Selection implements selection sort
func Selection(a sort.Interface) {
	n := a.Len()
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a.Less(j, min) {
				min = j
			}
		}
		a.Swap(i, min)
	}
}
