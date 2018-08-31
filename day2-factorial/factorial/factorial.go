package factorial

import "errors"

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Please provide an natural number.")
	}

	result := 1
	for i := 2; i <= n; i += 1 {
		result *= i
	}

	return result, nil
}
