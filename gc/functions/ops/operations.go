package ops

import (
	"errors"

	args "github.com/NumberXNumbers/types/gc/functions/arguments"
	m "github.com/NumberXNumbers/types/gc/matrices"
	mops "github.com/NumberXNumbers/types/gc/matrices/ops"
	gcvops "github.com/NumberXNumbers/types/gc/values/ops"
	v "github.com/NumberXNumbers/types/gc/vectors"
	vops "github.com/NumberXNumbers/types/gc/vectors/ops"
)

// Add will add two constants together
// For vectors, the two vectors need to be in the same space and size, else error.
// For matrices, the two matrices need to be of the same size, else error.
func Add(constA args.Const, constB args.Const) (args.Const, error) {
	if constA.Type() == args.Value && constB.Type() == args.Value {
		return args.MakeConst(gcvops.Add(constA.Value(), constB.Value())), nil
	}

	if constA.Type() == args.Vector && constB.Type() == args.Vector {
		vector, err := vops.Add(constA.Vector(), constB.Vector())
		if err != nil {
			return nil, err
		}
		return args.MakeConst(vector), nil
	}

	if constA.Type() == args.Matrix && constB.Type() == args.Matrix {
		matrix, err := mops.Add(constA.Matrix(), constB.Matrix())
		if err != nil {
			return nil, err
		}
		return args.MakeConst(matrix), nil
	}
	return nil, errors.New("One or More Types are not supported")
}

// MustAdd is the same as Add but will panic
func MustAdd(constA args.Const, constB args.Const) args.Const {
	constant, err := Add(constA, constB)
	if err != nil {
		panic(err)
	}
	return constant
}

// Sub will subtract two constants together
func Sub(constA args.Const, constB args.Const) (args.Const, error) {
	if constA.Type() == args.Value && constB.Type() == args.Value {
		return args.MakeConst(gcvops.Sub(constA.Value(), constB.Value())), nil
	}
	if constA.Type() == args.Vector && constB.Type() == args.Vector {
		vector, err := vops.Sub(constA.Vector(), constB.Vector())
		if err != nil {
			return nil, err
		}
		return args.MakeConst(vector), nil
	}
	if constA.Type() == args.Matrix && constB.Type() == args.Matrix {
		matrix, err := mops.Sub(constA.Matrix(), constB.Matrix())
		if err != nil {
			return nil, err
		}
		return args.MakeConst(matrix), nil
	}
	return nil, errors.New("One or More Types are not supported")
}

// MustSub is the same as Sub but will panic
func MustSub(constA args.Const, constB args.Const) args.Const {
	constant, err := Sub(constA, constB)
	if err != nil {
		panic(err)
	}
	return constant
}

// Div will divide two constants together
func Div(constA args.Const, constB args.Const) (args.Const, error) {
	if constA.Type() == args.Value && constB.Type() == args.Value {
		return args.MakeConst(gcvops.Div(constA.Value(), constB.Value())), nil
	}
	if constA.Type() == args.Vector && constB.Type() == args.Value {
		return args.MakeConst(vops.SDiv(constB.Value(), constA.Vector())), nil
	}
	if constA.Type() == args.Matrix && constB.Type() == args.Value {
		return args.MakeConst(mops.SDiv(constB.Value(), constA.Matrix())), nil
	}
	return nil, errors.New("One or More Types are not supported")
}

// MustDiv is the same as Div but will panic
func MustDiv(constA args.Const, constB args.Const) args.Const {
	constant, err := Div(constA, constB)
	if err != nil {
		panic(err)
	}
	return constant
}

// Mult will multiply two constants together
func Mult(constA args.Const, constB args.Const) (args.Const, error) {
	if constA.Type() == args.Value && constB.Type() == args.Value {
		return args.MakeConst(gcvops.Mult(constA.Value(), constB.Value())), nil
	}
	if constA.Type() == args.Vector && constB.Type() == args.Value {
		return args.MakeConst(vops.SMult(constB.Value(), constA.Vector())), nil
	}
	if constA.Type() == args.Value && constB.Type() == args.Vector {
		return args.MakeConst(vops.SMult(constA.Value(), constB.Vector())), nil
	}
	if constA.Type() == args.Vector && constB.Type() == args.Vector {
		vectorA := constA.Vector()
		vectorB := constB.Vector()
		if vectorA.Space() == v.RowSpace {
			vector, err := vops.InnerProduct(vectorA, vectorB)
			if err != nil {
				return nil, err
			}
			return args.MakeConst(vector), nil
		}
		matrix, err := vops.OuterProduct(vectorA, vectorB)
		if err != nil {
			return nil, err
		}
		return args.MakeConst(matrix), nil
	}
	if constA.Type() == args.Matrix && constB.Type() == args.Value {
		return args.MakeConst(mops.SMult(constB.Value(), constA.Matrix())), nil
	}
	if constA.Type() == args.Value && constB.Type() == args.Matrix {
		return args.MakeConst(mops.SMult(constA.Value(), constB.Matrix())), nil
	}
	if constA.Type() == args.Vector && constB.Type() == args.Matrix {
		vector, err := mops.VMMult(constA.Vector(), constB.Matrix())
		if err != nil {
			return nil, err
		}
		return args.MakeConst(vector), nil
	}
	if constA.Type() == args.Matrix && constB.Type() == args.Vector {
		vector, err := mops.MVMult(constB.Vector(), constA.Matrix())
		if err != nil {
			return nil, err
		}
		return args.MakeConst(vector), nil
	}
	matrix, err := mops.MultSimple(constA.Matrix(), constB.Matrix())
	if err != nil {
		return nil, err
	}
	return args.MakeConst(matrix), nil
}

