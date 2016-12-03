package functions

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

func TestFn0(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, x)
	value := function.MustEval(4)
	if value.Value().Real() != 4 {
		t.Fail()
	}
}

func TestFn1(t *testing.T) {
	x := args.NewVar(args.Value)
	y := args.NewVar(args.Value)
	regVars := []args.Var{x, y}
	constant := args.MakeConst(0)
	function := MakeFuncPanic(regVars, args.MakeConst(4), "+", 5, "-", y, "*", args.MakeConst(3), "/", args.MakeConst(7), "+", x, "+", args.MakeConst(constant), "+", args.MakeConst("0"))
	value := function.MustEval(4, 3)
	if value.Value().Real() != 11.714285714285715 {
		t.Fail()
	}
}

func TestFn2(t *testing.T) {
	x := args.NewVar(args.Value)
	y := args.NewVar(args.Matrix)
	regVars := []args.Var{x, y}
	vector := v.MakeVector(v.RowSpace, 2, 4, 6)
	constVect := args.MakeConst(vector)
	function := MakeFuncPanic(regVars, constVect, "*", y, "*", x, "/", args.MakeConst(4))
	matrix := m.NewIdentityMatrix(3)
	value := function.MustEval(2, matrix)
	if value.Vector().Get(0).Real() != 1 ||
		value.Vector().Get(1).Real() != 2 ||
		value.Vector().Get(2).Real() != 3 {
		t.Fail()
	}
}

func TestFn3(t *testing.T) {
	x := args.NewVar(args.Vector)
	regVars := []args.Var{x}
	vector := v.MakeVector(v.RowSpace, 2, 4, 6)
	function := MakeFuncPanic(regVars, x, "+", "(", args.MakeConst(5), "*", x, ")")
	value := function.MustEval(vector)
	if value.Vector().Get(0).Real() != 12 ||
		value.Vector().Get(1).Real() != 24 ||
		value.Vector().Get(2).Real() != 36 {
		t.Fail()
	}
}

func TestFn4(t *testing.T) {
	x := args.NewVar(args.Matrix)
	y := args.NewVar(args.Matrix)
	a := args.NewVar(args.Vector)
	b := args.NewVar(args.Vector)
	regVars := []args.Var{x, y, a, b}
	matrixA := m.NewIdentityMatrix(3)
	matrixB := m.NewIdentityMatrix(3)
	vectorA := v.MakeVector(v.RowSpace, 2, 4, 6)
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)
	function := MakeFuncPanic(regVars, a, "*", "(", x, "+", "(", y, "*", args.MakeConst(2), "-", x, ")", "/", args.MakeConst(2), ")", "-", b, "/", args.MakeConst(2))
	value := function.MustEval(matrixA, matrixB, vectorA, vectorB)
	if value.Vector().Get(0).Real() != 2 ||
		value.Vector().Get(1).Real() != 4 ||
		value.Vector().Get(2).Real() != 6 {
		t.Fail()
	}
}

func TestFn5(t *testing.T) {
	x := args.NewVar(args.Matrix)
	y := args.NewVar(args.Matrix)
	a := args.NewVar(args.Vector)
	b := args.NewVar(args.Vector)
	regVars := []args.Var{x, y, a, b}
	matrixA := m.NewIdentityMatrix(2)
	matrixB := m.NewIdentityMatrix(2)
	vectorA := v.MakeVector(v.RowSpace, 1, 0)
	vectorB := v.MakeVector(v.ColSpace, 0, 1)
	function := MakeFuncPanic(regVars, a, "*", b, "*", x, "+", y, "+", b, "*", a)

	value := function.MustEval(matrixA, matrixB, vectorA, vectorB)
	if value.Matrix().Get(0, 0).Real() != 1 ||
		value.Matrix().Get(0, 1).Real() != 0 ||
		value.Matrix().Get(1, 0).Real() != 1 ||
		value.Matrix().Get(1, 1).Real() != 1 {
		t.Fail()
	}
}

func TestFn6(t *testing.T) {
	x := args.NewVar(args.Matrix)
	y := args.NewVar(args.Matrix)
	a := args.NewVar(args.Vector)
	regVars := []args.Var{x, y, a}
	matrixA := m.NewIdentityMatrix(2)
	matrixB := m.NewIdentityMatrix(2)
	vectorA := v.MakeVector(v.ColSpace, 1, 0)
	function := MakeFuncPanic(regVars, x, "*", y, "*", a)
	value := function.MustEval(matrixA, matrixB, vectorA)
	if value.Vector().Get(0).Real() != 1 ||
		value.Vector().Get(1).Real() != 0 {
		t.Fail()
	}
}

func TestFn7(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, "Sin", "(", x, ")")
	value := function.MustEval(math.Pi)
	if value.Value().Real() >= 10e-15 {
		t.Fail()
	}
}

func TestFn8(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, "Sin", x)
	value := function.MustEval(math.Pi)
	if value.Value().Real() >= 10e-15 {
		t.Fail()
	}
}

