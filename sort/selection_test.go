package sort

import (
	"math/rand"
	"reflect"
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

func TestSort(t *testing.T) {
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
			Selection(IntSlice(a))
			got := a
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection() for input %v. got %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

// some benchmarks
func benchmarkSearch(size int, b *testing.B) {
	// prepare big slice
	bigSlice := generateRandomIntSlice(size)

	// run size * push-pop operations on stack b.N
	for i := 0; i < b.N; i++ {
		for o := 0; o <= size; o++ {
			Selection(IntSlice(bigSlice))
		}
	}
}

func BenchmarkSearch_1(b *testing.B)    { benchmarkSearch(1, b) }
func BenchmarkSearch_10(b *testing.B)   { benchmarkSearch(10, b) }
func BenchmarkSearch_100(b *testing.B)  { benchmarkSearch(100, b) }
func BenchmarkSearch_1000(b *testing.B) { benchmarkSearch(1000, b) }

//func BenchmarkSearch_10000(b *testing.B)    { benchmarkSearch(10000, b) }