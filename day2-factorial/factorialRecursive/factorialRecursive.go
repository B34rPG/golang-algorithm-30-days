package factorialRecursive

import "errors"

func factorialRecursive(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Please provide an natural number.")
	}

	if n > 1 {
		previousResult, _ := factorialRecursive(n - 1)
		return (n * previousResult), nil
	}

	return 1, nil
}
