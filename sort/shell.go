package sort

import "sort"

//Shell implements shell sort
func Shell(a sort.Interface) {
	n := a.Len()

	// 3x+1 increment sequence
	h := 1
	for h < n/3 {
		h = 3*h + 1
	} // 1, 4, 13, ...

	for h >= 1 {
		//h-sort the array
		for i := h; i < n; i++ {
			for j := i; j >= h && a.Less(j, j-h); j -= h {
				a.Swap(j, j-h)
			}
		}
		h = h / 3
	}

}