func TestFn9(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, x, "*", "Sin", "(", x, ")", "+", x)
	// fmt.Println(function.args)
	// fmt.Println(function.inputTypes)
	value := function.MustEval(math.Pi)
	if math.Abs(value.Value().Real()-3.1415926535897936) > 10e-15 {
		t.Fail()
	}
}

func TestFn10(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, "Sqrt", "(", x, "^", x, ")")
	value := function.MustEval(2)
	if math.Abs(value.Value().Real()-2) > 10e-15 {
		t.Fail()
	}
}

func TestFn11(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, "Cos", "(", x, ")", "*", "Sin", "(", x, ")")
	value := function.MustEval(math.Pi / 4)
	if math.Abs(value.Value().Real()-0.5000000) > 10e-15 {
		t.Fail()
	}
}

func TestFn12(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, "Cos", "(", "Sin", "(", x, ")", ")")
	value := function.MustEval(math.Pi / 4)
	if math.Abs(value.Value().Real()-0.760244) > 10e-6 {
		t.Fail()
	}
}

func TestFn13(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, "Cos", "Sin", x)
	value := function.MustEval(math.Pi / 4)
	if math.Abs(value.Value().Real()-0.760244) > 10e-6 {
		t.Fail()
	}
}

func TestFn14(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, x, pow, "Sin", math.Pi/2, pow, x)
	value := function.MustEval(2)
	if math.Abs(value.Value().Real()-2) > 10e-6 {
		t.Fail()
	}
}

func TestFn15(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	function := MakeFuncPanic(regVars, x, pow, 3, pow, x)
	value := function.MustEval(2)
	if math.Abs(value.Value().Real()-512) > 10e-6 {
		t.Fail()
	}
}

func TestFn16(t *testing.T) {
	x := args.NewVar(args.Value)
	matA := m.NewIdentityMatrix(2)
	val := 1 + 3i
	vect := v.NewVector(v.RowSpace, 2)
	regVars := []args.Var{x}
	functionA := MakeFuncPanic(regVars, "Conj", x, "*", matA, "^", 2)
	matrixSolutionA := functionA.MustEval(val)
	if matrixSolutionA.Matrix().Get(0, 0).Complex() != 1-3i {
		t.Fail()
	}

	matB := matA.Trim(0, 0, 1, 0)

	functionB := MakeFuncPanic(regVars, "Conj", x, "*", matB, "^", 2)
	_, err := functionB.Eval(val)
	if err == nil {
		t.Fail()
	}

	functionC := MakeFuncPanic(regVars, "Conj", vect, "*", "Conj", matA)
	matrixSolutionC := functionC.MustEval(val)
	if matrixSolutionC.Vector().Get(0).Complex() != 0 || matrixSolutionC.Vector().Get(1).Complex() != 0 {
		t.Fail()
	}
}

func TestFn17(t *testing.T) {
	x := args.NewVar(args.Value)
	val := 1
	regVarsA := []args.Var{x}
	functionA := MakeFuncPanic(regVarsA, "Atan", x)
	solutionA := functionA.MustEval(val)
	if cmplx.Abs(solutionA.Value().Complex()-0.78539816) > 10e-6 {
		t.Fail()
	}

	functionB := MakeFuncPanic(regVarsA, "Asin", x)
	solutionB := functionB.MustEval(val)
	if cmplx.Abs(solutionB.Value().Complex()-1.570796326) > 10e-6 {
		t.Fail()
	}

	functionC := MakeFuncPanic(regVarsA, "Acos", x)
	solutionC := functionC.MustEval(val)
	if cmplx.Abs(solutionC.Value().Complex()-0) > 10e-6 {
		t.Fail()
	}

	functionD := MakeFuncPanic(regVarsA, "Tan", x)
	solutionD := functionD.MustEval(val)
	if cmplx.Abs(solutionD.Value().Complex()-1.557407724) > 10e-6 {
		t.Fail()
	}

}

func TestMustCalculateA(t *testing.T) {
	matrixA := m.NewIdentityMatrix(3)
	matrixB := m.NewIdentityMatrix(3)
	vectorA := v.MakeVector(v.RowSpace, 2, 4, 6)
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)
	calculation := MustCalculate(vectorA, "*", "(", matrixA, "+", "(", matrixB, "*", args.MakeConst(2), "-", matrixA, ")", "/", 2, ")", "-", vectorB, "/", gcv.MakeValue(2))
	if calculation.Vector().Get(0).Real() != 2 ||
		calculation.Vector().Get(1).Real() != 4 ||
		calculation.Vector().Get(2).Real() != 6 {
		t.Fail()
	}
}

func TestMustCalculateB(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)
	calculation := MustCalculate("(", args.MakeConst(2), ")", "*", vectorB)
	if calculation.Vector().Get(0).Real() != 4 ||
		calculation.Vector().Get(1).Real() != 8 ||
		calculation.Vector().Get(2).Real() != 12 {
		t.Fail()
	}
}

