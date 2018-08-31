package factorial

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	/*  factorial(0)
	===================================================================== */
	result, err := factorial(0)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	expected := 1
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be 1")
	}

	/*  factorial(8)
	===================================================================== */
	result, err = factorial(8)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	expected = 40320
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be 40320")
	}

	/*  factorial(15)
	===================================================================== */
	result, err = factorial(15)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	expected = 1307674368000
	if !reflect.DeepEqual(result, expected) {
		t.Error("should be 1307674368000")
	}
}
