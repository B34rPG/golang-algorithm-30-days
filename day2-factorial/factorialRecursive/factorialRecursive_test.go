package factorialRecursive

import (
	"reflect"
	"testing"
)

func TestFactorialRecursive(t *testing.T) {
	/*  factorialRecursive(0)
	===================================================================== */
	result, err := factorialRecursive(0)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	expected := 1
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be 1")
	}

	/*  factorialRecursive(8)
	===================================================================== */
	result, err = factorialRecursive(8)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	expected = 40320
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be 40320")
	}

	/*  factorialRecursive(15)
	===================================================================== */
	result, err = factorialRecursive(15)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	expected = 1307674368000
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be 1307674368000")
	}
}
