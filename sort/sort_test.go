package sort

import (
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

func generateRandomIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Int()
	}
	return s
}

func testSort(t *testing.T, sortFunc func(slice sort.IntSlice)) {
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

// some benchmarks
func benchmarkRandom(sortFunc func(slice sort.IntSlice), size int, b *testing.B) {
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

// some benchmarks
func benchmarkSortedASC(sortFunc func(slice sort.IntSlice), size int, b *testing.B) {
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

// some benchmarks
func benchmarkSortedDESC(sortFunc func(slice sort.IntSlice), size int, b *testing.B) {
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

// set base results to compare
func stdSortWrapper(slice sort.IntSlice) {
	sort.Sort(slice)
}

func BenchmarkSTD_1(b *testing.B)     { benchmarkRandom(stdSortWrapper, 1, b) }
func BenchmarkSTD_10(b *testing.B)    { benchmarkRandom(stdSortWrapper, 10, b) }
func BenchmarkSTD_100(b *testing.B)   { benchmarkRandom(stdSortWrapper, 100, b) }
func BenchmarkSTD_1000(b *testing.B)  { benchmarkRandom(stdSortWrapper, 1000, b) }
func BenchmarkSTD_10000(b *testing.B) { benchmarkRandom(stdSortWrapper, 10000, b) }

func BenchmarkSTDSorted_1(b *testing.B)     { benchmarkSortedASC(stdSortWrapper, 1, b) }
func BenchmarkSTDSorted_10(b *testing.B)    { benchmarkSortedASC(stdSortWrapper, 10, b) }
func BenchmarkSTDSorted_100(b *testing.B)   { benchmarkSortedASC(stdSortWrapper, 100, b) }
func BenchmarkSTDSorted_1000(b *testing.B)  { benchmarkSortedASC(stdSortWrapper, 1000, b) }
func BenchmarkSTDSorted_10000(b *testing.B) { benchmarkSortedASC(stdSortWrapper, 10000, b) }

func BenchmarkSTDDescending_1(b *testing.B)    { benchmarkSortedDESC(stdSortWrapper, 1, b) }
func BenchmarkSTDDescending_10(b *testing.B)   { benchmarkSortedDESC(stdSortWrapper, 10, b) }
func BenchmarkSTDDescending_100(b *testing.B)  { benchmarkSortedDESC(stdSortWrapper, 100, b) }
func BenchmarkSTDDescending_1000(b *testing.B) { benchmarkSortedDESC(stdSortWrapper, 1000, b) }

// benchmark Selection

func benchmarkSelection(size int, b *testing.B) { benchmarkRandom(onIntSlice(Selection), size, b) }
func benchmarkSelectionSorted(size int, b *testing.B) {
	benchmarkSortedASC(onIntSlice(Selection), size, b)
}
func benchmarkSelectionDescending(size int, b *testing.B) {
	benchmarkSortedDESC(onIntSlice(Selection), size, b)
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

// benchmark Insertion

func benchmarkInsertion(size int, b *testing.B) { benchmarkRandom(onIntSlice(Insertion), size, b) }
func benchmarkInsertionSorted(size int, b *testing.B) {
	benchmarkSortedASC(onIntSlice(Insertion), size, b)
}
func benchmarkInsertionDescending(size int, b *testing.B) {
	benchmarkSortedDESC(onIntSlice(Insertion), size, b)
}

func BenchmarkInsertion_1(b *testing.B)     { benchmarkInsertion(1, b) }
func BenchmarkInsertion_10(b *testing.B)    { benchmarkInsertion(10, b) }
func BenchmarkInsertion_100(b *testing.B)   { benchmarkInsertion(100, b) }
func BenchmarkInsertion_1000(b *testing.B)  { benchmarkInsertion(1000, b) }
func BenchmarkInsertion_10000(b *testing.B) { benchmarkInsertion(10000, b) }

func BenchmarkInsertionSorted_1(b *testing.B)     { benchmarkInsertionSorted(1, b) }
func BenchmarkInsertionSorted_10(b *testing.B)    { benchmarkInsertionSorted(10, b) }
func BenchmarkInsertionSorted_100(b *testing.B)   { benchmarkInsertionSorted(100, b) }
func BenchmarkInsertionSorted_1000(b *testing.B)  { benchmarkInsertionSorted(1000, b) }
func BenchmarkInsertionSorted_10000(b *testing.B) { benchmarkInsertionSorted(10000, b) }

func BenchmarkInsertionDescending_1(b *testing.B)     { benchmarkInsertionDescending(1, b) }
func BenchmarkInsertionDescending_10(b *testing.B)    { benchmarkInsertionDescending(10, b) }
func BenchmarkInsertionDescending_100(b *testing.B)   { benchmarkInsertionDescending(100, b) }
func BenchmarkInsertionDescending_1000(b *testing.B)  { benchmarkInsertionDescending(1000, b) }
func BenchmarkInsertionDescending_10000(b *testing.B) { benchmarkInsertionDescending(10000, b) }

// benchmark Shell
func benchmarkShell(size int, b *testing.B)           { benchmarkRandom(onIntSlice(Shell), size, b) }
func benchmarkShellSorted(size int, b *testing.B)     { benchmarkSortedASC(onIntSlice(Shell), size, b) }
func benchmarkShellDescending(size int, b *testing.B) { benchmarkSortedDESC(onIntSlice(Shell), size, b) }

func BenchmarkShell_1(b *testing.B)     { benchmarkShell(1, b) }
func BenchmarkShell_10(b *testing.B)    { benchmarkShell(10, b) }
func BenchmarkShell_100(b *testing.B)   { benchmarkShell(100, b) }
func BenchmarkShell_1000(b *testing.B)  { benchmarkShell(1000, b) }
func BenchmarkShell_10000(b *testing.B) { benchmarkShell(10000, b) }

func BenchmarkShellSorted_1(b *testing.B)     { benchmarkShellSorted(1, b) }
func BenchmarkShellSorted_10(b *testing.B)    { benchmarkShellSorted(10, b) }
func BenchmarkShellSorted_100(b *testing.B)   { benchmarkShellSorted(100, b) }
func BenchmarkShellSorted_1000(b *testing.B)  { benchmarkShellSorted(1000, b) }
func BenchmarkShellSorted_10000(b *testing.B) { benchmarkShellSorted(10000, b) }

func BenchmarkShellDescending_1(b *testing.B)     { benchmarkShellDescending(1, b) }
func BenchmarkShellDescending_10(b *testing.B)    { benchmarkShellDescending(10, b) }
func BenchmarkShellDescending_100(b *testing.B)   { benchmarkShellDescending(100, b) }
func BenchmarkShellDescending_1000(b *testing.B)  { benchmarkShellDescending(1000, b) }
func BenchmarkShellDescending_10000(b *testing.B) { benchmarkShellDescending(10000, b) }

//benchmark merge

func benchmarkMerge(size int, b *testing.B)           { benchmarkRandom(Merge, size, b) }
func benchmarkMergeSorted(size int, b *testing.B)     { benchmarkSortedASC(Merge, size, b) }
func benchmarkMergeDescending(size int, b *testing.B) { benchmarkSortedDESC(Merge, size, b) }

func BenchmarkMergeSortInts_1(b *testing.B)     { benchmarkMerge(1, b) }
func BenchmarkMergeSortInts_10(b *testing.B)    { benchmarkMerge(10, b) }
func BenchmarkMergeSortInts_100(b *testing.B)   { benchmarkMerge(100, b) }
func BenchmarkMergeSortInts_1000(b *testing.B)  { benchmarkMerge(1000, b) }
func BenchmarkMergeSortInts_10000(b *testing.B) { benchmarkMerge(10000, b) }

func BenchmarkMergeSortIntsSorted_1(b *testing.B)     { benchmarkMergeSorted(1, b) }
func BenchmarkMergeSortIntsSorted_10(b *testing.B)    { benchmarkMergeSorted(10, b) }
func BenchmarkMergeSortIntsSorted_100(b *testing.B)   { benchmarkMergeSorted(100, b) }
func BenchmarkMergeSortIntsSorted_1000(b *testing.B)  { benchmarkMergeSorted(1000, b) }
func BenchmarkMergeSortIntsSorted_10000(b *testing.B) { benchmarkMergeSorted(10000, b) }

func BenchmarkMergeSortIntsDescending_1(b *testing.B)     { benchmarkMergeDescending(1, b) }
func BenchmarkMergeSortIntsDescending_10(b *testing.B)    { benchmarkMergeDescending(10, b) }
func BenchmarkMergeSortIntsDescending_100(b *testing.B)   { benchmarkMergeDescending(100, b) }
func BenchmarkMergeSortIntsDescending_1000(b *testing.B)  { benchmarkMergeDescending(1000, b) }
func BenchmarkMergeSortIntsDescending_10000(b *testing.B) { benchmarkMergeDescending(10000, b) }
