package sort

import (
	"sort"
)

//Quick3Way implement 3-way quick sort
func Quick3Way(a sort.IntSlice){
	q3way(a, 0, a.Len()-1)
}

func q3way(a sort.IntSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	lt := lo
	gt := hi
	v := a[lo]
	i := lo
	for i <= gt {
		cmp := compare(a, i, v)
		if cmp < 0 {
			a.Swap(lt, i)
			lt++
			i++
		} else if cmp > 0 {
			a.Swap(i, gt)
			gt--
		} else {
			i++
		}
	}
	q3way(a, lo, lt-1)
	q3way(a, gt+1, hi)

}

func compare(a sort.IntSlice, i int, v int) int {
	return a[i] - a[v]
}
