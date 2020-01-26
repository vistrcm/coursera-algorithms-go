package sort

import (
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func generateRandomIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Int()
	}
	return s
}

func testSort(t *testing.T, sortFunc func(sort.Interface)) {
	// prepare big slice
	bigSlice := generateRandomIntSlice(1000)
	bigSliceSorted := make([]int, len(bigSlice))
	copy(bigSliceSorted, bigSlice)
	sort.Ints(bigSliceSorted)

	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "integers",
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
			sortFunc(IntSlice(a))
			got := a
			if !reflect.DeepEqual(got, tt.want) {
				// found how to get function name here: https://stackoverflow.com/questions/10742749/get-name-of-function-using-reflection
				fName := runtime.FuncForPC(reflect.ValueOf(sortFunc).Pointer()).Name()
				t.Errorf("%v for input %v. got %v, want %v", fName, tt.args, got, tt.want)
			}
		})
	}
}

func TestSelection(t *testing.T) {
	testSort(t, Selection)
}

func TestInsertion(t *testing.T) {
	testSort(t, Insertion)
}

// some benchmarks
func benchmarkSearch(sortFunc func(sort.Interface), size int, b *testing.B) {
	// prepare big slice
	bigSlice := generateRandomIntSlice(size)

	b.ResetTimer()
	// run size * push-pop operations on stack b.N
	for i := 0; i < b.N; i++ {
		for o := 0; o <= size; o++ {
			sortFunc(IntSlice(bigSlice))
		}
	}
}

// some benchmarks
func benchmarkSearchSorted(sortFunc func(sort.Interface), size int, b *testing.B) {
	// prepare big slice
	bigSlice := make([]int, size)
	for i := range bigSlice {
		bigSlice[i] = i
	}

	b.ResetTimer()
	// run size * push-pop operations on stack b.N
	for i := 0; i < b.N; i++ {
		for o := 0; o <= size; o++ {
			sortFunc(IntSlice(bigSlice))
		}
	}
}

// some benchmarks
func benchmarkSearchDescending(sortFunc func(sort.Interface), size int, b *testing.B) {
	// prepare big slice
	bigSlice := make([]int, size)
	for i := range bigSlice {
		bigSlice[i] = size - i
	}

	b.ResetTimer()
	// run size * push-pop operations on stack b.N
	for i := 0; i < b.N; i++ {
		for o := 0; o <= size; o++ {
			sortFunc(IntSlice(bigSlice))
		}
	}
}

func benchmarkSelection(size int, b *testing.B)       { benchmarkSearch(Selection, size, b) }
func benchmarkSelectionSorted(size int, b *testing.B) { benchmarkSearchSorted(Selection, size, b) }
func benchmarkSelectionDescending(size int, b *testing.B) {
	benchmarkSearchDescending(Selection, size, b)
}

func BenchmarkSelection_1(b *testing.B)    { benchmarkSelection(1, b) }
func BenchmarkSelection_10(b *testing.B)   { benchmarkSelection(10, b) }
func BenchmarkSelection_100(b *testing.B)  { benchmarkSelection(100, b) }
func BenchmarkSelection_1000(b *testing.B) { benchmarkSelection(1000, b) }

func BenchmarkSelectionSorted_1(b *testing.B)    { benchmarkSelectionSorted(1, b) }
func BenchmarkSelectionSorted_10(b *testing.B)   { benchmarkSelectionSorted(10, b) }
func BenchmarkSelectionSorted_100(b *testing.B)  { benchmarkSelectionSorted(100, b) }
func BenchmarkSelectionSorted_1000(b *testing.B) { benchmarkSelectionSorted(1000, b) }

func BenchmarkSelectionDescending_1(b *testing.B)    { benchmarkSelectionDescending(1, b) }
func BenchmarkSelectionDescending_10(b *testing.B)   { benchmarkSelectionDescending(10, b) }
func BenchmarkSelectionDescending_100(b *testing.B)  { benchmarkSelectionDescending(100, b) }
func BenchmarkSelectionDescending_1000(b *testing.B) { benchmarkSelectionDescending(1000, b) }

func benchmarkInsertion(size int, b *testing.B)       { benchmarkSearch(Insertion, size, b) }
func benchmarkInsertionSorted(size int, b *testing.B) { benchmarkSearchSorted(Insertion, size, b) }
func benchmarkInsertionDescending(size int, b *testing.B) {
	benchmarkSearchDescending(Insertion, size, b)
}

func BenchmarkInsertion_1(b *testing.B)    { benchmarkInsertion(1, b) }
func BenchmarkInsertion_10(b *testing.B)   { benchmarkInsertion(10, b) }
func BenchmarkInsertion_100(b *testing.B)  { benchmarkInsertion(100, b) }
func BenchmarkInsertion_1000(b *testing.B) { benchmarkInsertion(1000, b) }
func BenchmarkInsertion_10000(b *testing.B) { benchmarkInsertion(10000, b) }

func BenchmarkInsertionSorted_1(b *testing.B)    { benchmarkInsertionSorted(1, b) }
func BenchmarkInsertionSorted_10(b *testing.B)   { benchmarkInsertionSorted(10, b) }
func BenchmarkInsertionSorted_100(b *testing.B)  { benchmarkInsertionSorted(100, b) }
func BenchmarkInsertionSorted_1000(b *testing.B) { benchmarkInsertionSorted(1000, b) }
func BenchmarkInsertionSorted_10000(b *testing.B) { benchmarkInsertionSorted(10000, b) }

func BenchmarkInsertionDescending_1(b *testing.B)    { benchmarkInsertionDescending(1, b) }
func BenchmarkInsertionDescending_10(b *testing.B)   { benchmarkInsertionDescending(10, b) }
func BenchmarkInsertionDescending_100(b *testing.B)  { benchmarkInsertionDescending(100, b) }
func BenchmarkInsertionDescending_1000(b *testing.B) { benchmarkInsertionDescending(1000, b) }
func BenchmarkInsertionDescending_10000(b *testing.B) { benchmarkInsertionDescending(10000, b) }
