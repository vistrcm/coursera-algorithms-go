package sort

import (
	"math/rand"
	"sort"
)

//Quick3Way implement 3-way quick sort
func Quick3Way(a sort.IntSlice) {
	rand.Shuffle(a.Len(), a.Swap)
	q3way(a, 0, a.Len()-1)
}

func q3way(a sort.IntSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	lt := lo
	gt := hi
	v := a[lo] // partition key
	i := lo + 1
	for i <= gt {
		cmp := a[i] - v
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
