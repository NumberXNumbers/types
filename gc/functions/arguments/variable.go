package arguments

import (
	"fmt"
)

var (
	varTypeMap = map[Type]string{
		Value:    "Value",
		Vector:   "Vector",
		Matrix:   "Matrix",
		Constant: "Constant",
	}
)

// Var is the GoCalculate variable type
type Var interface {
	// Eval will take a variable and return a constant.
	// the variable type, either type value, vector, matrix or const
	// must match the passed in variable type, else Eval will panic
	Eval(x interface{}) (Const, error)

	// MustEval is the same as Eval but will panic
	MustEval(x interface{}) Const
}

type variable struct {
	varType Type
}

func (v *variable) Eval(x interface{}) (Const, error) {
	constant := MakeConst(x)
	if v.varType != constant.Type() {
		return nil, fmt.Errorf("Expected %v, received %v", varTypeMap[v.varType], varTypeMap[constant.Type()])
	}
	return constant, nil
}

func (v *variable) MustEval(x interface{}) Const {
	constant, err := v.Eval(x)
	if err != nil {
		panic(err)
	}
	return constant
}

// NewVar will make a new Variable of type varType
func NewVar(varType Type) Var {
	variable := new(variable)
	variable.varType = varType
	return variable
}
