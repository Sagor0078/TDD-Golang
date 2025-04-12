
package concurrency

func SquareWorker(numbers []int) []int {
	results := make(chan int, len(numbers))


	for _, num := range numbers {
		go func(n int) {
			results <- n * n
		}(num)
	}

	var squares []int

	for i := 0; i < len(numbers); i++ {
		squares = append(squares, <- results)
	}

	return squares
}