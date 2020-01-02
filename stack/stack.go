package stack

//OfStrings interface representing stack of strings
type OfStrings interface {
	Push(item string)
	Pop() string
	IsEmpty() bool
}

type node struct {
	Item string
	Next *node
}

//OfStringsLL linked list implementation of OfStrings
type OfStringsLL struct {
	first *node
}

//NewStackOfStringsLL returns new structure OfStringsLL
func NewStackOfStringsLL() *OfStringsLL {
	return &OfStringsLL{}
}

//Push add item to the stack
func (s *OfStringsLL) Push(item string) {
	oldFirst := s.first
	s.first = &node{
		Item: item,
		Next: oldFirst,
	}
}

//Pop item from the stack
func (s *OfStringsLL) Pop() string {
	item := s.first.Item
	s.first = s.first.Next
	return item
}

//IsEmpty check if stack is empty
func (s *OfStringsLL) IsEmpty() bool {
	return s.first == nil
}

// Array implementation

type FixedCapacityStackOfStrings struct {
	s []string
	n int
}

//Push add item to the stack
func (s *FixedCapacityStackOfStrings) Push(item string) {
	s.s[s.n] = item
	s.n++
}

//Pop item from the stack
func (s *FixedCapacityStackOfStrings) Pop() string {
	s.n--
	return s.s[s.n]
}

//IsEmpty check if stack is empty
func (s *FixedCapacityStackOfStrings) IsEmpty() bool {
	return s.n == 0
}

//NewStackOfStringsLL returns new structure FixedCapacityStackOfStrings
func NewFixedCapacityStackOfStrings(n int) *FixedCapacityStackOfStrings {
	return &FixedCapacityStackOfStrings{
		n: 0,
		s: make([]string, n),
	}
}
