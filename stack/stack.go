package stack

type OfStrings interface {
	Push(item string)
	Pop() string
	IsEmpty() bool
}

type node struct {
	Item string
	Next *node
}

type OfStringsLL struct {
	first *node
}

func NewStackOfStringsLL() *OfStringsLL {
	return &OfStringsLL{}
}

func (s *OfStringsLL) Push(item string) {
	oldFirst := s.first
	s.first = &node{
		Item: item,
		Next: oldFirst,
	}
}

func (s *OfStringsLL) Pop() string {
	item := s.first.Item
	s.first = s.first.Next
	return item
}

func (s *OfStringsLL) IsEmpty() bool {
	return s.first == nil
}
