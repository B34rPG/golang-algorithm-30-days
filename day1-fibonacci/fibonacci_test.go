package fibonacci

import (
	"testing"
	"reflect"
)

func TestFibonacci(t *testing.T) {
	/*  fibonacci(1)
  	===================================================================== */
	result, err := fibonacci(1)
	if err != nil { t.Fatal("got errors:", err) }

	expected := []int {0, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be [0, 1, 1, 2, 3, 5, 8]")
	}

	/*  fibonacci(2)
  	===================================================================== */
	result, err = fibonacci(2)
	if err != nil { t.Fatal("got errors:", err) }

	expected = []int {0, 1, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be [0, 1, 1]")
	}

	/*  fibonacci(3)
  	===================================================================== */
	  result, err = fibonacci(3)
	  if err != nil { t.Fatal("got errors:", err) }

	  expected = []int {0, 1, 1, 2}
	  if !reflect.DeepEqual(result, expected) {
		  t.Error("should be [0, 1, 1, 2]")
	  }

	/*  fibonacci(4)
  	===================================================================== */
	result, err = fibonacci(4)
	if err != nil { t.Fatal("got errors:", err) }

	expected = []int {0, 1, 1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be [0, 1, 1, 2, 3]")
	}

	/*  fibonacci(5)
  	===================================================================== */
	result, err = fibonacci(5)
	if err != nil { t.Fatal("got errors:", err) }

	expected = []int {0, 1, 1, 2, 3, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be [0, 1, 1, 2, 3, 5]")
	}

	/*  fibonacci(6)
  	===================================================================== */
	result, err = fibonacci(6)
	if err != nil { t.Fatal("got errors:", err) }

	expected = []int {0, 1, 1, 2, 3, 5, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be [0, 1, 1, 2, 3, 5, 8]")
	}

	/*  fibonacci(7)
  	===================================================================== */
	result, err = fibonacci(7)
	if err != nil { t.Fatal("got errors:", err) }

	expected = []int {0, 1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be [0, 1, 1, 2, 3, 5, 8, 13]")
	}

	/*  fibonacci(8)
  	===================================================================== */
	result, err = fibonacci(8)
	if err != nil { t.Fatal("got errors:", err) }

	expected = []int {0, 1, 1, 2, 3, 5, 8, 13, 21}
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be [0, 1, 1, 2, 3, 5, 8, 13, 21]")
	}
}
