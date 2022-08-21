// queue implementation
package queue

type OfStrings interface {
	//Enqueue inserts an item into the queue
	Enqueue(item string)
	//Dequeue removes and returns item least recently added
	Dequeue() string
	//IsEmpty check if the queue is empty
	IsEmpty() bool
	//Length returns the number of items in the queue
	Length() int
}

type node struct {
	item string
	next *node
}

// OfStringsLL linked list implementation of queue
type OfStringsLL struct {
	first *node
	last  *node
	len   int
}

// NewOfStringsLL create new instance of OfStringsLL
func NewOfStringsLL() *OfStringsLL {
	return &OfStringsLL{}
}

// Enqueue inserts an item into the queue
func (q *OfStringsLL) Enqueue(item string) {
	q.len++
	oldLast := q.last
	//new node for the end of the list
	q.last = &node{
		item: item,
		next: nil,
	}
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
}

// Dequeue removes and returns item least recently added
func (q *OfStringsLL) Dequeue() string {
	q.len--
	item := q.first.item
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	return item
}

// IsEmpty check if the queue is empty
func (q *OfStringsLL) IsEmpty() bool {
	return q.first == nil
}

// Length returns the number of items in the queue
func (q *OfStringsLL) Length() int {
	return q.len
}
