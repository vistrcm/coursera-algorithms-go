package random

import (
	"math/rand"
)

//Shuffle n times
func Shuffle(n int64, swap func(p, n int64)) {
	var i int64
	for i = 0; i < n; i++ {
		r := rand.Int63n(i + 1)
		swap(i, r)
	}
}
