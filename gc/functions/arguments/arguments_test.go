package arguments

import (
	"fmt"
	"testing"

	m "github.com/NumberXNumbers/types/gc/matrices"
	gcv "github.com/NumberXNumbers/types/gc/values"
	v "github.com/NumberXNumbers/types/gc/vectors"
)

func TestMakeConst(t *testing.T) {
	c := MakeConst(3)

	solution := c.Value()

	if solution == nil || solution.Real() != 3 {
		t.Errorf("Expected %v, received %v", 3, solution.Real())
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
