package ops

import (
	"fmt"
	"math"
	"math/cmplx"
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

func TestConj(t *testing.T) {
	solutionA := MustConj(args.MakeConst(3 + 5i))
	if solutionA.Value().Complex() != 3-5i {
		t.Fail()
	}

	solutionB := MustConj(args.MakeConst(v.MakeVector(v.RowSpace, 2+1i, -3-3i)))
	if solutionB.Vector().Get(0).Complex() != 2-1i ||
		solutionB.Vector().Get(1).Complex() != -3+3i {
		t.Fail()
	}

	solutionC := MustConj(args.MakeConst(m.MakeMatrix(v.MakeVector(v.RowSpace, 2+1i, -3-3i), v.MakeVector(v.RowSpace, 4+2i, -1i))))
	if solutionC.Matrix().Get(0, 0).Complex() != 2-1i ||
		solutionC.Matrix().Get(0, 1).Complex() != -3+3i ||
		solutionC.Matrix().Get(1, 0).Complex() != 4-2i ||
		solutionC.Matrix().Get(1, 1).Complex() != 1i {
		t.Fail()
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

func TestSin(t *testing.T) {
	solutionA := MustSin(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-0) >= 1e-10 {
		t.Fail()
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

func TestCos(t *testing.T) {
	solutionA := MustCos(args.MakeConst(math.Pi / 2.0))
	if cmplx.Abs(solutionA.Value().Complex()-0) >= 1e-10 {
		t.Fail()
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

func TestTan(t *testing.T) {
	solutionA := MustTan(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-0) >= 1e-10 {
		t.Fail()
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

func TestAtan(t *testing.T) {
	solutionA := MustAtan(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-1.26262726) >= 1e-6 {
		t.Fail()
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

func TestAcos(t *testing.T) {
	solutionA := MustAcos(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-1.81152627) >= 1e-6 {
		t.Fail()
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

func TestAsin(t *testing.T) {
	solutionA := MustAsin(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-1.57079632) >= 1e-6 {
		t.Fail()
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

func TestSinh(t *testing.T) {
	solutionA := MustSinh(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-11.5487393) >= 1e-6 {
		t.Fail()
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

func TestCosh(t *testing.T) {
	solutionA := MustCosh(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-11.5919532) >= 1e-6 {
		t.Fail()
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

func TestTanh(t *testing.T) {
	solutionA := MustTanh(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-0.9962720) >= 1e-6 {
		t.Fail()
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

func TestAtanh(t *testing.T) {
	solutionA := MustAtanh(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-0.3297653) >= 1e-6 {
		t.Fail()
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

func TestAcosh(t *testing.T) {
	solutionA := MustAcosh(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-1.8115262) >= 1e-6 {
		t.Fail()
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

func TestAsinh(t *testing.T) {
	solutionA := MustAsinh(args.MakeConst(math.Pi))
	if cmplx.Abs(solutionA.Value().Complex()-1.862295) >= 1e-6 {
		t.Fail()
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
