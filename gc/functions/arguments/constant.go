package arguments

import (
	m "github.com/NumberXNumbers/types/gc/matrices"
	gcv "github.com/NumberXNumbers/types/gc/values"
	v "github.com/NumberXNumbers/types/gc/vectors"
)

// Const is the GoCalculate constant type
type Const interface {
	// Type will return the type of the constant
	Type() Type

	// Matrix will return a matrix if the type is matrix, else will panic
	Matrix() m.Matrix

	// Vector will return a vector if the type is vector, else will panic
	Vector() v.Vector

	// Value will return a value if the type is value, else will panic
	Value() gcv.Value
}

type constant struct {
	constType Type
	constant  interface{}
}

func (c *constant) Type() Type { return c.constType }

func (c *constant) Matrix() m.Matrix {
	if c.Type() != Matrix {
		panic("constant is not of type Matrix")
	}
	return c.constant.(m.Matrix)
}

func (c *constant) Vector() v.Vector {
	if c.Type() != Vector {
		panic("constant is not of type Vector")
	}
	return c.constant.(v.Vector)
}

func (c *constant) Value() gcv.Value {
	if c.Type() != Value {
		panic("constant is not of type Value")
	}
	return c.constant.(gcv.Value)
}

// Eval is mostly just to make sure the constant meets the requirements for Var interface as well as Const
func (c *constant) Eval(x interface{}) (Const, error) {
	return c, nil
}

// MustEval is mostly just to make sure constant meets the requirements for Var interface as well as Const
func (c *constant) MustEval(x interface{}) Const {
	return c
}

// MakeConst will take an interface of type Value, Vector or Matrix
// And will return a constant type. if the type is not a supported Type, the
// zero Value will be returned.
func MakeConst(c interface{}) Const {
	constant := new(constant)
	switch c.(type) {
	case Const:
		return c.(Const)
	case m.Matrix:
		constant.constType = Matrix
	case v.Vector:
		constant.constType = Vector
	case gcv.Value:
		constant.constType = Value
	case int, int32, int64, float32, float64, complex64, complex128:
		c = gcv.MakeValue(c)
		constant.constType = Value
	default:
		c = gcv.Zero()
		constant.constType = Value
	}
	constant.constant = c
	return constant
}
