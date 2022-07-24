package st_test

import (
	"fmt"
	"github.com/vistrcm/coursera-algorithms-go/st"
	"github.com/vistrcm/coursera-algorithms-go/st/rbbst"
)

type comparableString string

func (s comparableString) CompareTo(o interface{}) int {
	other := o.(comparableString)
	if s < other {
		return -1
	} else if s > other {
		return +1
	}
	return 0
}

func ExampleBST() {
	keys := []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	stable := st.BST{}

	for i, key := range keys {
		stable.Put(comparableString(key), i)
	}

	for _, k := range stable.Keys() {
		fmt.Printf("%s -> %d\n", k, stable.Get(k))
	}
	// Output: A -> 8
	// C -> 4
	// E -> 12
	// H -> 5
	// L -> 11
	// M -> 9
	// P -> 10
	// R -> 3
	// S -> 0
	// X -> 7
}

func ExampleRBBST() {
	keys := []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	stable := rbbst.RBBST{}

	for i, key := range keys {
		stable.Put(comparableString(key), i)
	}

	for _, k := range stable.Keys() {
		fmt.Printf("%s -> %d\n", k, stable.Get(k))
	}
	// Output: A -> 8
	// C -> 4
	// E -> 12
	// H -> 5
	// L -> 11
	// M -> 9
	// P -> 10
	// R -> 3
	// S -> 0
	// X -> 7
}
