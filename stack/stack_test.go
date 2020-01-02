package stack

import (
	"reflect"
	"testing"
)

func TestNewStackOfStringsLL(t *testing.T) {
	tests := []struct {
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
	for _, tt := range tests {
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

func TestOfStringsLL_IsEmpty(t *testing.T) {
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
			stack := NewStackOfStringsLL()
			for _, i := range tt.input {
				stack.Push(i)
			}

			if got := stack.IsEmpty(); got != tt.want {
				t.Errorf("NewStackOfStringsLL for input %v = IsEmpty() = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
