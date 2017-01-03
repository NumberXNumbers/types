package values

import (
	"reflect"
	"testing"
)

func TestZero(t *testing.T) {
	myZero := Zero()
	if !reflect.DeepEqual(zero, myZero) {
		t.Errorf("Expected %v, received %v", zero, myZero)
	}

	if !myZero.IsZero() {
		t.Errorf("Expected %t, received %t", true, myZero.IsZero())
	}

	if myZero.Real() != 0 {
		t.Errorf("Expected %f, received %f", 0.0, myZero.Real())
	}

	if myZero.Imag() != 0 {
		t.Errorf("Expected %f, received %f", 0.0, myZero.Imag())
	}

	if myZero.Complex() != 0 {
		t.Errorf("Expected %v, received %v", 0+0i, myZero.Complex())
	}

	if myZero.Type() != Real {
		t.Errorf("Expected %v, received %v", Real, myZero.Type())
	}
}