// MustMult is the same as Mult but will panic
func MustMult(constA args.Const, constB args.Const) args.Const {
	constant, err := Mult(constA, constB)
	if err != nil {
		panic(err)
	}
	return constant
}

// Pow will raise one constant to the power of another constant
// for matrix consts, it is assumed that the value it will be raised to is an integer
func Pow(constA args.Const, constB args.Const) (args.Const, error) {
	if constA.Type() == args.Value && constB.Type() == args.Value {
		return args.MakeConst(gcvops.Pow(constA.Value(), constB.Value())), nil
	}
	if constA.Type() == args.Matrix && constB.Type() == args.Value {
		matrix, err := mops.Pow(constA.Matrix(), int(constB.Value().Real()))
		if err != nil {
			return nil, err
		}
		return args.MakeConst(matrix), nil
	}
	return nil, errors.New("One or More Types are not supported")
}

// MustPow is the same as Pow but will panic
func MustPow(constA args.Const, constB args.Const) args.Const {
	con, err := Pow(constA, constB)
	if err != nil {
		panic(err)
	}
	return con
}

// Sqrt will find the square root of a Const
func Sqrt(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Sqrt(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Sqrt")
}

// MustSqrt is the same as Sqrt but will panic
func MustSqrt(constant args.Const) args.Const {
	con, err := Sqrt(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Conj will find the conjuage of a Const
func Conj(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Conj(constant.Value())), nil
	}
	if constant.Type() == args.Vector {
		return args.MakeConst(v.MakeConjVector(constant.Vector())), nil
	}
	if constant.Type() == args.Matrix {
		return args.MakeConst(m.MakeConjMatrix(constant.Matrix())), nil
	}
	return nil, errors.New("Const Type is not supported for Conj")
}

// MustConj is the same as Conj but will panic
func MustConj(constant args.Const) args.Const {
	con, err := Conj(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Sin will find the sine of a Const
func Sin(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Sin(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Sin")
}

// MustSin is the same as Sin but will panic
func MustSin(constant args.Const) args.Const {
	con, err := Sin(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Cos will find the cosine of a Const
func Cos(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Cos(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Cos")
}

// MustCos is the same as Cos but will panic
func MustCos(constant args.Const) args.Const {
	con, err := Cos(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Tan will find the tangent of a Const
func Tan(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Tan(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Tan")
}

// MustTan is the same as Tan but will panic
func MustTan(constant args.Const) args.Const {
	con, err := Tan(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Asin will find the arcsine of a Const
func Asin(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Asin(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Asin")
}

// MustAsin is the same as Asin but will panic
func MustAsin(constant args.Const) args.Const {
	con, err := Asin(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Acos will find the arccosine of a Const
func Acos(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Acos(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Acos")
}

// MustAcos is the same as Acos but will panic
func MustAcos(constant args.Const) args.Const {
	con, err := Acos(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Atan will find the arctangent of a Const
func Atan(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Atan(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Atan")
}

// MustAtan is the same as Atan but will panic
func MustAtan(constant args.Const) args.Const {
	con, err := Atan(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Sinh will find the hyperbolicSine of a Const
func Sinh(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Sinh(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Sinh")
}

// MustSinh is the same as Sinh but will panic
func MustSinh(constant args.Const) args.Const {
	con, err := Sinh(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Cosh will find the hyperbolicCosine of a Const
func Cosh(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Cosh(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Cosh")
}

// MustCosh is the same as Cosh but will panic
func MustCosh(constant args.Const) args.Const {
	con, err := Cosh(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Tanh will find the hyperbolicTangent of a Const
func Tanh(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Tanh(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Tanh")
}

// MustTanh is the same as Tanh but will panic
func MustTanh(constant args.Const) args.Const {
	con, err := Tanh(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Asinh will find the hyperbolicArcSine of a Const
func Asinh(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Asinh(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Asinh")
}

// MustAsinh is the same as Asinh but will panic
func MustAsinh(constant args.Const) args.Const {
	con, err := Asinh(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Acosh will find the hyperbolicArcCosine of a Const
func Acosh(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Acosh(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Acosh")
}

// MustAcosh is the same as Acosh but will panic
func MustAcosh(constant args.Const) args.Const {
	con, err := Acosh(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Atanh will find the hyperbolicArcTangent of a Const
func Atanh(constant args.Const) (args.Const, error) {
	if constant.Type() == args.Value {
		return args.MakeConst(gcvops.Atanh(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Atanh")
}

// MustAtanh is the same as Atanh but will panic
func MustAtanh(constant args.Const) args.Const {
	con, err := Atanh(constant)
	if err != nil {
		panic(err)
	}
	return con
}
