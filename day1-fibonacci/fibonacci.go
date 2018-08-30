package fibonacci

import "errors"

func fibonacci(n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("Please provide a positive number")
	}

	previousValue := 0
	currentValue := 1

	fibonacciSequence := []int {previousValue, currentValue}

	if n == 1 {
		return fibonacciSequence, nil
	}

	counter := n - 1
	for counter > 0 {
		currentValue += previousValue
		previousValue = currentValue - previousValue

		fibonacciSequence = append(fibonacciSequence, currentValue)

		counter -= 1
	}
	return fibonacciSequence, nil
}
