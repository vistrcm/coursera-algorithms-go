package priorityQueue_test

import (
	"fmt"
	"github.com/vistrcm/coursera-algorithms-pi-go/priorityQueue"
)

func ExampleUnorderedMaxPQ() {
	q := priorityQueue.NewUnorderedMaxPQ(10)
	q.Insert(priorityQueue.Item{"P"})
	q.Insert(priorityQueue.Item{"Q"})
	q.Insert(priorityQueue.Item{"E"})
	fmt.Println(q.DelMax().Value) //Q

	q.Insert(priorityQueue.Item{"X"})
	q.Insert(priorityQueue.Item{"A"})
	q.Insert(priorityQueue.Item{"M"})
	fmt.Println(q.DelMax().Value) //X

	q.Insert(priorityQueue.Item{"P"})
	q.Insert(priorityQueue.Item{"L"})
	q.Insert(priorityQueue.Item{"E"})
	fmt.Println(q.DelMax().Value) //P
	// Output: Q
	// X
	// P
}
