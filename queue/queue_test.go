package queue_test

import (
	"fmt"
	"github.com/vistrcm/coursera-algorithms-go/queue"
)

func ExampleOfStringsLL() {
	q := queue.NewOfStringsLL()

	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")
	q.Enqueue("one more")

	fmt.Println(q.Length())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println(q.Length())

	fmt.Println(q.Dequeue())
	fmt.Println(q.Length())

	// Output: 4
	// first
	// second
	// third
	// 1
	// one more
	// 0
}
