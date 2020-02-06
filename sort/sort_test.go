package sort

import (
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
func stdSortWrapper(slice sort.IntSlice) {
	sort.Sort(slice)
}

// benchmark Selection

func benchmarkSelection(size int, b *testing.B) { benchmarkRandom(onIntSlice(Selection), size, b) }
func benchmarkSelectionSorted(size int, b *testing.B) {
	benchmarkSortedASC(onIntSlice(Selection), size, b)
}
func benchmarkSelectionDescending(size int, b *testing.B) {
	benchmarkSortedDESC(onIntSlice(Selection), size, b)
}

// benchmark Insertion

func benchmarkInsertion(size int, b *testing.B) { benchmarkRandom(onIntSlice(Insertion), size, b) }
func benchmarkInsertionSorted(size int, b *testing.B) {
	benchmarkSortedASC(onIntSlice(Insertion), size, b)
}
func benchmarkInsertionDescending(size int, b *testing.B) {
	benchmarkSortedDESC(onIntSlice(Insertion), size, b)
}

// benchmark Shell
func benchmarkShell(size int, b *testing.B)           { benchmarkRandom(onIntSlice(Shell), size, b) }
func benchmarkShellSorted(size int, b *testing.B)     { benchmarkSortedASC(onIntSlice(Shell), size, b) }
func benchmarkShellDescending(size int, b *testing.B) { benchmarkSortedDESC(onIntSlice(Shell), size, b) }

//benchmark merge

func benchmarkMerge(size int, b *testing.B)           { benchmarkRandom(Merge, size, b) }
func benchmarkMergeSorted(size int, b *testing.B)     { benchmarkSortedASC(Merge, size, b) }
func benchmarkMergeDescending(size int, b *testing.B) { benchmarkSortedDESC(Merge, size, b) }

//benchmark merge bottom up

func benchmarkMergeBU(size int, b *testing.B)           { benchmarkRandom(MergeBU, size, b) }
func benchmarkMergeBUSorted(size int, b *testing.B)     { benchmarkSortedASC(MergeBU, size, b) }
func benchmarkMergeBUDescending(size int, b *testing.B) { benchmarkSortedDESC(MergeBU, size, b) }

//benchmark quicksort

func benchmarkQuick(size int, b *testing.B)           { benchmarkRandom(Quick, size, b) }
func benchmarkQuickSorted(size int, b *testing.B)     { benchmarkSortedASC(Quick, size, b) }
func benchmarkQuickDescending(size int, b *testing.B) { benchmarkSortedDESC(Quick, size, b) }

func Benchmark_10_STD(b *testing.B)       { benchmarkRandom(stdSortWrapper, 10, b) }
func Benchmark_10_Selection(b *testing.B) { benchmarkSelection(10, b) }
func Benchmark_10_Insertion(b *testing.B) { benchmarkInsertion(10, b) }
func Benchmark_10_Shell(b *testing.B)     { benchmarkShell(10, b) }
func Benchmark_10_Merge(b *testing.B)     { benchmarkMerge(10, b) }
func Benchmark_10_MergeBU(b *testing.B)   { benchmarkMergeBU(10, b) }
func Benchmark_10_Quick(b *testing.B)     { benchmarkQuick(10, b) }

func Benchmark_100_STD(b *testing.B)       { benchmarkRandom(stdSortWrapper, 100, b) }
func Benchmark_100_Selection(b *testing.B) { benchmarkSelection(100, b) }
func Benchmark_100_Insertion(b *testing.B) { benchmarkInsertion(100, b) }
func Benchmark_100_Shell(b *testing.B)     { benchmarkShell(100, b) }
func Benchmark_100_Merge(b *testing.B)     { benchmarkMerge(100, b) }
func Benchmark_100_MergeBU(b *testing.B)   { benchmarkMergeBU(100, b) }
func Benchmark_100_Quick(b *testing.B)     { benchmarkQuick(100, b) }

func Benchmark_1000_STD(b *testing.B)       { benchmarkRandom(stdSortWrapper, 1000, b) }
func Benchmark_1000_Selection(b *testing.B) { benchmarkSelection(1000, b) }
func Benchmark_1000_Insertion(b *testing.B) { benchmarkInsertion(1000, b) }
func Benchmark_1000_Shell(b *testing.B)     { benchmarkShell(1000, b) }
func Benchmark_1000_Merge(b *testing.B)     { benchmarkMerge(1000, b) }
func Benchmark_1000_MergeBU(b *testing.B)   { benchmarkMergeBU(1000, b) }
func Benchmark_1000_Quick(b *testing.B)     { benchmarkQuick(1000, b) }

func Benchmark_10000_STD(b *testing.B)       { benchmarkRandom(stdSortWrapper, 10000, b) }
func Benchmark_10000_Insertion(b *testing.B) { benchmarkInsertion(10000, b) }
func Benchmark_10000_Shell(b *testing.B)     { benchmarkShell(10000, b) }
func Benchmark_10000_Merge(b *testing.B)     { benchmarkMerge(10000, b) }
func Benchmark_10000_MergeBU(b *testing.B)   { benchmarkMergeBU(10000, b) }
func Benchmark_10000_Quick(b *testing.B)     { benchmarkQuick(10000, b) }

func Benchmark_10_ASC_STD(b *testing.B)       { benchmarkSortedASC(stdSortWrapper, 10, b) }
func Benchmark_10_ASC_Selection(b *testing.B) { benchmarkSelectionSorted(10, b) }
func Benchmark_10_ASC_Insertion(b *testing.B) { benchmarkInsertionSorted(10, b) }
func Benchmark_10_ASC_Shell(b *testing.B)     { benchmarkShellSorted(10, b) }
func Benchmark_10_ASC_Merge(b *testing.B)     { benchmarkMergeSorted(10, b) }
func Benchmark_10_ASC_MergeBU(b *testing.B)   { benchmarkMergeBUSorted(10, b) }
func Benchmark_10_ASC_Quick(b *testing.B)     { benchmarkQuickSorted(10, b) }

func Benchmark_100_ASC_STD(b *testing.B)       { benchmarkSortedASC(stdSortWrapper, 100, b) }
func Benchmark_100_ASC_Selection(b *testing.B) { benchmarkSelectionSorted(100, b) }
func Benchmark_100_ASC_Insertion(b *testing.B) { benchmarkInsertionSorted(100, b) }
func Benchmark_100_ASC_Shell(b *testing.B)     { benchmarkShellSorted(100, b) }
func Benchmark_100_ASC_Merge(b *testing.B)     { benchmarkMergeSorted(100, b) }
func Benchmark_100_ASC_MergeBU(b *testing.B)   { benchmarkMergeBUSorted(100, b) }
func Benchmark_100_ASC_Quick(b *testing.B)     { benchmarkQuickSorted(100, b) }

func Benchmark_1000_ASC_STD(b *testing.B)       { benchmarkSortedASC(stdSortWrapper, 1000, b) }
func Benchmark_1000_ASC_Selection(b *testing.B) { benchmarkSelectionSorted(1000, b) }
func Benchmark_1000_ASC_Insertion(b *testing.B) { benchmarkInsertionSorted(1000, b) }
func Benchmark_1000_ASC_Shell(b *testing.B)     { benchmarkShellSorted(1000, b) }
func Benchmark_1000_ASC_Merge(b *testing.B)     { benchmarkMergeSorted(1000, b) }
func Benchmark_1000_ASC_MergeBU(b *testing.B)   { benchmarkMergeBUSorted(1000, b) }
func Benchmark_1000_ASC_Quick(b *testing.B)     { benchmarkQuickSorted(1000, b) }

func Benchmark_10000_ASC_STD(b *testing.B)       { benchmarkSortedASC(stdSortWrapper, 10000, b) }
func Benchmark_10000_ASC_Insertion(b *testing.B) { benchmarkInsertionSorted(10000, b) }
func Benchmark_10000_ASC_Shell(b *testing.B)     { benchmarkShellSorted(10000, b) }
func Benchmark_10000_ASC_Merge(b *testing.B)     { benchmarkMergeSorted(10000, b) }
func Benchmark_10000_ASC_MergeBU(b *testing.B)   { benchmarkMergeBUSorted(10000, b) }
func Benchmark_10000_ASC_Quick(b *testing.B)     { benchmarkQuickSorted(10000, b) }

func Benchmark_10_DESC_STD(b *testing.B)       { benchmarkSortedDESC(stdSortWrapper, 10, b) }
func Benchmark_10_DESC_Selection(b *testing.B) { benchmarkSelectionDescending(10, b) }
func Benchmark_10_DESC_Insertion(b *testing.B) { benchmarkInsertionDescending(10, b) }
func Benchmark_10_DESC_Shell(b *testing.B)     { benchmarkShellDescending(10, b) }
func Benchmark_10_DESC_Merge(b *testing.B)     { benchmarkMergeDescending(10, b) }
func Benchmark_10_DESC_MergeBU(b *testing.B)   { benchmarkMergeBUDescending(10, b) }
func Benchmark_10_DESC_Quick(b *testing.B)     { benchmarkQuickDescending(10, b) }

func Benchmark_100_DESC_STD(b *testing.B)       { benchmarkSortedDESC(stdSortWrapper, 100, b) }
func Benchmark_100_DESC_Selection(b *testing.B) { benchmarkSelectionDescending(100, b) }
func Benchmark_100_DESC_Insertion(b *testing.B) { benchmarkInsertionDescending(100, b) }
func Benchmark_100_DESC_Shell(b *testing.B)     { benchmarkShellDescending(100, b) }
func Benchmark_100_DESC_Merge(b *testing.B)     { benchmarkMergeDescending(100, b) }
func Benchmark_100_DESC_MergeBU(b *testing.B)   { benchmarkMergeBUDescending(100, b) }
func Benchmark_100_DESC_Quick(b *testing.B)     { benchmarkQuickDescending(100, b) }

func Benchmark_1000_DESC_STD(b *testing.B)       { benchmarkSortedDESC(stdSortWrapper, 1000, b) }
func Benchmark_1000_DESC_Selection(b *testing.B) { benchmarkSelectionDescending(1000, b) }
func Benchmark_1000_DESC_Insertion(b *testing.B) { benchmarkInsertionDescending(1000, b) }
func Benchmark_1000_DESC_Shell(b *testing.B)     { benchmarkShellDescending(1000, b) }
func Benchmark_1000_DESC_Merge(b *testing.B)     { benchmarkMergeDescending(1000, b) }
func Benchmark_1000_DESC_MergeBU(b *testing.B)   { benchmarkMergeBUDescending(1000, b) }
func Benchmark_1000_DESC_Quick(b *testing.B)     { benchmarkQuickDescending(1000, b) }

func Benchmark_10000_DESC_STD(b *testing.B)       { benchmarkSortedDESC(stdSortWrapper, 10000, b) }
func Benchmark_10000_DESC_Insertion(b *testing.B) { benchmarkInsertionDescending(10000, b) }
func Benchmark_10000_DESC_Shell(b *testing.B)     { benchmarkShellDescending(10000, b) }
func Benchmark_10000_DESC_Merge(b *testing.B)     { benchmarkMergeDescending(10000, b) }
func Benchmark_10000_DESC_MergeBU(b *testing.B)   { benchmarkMergeBUDescending(10000, b) }
func Benchmark_10000_DESC_Quick(b *testing.B)     { benchmarkQuickDescending(10000, b) }
