package ops

import (
	"fmt"
	"testing"

	args "github.com/NumberXNumbers/types/gc/functions/arguments"
	m "github.com/NumberXNumbers/types/gc/matrices"
	gcv "github.com/NumberXNumbers/types/gc/values"
	v "github.com/NumberXNumbers/types/gc/vectors"
)

func TestAdd(t *testing.T) {
	solutionA := MustAdd(args.MakeConst(4), args.MakeConst(3))
	if solutionA.Value().Complex() != 7 {
		t.Fail()
	}

	solutionB := MustAdd(args.MakeConst(v.NewVector(v.RowSpace, 3)), args.MakeConst(v.MakeVector(v.RowSpace, 1, 2, 3)))
	if solutionB.Vector().Get(0).Complex() != 1 ||
		solutionB.Vector().Get(1).Complex() != 2 ||
		solutionB.Vector().Get(2).Complex() != 3 {
		t.Fail()
	}

	solutionC := MustAdd(args.MakeConst(m.NewMatrix(2, 2)), args.MakeConst(m.NewIdentityMatrix(2)))
	if solutionC.Matrix().Get(0, 0).Complex() != 1 ||
		solutionC.Matrix().Get(1, 1).Complex() != 1 ||
		solutionC.Matrix().Get(0, 1).Complex() != 0 ||
		solutionC.Matrix().Get(1, 0).Complex() != 0 {
		t.Fail()
	}
}

