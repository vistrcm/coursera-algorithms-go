package queue_test

import (
	"fmt"
	"github.com/vistrcm/coursera-algorithms-pi-go/queue"
)

func ExampleOfStringsLL() {
	q := queue.NewOfStringsLL()

	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")
	q.Enqueue("one more")

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	// Output: first
	// second
	// third
}
