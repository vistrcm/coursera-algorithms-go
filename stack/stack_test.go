package stack

import (
	"reflect"
	"testing"
)

var (
	flowTests = []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "simple client",
			input: []string{"to", "be", "or", "not", "to", "-", "be", "-", "-", "that", "-", "-", "-", "is"},
			want:  []string{"to", "be", "not", "that", "or", "be"},
		},
		{
			name:  "five_1",
			input: []string{"1", "2", "3", "4", "5", "-", "-", "-", "-", "-"},
			want:  []string{"5", "4", "3", "2", "1"},
		},
		{
			name:  "five_2",
			input: []string{"1", "2", "5", "-", "3", "4", "-", "-", "-", "-"},
			want:  []string{"5", "4", "3", "2", "1"},
		},
		{
			name:  "five_3",
			input: []string{"5", "-", "1", "2", "3", "-", "4", "-", "-", "-"},
			want:  []string{"5", "3", "4", "2", "1"},
		},
		{
			name:  "five_4",
			input: []string{"5", "-", "4", "-", "3", "-", "2", "-", "1", "-"},
			want:  []string{"5", "4", "3", "2", "1"},
		},
	}
)

func TestNewStackOfStringsLL(t *testing.T) {

	for _, tt := range flowTests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			stack := NewStackOfStringsLL()

			for _, w := range tt.input {

				if w == "-" {
					got = append(got, stack.Pop())
				} else {
					stack.Push(w)
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStackOfStringsLL() on %v. Got: %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func testStack_IsEmpty(stack OfStrings, t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  bool
	}{
		{
			name:  "empty",
			input: nil,
			want:  true,
		},
		{
			name:  "nonempty",
			input: []string{"nonempty"},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, i := range tt.input {
				stack.Push(i)
			}

			if got := stack.IsEmpty(); got != tt.want {
				t.Errorf("NewStackOfStringsLL for input %v = IsEmpty() = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestOfStringsLL_IsEmpty(t *testing.T) {
	stack := NewStackOfStringsLL()
	testStack_IsEmpty(stack, t)
}

func TestFixedCapacityStackOfStrings_IsEmpty(t *testing.T) {
	stack := NewFixedCapacityStackOfStrings(10)
	testStack_IsEmpty(stack, t)
}

func TestResizingArrayStackOfStrings_IsEmpty(t *testing.T) {
	stack := NewResizingArrayStackOfStrings()
	testStack_IsEmpty(stack, t)
}

func TestNewFixedCapacityStackOfStrings(t *testing.T) {
	for _, tt := range flowTests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			stack := NewFixedCapacityStackOfStrings(10)

			for _, w := range tt.input {

				if w == "-" {
					got = append(got, stack.Pop())
				} else {
					stack.Push(w)
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStackOfStringsLL() on %v. Got: %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNewResizingArrayStackOfStrings(t *testing.T) {
	for _, tt := range flowTests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			stack := NewResizingArrayStackOfStrings()

			for _, w := range tt.input {

				if w == "-" {
					got = append(got, stack.Pop())
				} else {
					stack.Push(w)
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStackOfStringsLL() on %v. Got: %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// some benchmarks
func benchmarkStack(stack OfStrings, size int, b *testing.B) {
	// run size * push-pop operations on stack b.N
	for i := 0; i < b.N; i++ {
		for o := 0; o <= size; o++ {
			stack.Push("dummy")
		}
		for o := 0; o <= size; o++ {
			stack.Pop()
		}
	}
}

func benchmarkOfStringsLL(size int, b *testing.B) {
	stack := NewStackOfStringsLL()
	benchmarkStack(stack, size, b)
}

func BenchmarkOfStringsLL_1(b *testing.B)     { benchmarkOfStringsLL(1, b) }
func BenchmarkOfStringsLL_10(b *testing.B)    { benchmarkOfStringsLL(10, b) }
func BenchmarkOfStringsLL_100(b *testing.B)   { benchmarkOfStringsLL(100, b) }
func BenchmarkOfStringsLL_1000(b *testing.B)  { benchmarkOfStringsLL(1000, b) }
func BenchmarkOfStringsLL_10000(b *testing.B) { benchmarkOfStringsLL(10000, b) }

func benchmarkFixedCapacityStackOfStrings(size int, b *testing.B) {
	stack := NewFixedCapacityStackOfStrings(10001)
	benchmarkStack(stack, size, b)
}

func BenchmarkFixedCapacityStackOfStrings_1(b *testing.B) { benchmarkFixedCapacityStackOfStrings(1, b) }
func BenchmarkFixedCapacityStackOfStrings_10(b *testing.B) {
	benchmarkFixedCapacityStackOfStrings(10, b)
}
func BenchmarkFixedCapacityStackOfStrings_100(b *testing.B) {
	benchmarkFixedCapacityStackOfStrings(100, b)
}
func BenchmarkFixedCapacityStackOfStrings_1000(b *testing.B) {
	benchmarkFixedCapacityStackOfStrings(1000, b)
}
func BenchmarkFixedCapacityStackOfStrings_10000(b *testing.B) {
	benchmarkFixedCapacityStackOfStrings(10000, b)
}

func benchmarkResizingArrayStackOfStrings(size int, b *testing.B) {
	stack := NewResizingArrayStackOfStrings()
	benchmarkStack(stack, size, b)
}

func BenchmarkResizingArrayStackOfStrings_1(b *testing.B) { benchmarkResizingArrayStackOfStrings(1, b) }
func BenchmarkResizingArrayStackOfStrings_10(b *testing.B) {
	benchmarkResizingArrayStackOfStrings(10, b)
}
func BenchmarkResizingArrayStackOfStrings_100(b *testing.B) {
	benchmarkResizingArrayStackOfStrings(100, b)
}
func BenchmarkResizingArrayStackOfStrings_1000(b *testing.B) {
	benchmarkResizingArrayStackOfStrings(1000, b)
}
func BenchmarkResizingArrayStackOfStrings_10000(b *testing.B) {
	benchmarkResizingArrayStackOfStrings(10000, b)
}