func TestPanicAddVector(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))
	v2 := args.MakeConst(v.NewVector(v.ColSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solutionV := MustAdd(v1, v2)

	if solutionV != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicAddMatrix(t *testing.T) {
	m1 := args.MakeConst(m.NewMatrix(2, 2))
	m2 := args.MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solutionM := MustAdd(m1, m2)

	if solutionM != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicAddMismatch(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))
	m2 := args.MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustAdd(v1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestSub(t *testing.T) {
	solutionA := MustSub(args.MakeConst(4), args.MakeConst(3))
	if solutionA.Value().Complex() != 1 {
		t.Fail()
	}

	solutionB := MustSub(args.MakeConst(v.NewVector(v.RowSpace, 3)), args.MakeConst(v.MakeVector(v.RowSpace, 1, 2, 3)))
	if solutionB.Vector().Get(0).Complex() != -1 ||
		solutionB.Vector().Get(1).Complex() != -2 ||
		solutionB.Vector().Get(2).Complex() != -3 {
		t.Fail()
	}

	solutionC := MustSub(args.MakeConst(m.NewMatrix(2, 2)), args.MakeConst(m.NewIdentityMatrix(2)))
	if solutionC.Matrix().Get(0, 0).Complex() != -1 ||
		solutionC.Matrix().Get(1, 1).Complex() != -1 ||
		solutionC.Matrix().Get(0, 1).Complex() != 0 ||
		solutionC.Matrix().Get(1, 0).Complex() != 0 {
		t.Fail()
	}
}

func TestPanicSubVector(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))
	v2 := args.MakeConst(v.NewVector(v.ColSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solutionV := MustSub(v1, v2)

	if solutionV != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicSubMatrix(t *testing.T) {
	m1 := args.MakeConst(m.NewMatrix(2, 2))
	m2 := args.MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solutionM := MustSub(m1, m2)

	if solutionM != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicSubMismatch(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))
	m2 := args.MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustSub(v1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestDiv(t *testing.T) {
	solutionA := MustDiv(args.MakeConst(3), args.MakeConst(3))
	if solutionA.Value().Complex() != 1 {
		t.Fail()
	}

	solutionB := MustDiv(args.MakeConst(v.MakeVector(v.RowSpace, 4, 2, 4)), args.MakeConst(2))
	if solutionB.Vector().Get(0).Complex() != 2 ||
		solutionB.Vector().Get(1).Complex() != 1 ||
		solutionB.Vector().Get(2).Complex() != 2 {
		t.Fail()
	}

	solutionC := MustDiv(args.MakeConst(m.NewIdentityMatrix(2)), args.MakeConst(2))
	if solutionC.Matrix().Get(0, 0).Complex() != 0.5 ||
		solutionC.Matrix().Get(1, 1).Complex() != 0.5 ||
		solutionC.Matrix().Get(0, 1).Complex() != 0 ||
		solutionC.Matrix().Get(1, 0).Complex() != 0 {
		t.Fail()
	}
}

func TestPanicDivMismatch(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))
	m2 := args.MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustDiv(v1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestMult(t *testing.T) {
	solutionA := MustMult(args.MakeConst(3), args.MakeConst(3))
	if solutionA.Value().Complex() != 9 {
		t.Fail()
	}

	solutionB := MustMult(args.MakeConst(v.MakeVector(v.RowSpace, 4, 2, 4)), args.MakeConst(2))
	if solutionB.Vector().Get(0).Complex() != 8 ||
		solutionB.Vector().Get(1).Complex() != 4 ||
		solutionB.Vector().Get(2).Complex() != 8 {
		t.Fail()
	}

	solutionC := MustMult(args.MakeConst(m.NewIdentityMatrix(2)), args.MakeConst(2))
	if solutionC.Matrix().Get(0, 0).Complex() != 2 ||
		solutionC.Matrix().Get(1, 1).Complex() != 2 ||
		solutionC.Matrix().Get(0, 1).Complex() != 0 ||
		solutionC.Matrix().Get(1, 0).Complex() != 0 {
		t.Fail()
	}

	solutionD := MustMult(args.MakeConst(2), args.MakeConst(v.MakeVector(v.RowSpace, 4, 2, 4)))
	if solutionD.Vector().Get(0).Complex() != 8 ||
		solutionD.Vector().Get(1).Complex() != 4 ||
		solutionD.Vector().Get(2).Complex() != 8 {
		t.Fail()
	}

	solutionE := MustMult(args.MakeConst(2), args.MakeConst(m.NewIdentityMatrix(2)))
	if solutionE.Matrix().Get(0, 0).Complex() != 2 ||
		solutionE.Matrix().Get(1, 1).Complex() != 2 ||
		solutionE.Matrix().Get(0, 1).Complex() != 0 ||
		solutionE.Matrix().Get(1, 0).Complex() != 0 {
		t.Fail()
	}
}

func TestPanicMultDoubleRowVector(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))
	v2 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustMult(v1, v2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultDoubleColVector(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.ColSpace, 3))
	v2 := args.MakeConst(v.NewVector(v.ColSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustMult(v1, v2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultVM(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.ColSpace, 3))
	m2 := args.MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustMult(v1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultMV(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))
	m2 := args.MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustMult(m2, v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicMultMatrix(t *testing.T) {
	m1 := args.MakeConst(m.NewMatrix(2, 2))
	m2 := args.MakeConst(m.NewMatrix(3, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustMult(m1, m2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPow(t *testing.T) {
	solutionA := MustPow(args.MakeConst(3), args.MakeConst(3))
	if solutionA.Value().Complex() != 27 {
		t.Fail()
	}

	solutionC := MustPow(args.MakeConst(m.NewIdentityMatrix(2)), args.MakeConst(2))
	if solutionC.Matrix().Get(0, 0).Complex() != 1 ||
		solutionC.Matrix().Get(1, 1).Complex() != 1 ||
		solutionC.Matrix().Get(0, 1).Complex() != 0 ||
		solutionC.Matrix().Get(1, 0).Complex() != 0 {
		t.Fail()
	}
}

func TestPanicBadPowMatrix(t *testing.T) {
	v1 := args.MakeConst(m.NewMatrix(2, 3))
	v2 := args.MakeConst(2)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustPow(v1, v2)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadPowUnsupportedType(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustPow(v1, v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestSqrt(t *testing.T) {
	solutionA := MustSqrt(args.MakeConst(9))
	if solutionA.Value().Complex() != 3 {
		t.Fail()
	}
}

func TestPanicBadSqrt(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustSqrt(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

type testConst struct{}

func (c *testConst) Type() args.Type { return args.Constant }

func (c *testConst) Matrix() m.Matrix { return nil }

func (c *testConst) Vector() v.Vector { return nil }

func (c *testConst) Value() gcv.Value { return nil }

func TestPanicBadConj(t *testing.T) {
	v1 := &testConst{}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustConj(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadSin(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustSin(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadCos(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustCos(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadTan(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustTan(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadAtan(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustAtan(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadAcos(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustAcos(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadAsin(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustAsin(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadSinh(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustSinh(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadCosh(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustCosh(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadTanh(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustTanh(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadAtanh(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustAtanh(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadAcosh(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustAcosh(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadAsinh(t *testing.T) {
	v1 := args.MakeConst(v.NewVector(v.RowSpace, 3))

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	solution := MustAsinh(v1)

	if solution != nil {
		t.Error("Expected Panic")
	}
}