func TestMustCalculateC(t *testing.T) {
	calculation := MustCalculate("Sin", "(", math.Pi, ")")
	if calculation.Value().Real() >= 10e-15 {
		t.Fail()
	}
}

func TestMustCalculateD(t *testing.T) {
	calculation := MustCalculate(2, pow, 2, pow, 3)
	if calculation.Value().Real() == 8 {
		t.Fail()
	}
}

func TestMustCalculateE(t *testing.T) {
	calculation := MustCalculate("Cos", "(", "Sin", "(", math.Pi/4, ")", ")")
	if math.Abs(calculation.Value().Real()-0.760244) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateF(t *testing.T) {
	calculation := MustCalculate(math.Pi, "*", "Sin", "(", 2, "*", math.Pi, "-", math.Pi, ")", "+", math.Pi)
	// fmt.Println(function.args)
	// fmt.Println(function.inputTypes)
	if math.Abs(calculation.Value().Real()-3.1415926535897936) > 10e-15 {
		t.Fail()
	}
}

func TestMustCalculateG(t *testing.T) {
	calculation := MustCalculate(2, "*", 2, "/", 2, "*", 2)
	if math.Abs(calculation.Value().Real()-4) > 10e-15 {
		t.Fail()
	}
}

func TestMustCalculateH(t *testing.T) {
	calculation := MustCalculate("Cos", "Sin", math.Pi/4)
	if math.Abs(calculation.Value().Real()-0.760244) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateI(t *testing.T) {
	calculation := MustCalculate(2, pow, "Sin", math.Pi/2, pow, 2)
	if math.Abs(calculation.Value().Real()-2) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateJ(t *testing.T) {
	calculation := MustCalculate(2, "+", "Sin", math.Pi/2, "+", 2)
	if math.Abs(calculation.Value().Real()-5) > 10e-6 {
		t.Fail()
	}
}

func TestMustCalculateK(t *testing.T) {
	calculation := MustCalculate("Sqrt", "(", "Sin", math.Pi/2, pow, 2, "+", "Cos", math.Pi/2, pow, 2, ")")
	if math.Abs(calculation.Value().Real()-1) > 10e-6 {
		t.Fail()
	}
}

func TestFunctionPanicOperatorNotSupported(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	regVars := []args.Var{}
	function := MakeFuncPanic(regVars, "(", args.MakeConst(2), ")", "=", vectorB)

	value := function.MustEval()

	if value.Value() != nil {
		t.Error("Expected Panic")
	}
}

func TestFunctionPanicOperatorParensMismatch(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	regVars := []args.Var{}
	function := MakeFuncPanic(regVars, 2, "+", args.MakeConst(2), ")", "+", vectorB)

	value := function.MustEval()

	if value.Value() != nil {
		t.Error("Expected Panic")
	}
}

func TestMustCalculatePanicOperatorsOperandMismatch(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	calculation := MustCalculate("(", args.MakeConst(2), ")", ")", "*", vectorB)

	if calculation != nil {
		t.Error("Expected Panic")
	}
}

func TestMustCalculatePanicOperatorParensMismatch(t *testing.T) {
	vectorB := v.MakeVector(v.RowSpace, 2, 4, 6)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	calculation := MustCalculate(2, "+", args.MakeConst(2), ")", "+", vectorB)

	if calculation.Value() != nil {
		t.Error("Expected Panic")
	}
}

func TestMustCalculatePanicUnsupportedType(t *testing.T) {
	vectorB := args.NewVar(args.Vector)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	calculation := MustCalculate("(", args.MakeConst(2), ")", "*", vectorB)

	if calculation != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicDuplicateRegVarsForFunc(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x, x}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFuncPanic(regVars, args.MakeConst(4), "+", args.MakeConst(5))

	if function != nil {
		t.Error("Expected Panic")
	}

}

func TestPanicNotRegVarsForFunc(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}
	y := args.NewVar(args.Value)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFuncPanic(regVars, y)

	if function != nil {
		t.Error("Expected Panic")
	}

}

func TestPanicNotSupportedTypeForFunc(t *testing.T) {
	x := uint(2)
	regVars := []args.Var{}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFuncPanic(regVars, x)

	if function != nil {
		t.Error("Expected Panic")
	}

}

func TestPanicNotEnoughArgsFunc(t *testing.T) {
	x := args.NewVar(args.Matrix)
	regVars := []args.Var{x}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v error\n", r)
		}
	}()

	function := MakeFuncPanic(regVars, x)

	function.MustEval()

	if function != nil {
		t.Error("Expected Panic")
	}
}

func TestPanicBadgetOpFunc(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}

	function := MakeFuncPanic(regVars, x)

	_, err := function.getOp(0)

	if err == nil {
		t.Error("Expected Error")
	}
}

func TestPanicBadgetVarFunc(t *testing.T) {
	x := args.NewVar(args.Value)
	regVars := []args.Var{x}

	function := MakeFuncPanic(regVars, x, "+", x)

	_, err := function.getVar(2)

	if err == nil {
		t.Error("Expected Error")
	}
}
