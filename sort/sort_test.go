package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"testing"
	"time"
)

func generateRandomIntSlice(n int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Int()
	}
	return s
}

func testSort(t *testing.T, sortFunc func(slice sort.IntSlice)) {
	// prepare big slice
	bigSlice := generateRandomIntSlice(10000)
	bigSliceSorted := make([]int, len(bigSlice))
	copy(bigSliceSorted, bigSlice)
	sort.Ints(bigSliceSorted)

	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "integers_even",
			args: []int{5, 4, 3, 2, 1, 0},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "integers_odd",
			args: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "more than ten",
			args: []int{11, 10, 8, 9, 7, 5, 6, 3, 4, 1, 2},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		{
			name: "BigSlice",
			args: bigSlice,
			want: bigSliceSorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// make a copy
			a := make([]int, len(tt.args))
			copy(a, tt.args)

			// sort
			sortFunc(a)
			got := a
			if !reflect.DeepEqual(got, tt.want) {
				// found how to get function name here: https://stackoverflow.com/questions/10742749/get-name-of-function-using-reflection
				fName := runtime.FuncForPC(reflect.ValueOf(sortFunc).Pointer()).Name()
				t.Errorf("%v for input %v. got %v, want %v", fName, tt.args, got, tt.want)
			}
		})
	}
}

//onIntSlice change signature of a function to use sort.Interface as sort.IntSlice
func onIntSlice(src func(a sort.Interface)) func(slice sort.IntSlice) {
	return func(slice sort.IntSlice) {
		src(slice)
	}
}

func TestSelection(t *testing.T) {
	testSort(t, onIntSlice(Selection))
}

func TestInsertion(t *testing.T) {
	testSort(t, onIntSlice(Insertion))
}

func TestShell(t *testing.T) {
	testSort(t, onIntSlice(Shell))
}

func TestMerge(t *testing.T) {
	testSort(t, Merge)
}

func TestMergeBU(t *testing.T) {
	testSort(t, MergeBU)
}

func TestQuick(t *testing.T) {
	testSort(t, Quick)
}

func TestQuick3Way(t *testing.T) {
	testSort(t, Quick3Way)
}

// benchmark helpers
func bRandom(sortFunc func(slice sort.IntSlice), size int, b *testing.B) {
	// prepare big slice
	bigSlice := generateRandomIntSlice(size)

	b.ResetTimer()
	// run size * push-pop operations on stack b.N
	for i := 0; i < b.N; i++ {
		for o := 0; o <= size; o++ {
			sortFunc(bigSlice)
		}
	}
}
func bASC(sortFunc func(slice sort.IntSlice), size int, b *testing.B) {
	// prepare big slice
	bigSlice := make([]int, size)
	for i := range bigSlice {
		bigSlice[i] = i
	}

	b.ResetTimer()
	// run size * push-pop operations on stack b.N
	for i := 0; i < b.N; i++ {
		for o := 0; o <= size; o++ {
			sortFunc(bigSlice)
		}
	}
}
func bDESC(sortFunc func(slice sort.IntSlice), size int, b *testing.B) {
	// prepare big slice
	bigSlice := make([]int, size)
	for i := range bigSlice {
		bigSlice[i] = size - i
	}

	b.ResetTimer()
	// run size * push-pop operations on stack b.N
	for i := 0; i < b.N; i++ {
		for o := 0; o <= size; o++ {
			sortFunc(bigSlice)
		}
	}
}
func stdSortWrapper(slice sort.IntSlice) {
	sort.Sort(slice)
}

type testFunc func(sortFunc func(slice sort.IntSlice), size int, b *testing.B)
type sortFunc func(slice sort.IntSlice)

func benchGenerator(testF testFunc, sortF sortFunc, size int) func(b *testing.B) {
	return func(b *testing.B) {
		testF(sortF, size, b)
	}
}

func BenchmarkSort(b *testing.B) {
	tests := []struct {
		name string
		f    func(a sort.IntSlice)
	}{
		//{name: "Selection", f: onIntSlice(Selection)},
		{name: "STD", f: stdSortWrapper},
		{name: "Insertion", f: onIntSlice(Insertion)},
		{name: "Shell", f: onIntSlice(Shell)},
		{name: "Merge", f: Merge},
		{name: "MergeBU", f: MergeBU},
		{name: "Quick", f: Quick},
		{name: "Quick3Way", f: Quick3Way},
	}

	testTypes := []struct {
		name string
		f    func(sortFunc func(slice sort.IntSlice), size int, b *testing.B)
	}{
		{name: "random", f: bRandom},
		{name: "ASC", f: bASC},
		{name: "DESC", f: bDESC},
	}

	dataSizes := []int{10, 100, 1000, 10000}
	for _, size := range dataSizes {
		for _, testType := range testTypes {
			for _, test := range tests {
				benchName := fmt.Sprintf("%s-%s-%d", testType.name, test.name, size)
				b.Run(benchName, benchGenerator(testType.f, test.f, size))
			}
		}
	}
}
