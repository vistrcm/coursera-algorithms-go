package random_test

import (
	"fmt"
	"github.com/vistrcm/coursera-algorithms-go/random"
	"math/rand"
)

func ExampleShuffle() {

	rand.Seed(42)

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	random.Shuffle(int64(len(numbers)), func(i, j int64) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	for _, val := range numbers {
		fmt.Println(val)
	}
	// Output: 7
	// 6
	// 4
	// 5
	// 9
	// 1
	// 3
	// 2
	// 8
	// 10
	// 0
}
