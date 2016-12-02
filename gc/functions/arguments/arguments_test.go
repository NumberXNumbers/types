package arguments

import (
	"fmt"
	"reflect"
	"testing"

	m "github.com/NumberXNumbers/types/gc/matrices"
	gcv "github.com/NumberXNumbers/types/gc/values"
	v "github.com/NumberXNumbers/types/gc/vectors"
)

func TestMakeConst(t *testing.T) {
	c1 := MakeConst(3)

	if c1.Type() != Value {
		t.Fail()
	}

	solution1 := c1.Value()

	if solution1 == nil || solution1.Real() != 3 {
		t.Errorf("Expected %v, received %v", 3, solution1.Real())
	}

	c2 := MakeConst(v.NewVector(v.RowSpace, 2))

	if c2.Type() != Vector {
		t.Fail()
	}

	solution2 := c2.Vector()

	if solution2 == nil || solution2.Get(0).Complex() != 0 || solution2.Get(1).Complex() != 0 {
		t.Errorf("Expected %v, received %v", 3, solution1.Real())
	}

	c3 := MakeConst(m.NewIdentityMatrix(2))

	if c3.Type() != Matrix {
		t.Fail()
	}

	solution3 := c3.Matrix()

	if solution3 == nil ||
		solution3.Get(0, 0).Complex() != 1 || solution3.Get(0, 1).Complex() != 0 ||
		solution3.Get(1, 0).Complex() != 0 || solution3.Get(1, 1).Complex() != 1 {
		t.Errorf("Expected %v, received %v", 3, solution1.Real())
	}

	c4 := MakeConst(c1)
	if !reflect.DeepEqual(c4, c1) {
		t.Fail()
	}

	c5 := MakeConst("0")
	if c5.Type() != Value {
		t.Fail()
	}
}

func TestConstantEvalAndMustEval(t *testing.T) {
	c1 := new(constant)

	if c2, err := c1.Eval(4); err != nil || !reflect.DeepEqual(c2, c1) {
		t.Fail()
	}

	if !reflect.DeepEqual(c1.MustEval(4), c1) {
		t.Fail()
	}
}

func TestPanicBadValue(t *testing.T) {
	v := MakeConst(v.NewVector(v.RowSpace, 2))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := v.Value()

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadVector(t *testing.T) {
	v := MakeConst(m.NewMatrix(2, 2))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := v.Vector()

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadMatrix(t *testing.T) {
	m := MakeConst(gcv.NewValue())

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := m.Matrix()

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadVariable(t *testing.T) {
	m := NewVar(Matrix)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := m.MustEval(gcv.NewValue())

	if solution != nil {
		t.Error("Expected Panic")
	}
}
