package random_test

import (
	"fmt"
	"github.com/vistrcm/coursera-algorithms-pi-go/random"
)

func ExampleShuffle() {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	random.Shuffle(int64(len(numbers)), func(i, j int64) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	for _, val := range numbers {
		fmt.Println(val)
	}
	// Unordered output: 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}
