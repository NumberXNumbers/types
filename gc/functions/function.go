package functions

import (
	"errors"
	"fmt"

	args "github.com/NumberXNumbers/types/gc/functions/arguments"
	"github.com/NumberXNumbers/types/gc/functions/ops"
	m "github.com/NumberXNumbers/types/gc/matrices"
	gcv "github.com/NumberXNumbers/types/gc/values"
	v "github.com/NumberXNumbers/types/gc/vectors"
)

var (
	leftParen  = "("
	rightParen = ")"
	pow        = "^"
	unaryFuncs = map[string]func(args.Const) (args.Const, error){
		"Sqrt":  ops.Sqrt,
		"Conj":  ops.Conj,
		"Sin":   ops.Sin,
		"Cos":   ops.Cos,
		"Tan":   ops.Tan,
		"Asin":  ops.Asin,
		"Acos":  ops.Acos,
		"Atan":  ops.Atan,
		"Sinh":  ops.Sinh,
		"Cosh":  ops.Cosh,
		"Tanh":  ops.Tanh,
		"Asinh": ops.Asinh,
		"Acosh": ops.Acosh,
		"Atanh": ops.Atanh,
	}
	binaryFuncs = map[string]func(args.Const, args.Const) (args.Const, error){
		"+": ops.Add,
		"-": ops.Sub,
		"*": ops.Mult,
		"/": ops.Div,
		pow: ops.Pow,
	}
	orderOfOperations = map[string]uint{
		pow: 3,
		"*": 2,
		"/": 2,
		"+": 1,
		"-": 1,
	}
)

// Function is the function type for GoCalculate
type Function struct {
	inputTypes map[int]args.Type
	Args       []interface{}
	varNum     map[args.Var]int
	numVars    int
	regVars    []args.Var
}

func newConstVar(c args.Const) args.Var {
	return c.(args.Var)
}

func (f *Function) getVar(i int) (args.Var, error) {
	if f.typeInput(i) == args.Constant {
		return newConstVar(f.Args[i].(args.Const)), nil
	} else if f.typeInput(i) == args.Variable {
		return f.Args[i].(args.Var), nil
	}
	return nil, fmt.Errorf("Index %d, is not of type Var or Const", i)
}

func (f *Function) getOp(i int) (string, error) {
	if f.typeInput(i) == args.Operation {
		return f.Args[i].(string), nil
	}
	return "", fmt.Errorf("Index %d, is not of type Operations", i)
}

func (f *Function) typeInput(x int) args.Type { return f.inputTypes[x] }

// Eval will evaluate a function
func (f *Function) Eval(inputs ...interface{}) (args.Const, error) {
	lenInputs := len(inputs)
	if lenInputs != f.numVars {
		return nil, errors.New("Number of inputs is not equal to the number of variables in function")
	}

	var operand1 args.Const
	var operand2 args.Const
	var operandStack []args.Const

	i := 0
	for i < len(f.Args) {
		if f.typeInput(i) == args.Constant || f.typeInput(i) == args.Variable {
			variable, err := f.getVar(i)
			if err != nil {
				return nil, err
			}

			if lenInputs != 0 {
				operand, err := variable.Eval(inputs[f.varNum[variable]])
				if err != nil {
					return nil, err
				}

				operandStack = append(operandStack, operand)
			} else {
				// If length inputs is 0, then all variables must be constant.
				// This code assumes variable is a constant and so uses 0 as an input
				// to MustEval as it will never fail as the input does not matter for constants
				operandStack = append(operandStack, variable.MustEval(0))
			}
		} else if f.typeInput(i) == args.Operation {
			operation, err := f.getOp(i)
			if err != nil {
				return nil, err
			}

			if h, ok := unaryFuncs[operation]; ok {
				if len(operandStack) == 0 {
					return nil, errors.New("Not enough operands")
				}

				operand1, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
				result, err := h(operand1)
				if err != nil {
					return nil, err
				}

				operandStack = append(operandStack, result)
			} else if h, ok := binaryFuncs[operation]; ok {
				if len(operandStack) < 2 {
					return nil, errors.New("Not enough operands")
				}

				operand2, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
				operand1, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
				result, err := h(operand1, operand2)
				if err != nil {
					return nil, err
				}

				operandStack = append(operandStack, result)
			} else {
				return nil, errors.New("Operation not supported")
			}
		}
		i++
	}

	if len(operandStack) > 1 {
		return nil, errors.New("To many operands left over after calculation")
	}

	return operandStack[0], nil
}

// MustEval is like Eval but will panic
func (f *Function) MustEval(inputs ...interface{}) args.Const {
	constant, err := f.Eval(inputs...)
	if err != nil {
		panic(err)
	}
	return constant
}

// MakeFunc will make a gcf function struct.
// Else error
func MakeFunc(regVars []args.Var, inputs ...interface{}) (*Function, error) {
	function := new(Function)

	function.regVars = regVars
	var varNum = make(map[args.Var]int)
	var numVars int
	var tempOpsStack []string
	var postfixStack []interface{}
	for i, v := range regVars {
		if _, ok := varNum[v]; !ok {
			varNum[v] = numVars
			numVars++
			continue
		}
		return nil, fmt.Errorf("Error registering variables. Variable at index %d, is a duplicate", i)

	}
	var inputType = make(map[int]args.Type)
	for i, n := range inputs {
		topIndexInPostfixStack := len(postfixStack) - 1
		switch n.(type) {
		case string:
			operation := n.(string)
			var finishComparing bool
			topIndexInTempOpsStack := len(tempOpsStack) - 1
			if len(tempOpsStack) == 0 ||
				(tempOpsStack[topIndexInTempOpsStack] == leftParen && operation != rightParen) {
				tempOpsStack = append(tempOpsStack, operation)
			} else if operation == leftParen {
				tempOpsStack = append(tempOpsStack, operation)
			} else if operation == rightParen {
				for !finishComparing {
					if len(tempOpsStack) == 0 {
						return nil, errors.New("Mismatch of Parentheses found")
					}
					topOperationInTempOpsStack := tempOpsStack[topIndexInTempOpsStack]
					if topOperationInTempOpsStack == leftParen {
						tempOpsStack = tempOpsStack[:topIndexInTempOpsStack]
						finishComparing = true
					} else {
						inputType[topIndexInPostfixStack+1] = args.Operation
						postfixStack, tempOpsStack = append(postfixStack, topOperationInTempOpsStack), tempOpsStack[:topIndexInTempOpsStack]
					}
					topIndexInTempOpsStack = len(tempOpsStack) - 1
					topIndexInPostfixStack = len(postfixStack) - 1
				}
			} else {
				topOperationInTempOpsStack := tempOpsStack[topIndexInTempOpsStack]
				var isPreviousUnary bool
				var isUnary bool
				if _, ok := unaryFuncs[topOperationInTempOpsStack]; ok {
					isPreviousUnary = true
				}
				if _, ok := unaryFuncs[operation]; ok {
					isUnary = true
				}
				if isPreviousUnary || orderOfOperations[operation] < orderOfOperations[topOperationInTempOpsStack] {
					for !finishComparing {
						if isUnary && isPreviousUnary {
							tempOpsStack = append(tempOpsStack, operation)
							finishComparing = true
						} else if (topOperationInTempOpsStack == leftParen ||
							orderOfOperations[operation] > orderOfOperations[topOperationInTempOpsStack] ||
							isUnary) &&
							!isPreviousUnary {
							tempOpsStack = append(tempOpsStack, operation)
							finishComparing = true
						} else if orderOfOperations[operation] == orderOfOperations[topOperationInTempOpsStack] {
							if operation == pow {
								tempOpsStack = append(tempOpsStack, operation)
								finishComparing = true
							} else {
								inputType[topIndexInPostfixStack+1] = args.Operation
								postfixStack, tempOpsStack = append(postfixStack, topOperationInTempOpsStack), tempOpsStack[:topIndexInTempOpsStack]
								topIndexInTempOpsStack = len(tempOpsStack) - 1
							}
						} else if orderOfOperations[operation] < orderOfOperations[topOperationInTempOpsStack] || isPreviousUnary {
							inputType[topIndexInPostfixStack+1] = args.Operation
							postfixStack, tempOpsStack = append(postfixStack, topOperationInTempOpsStack), tempOpsStack[:topIndexInTempOpsStack]
							topIndexInTempOpsStack = len(tempOpsStack) - 1
						}

						if len(tempOpsStack) == 0 {
							tempOpsStack = append(tempOpsStack, operation)
							finishComparing = true
						} else {
							topOperationInTempOpsStack = tempOpsStack[topIndexInTempOpsStack]
							topIndexInPostfixStack = len(postfixStack) - 1
							if _, ok := unaryFuncs[topOperationInTempOpsStack]; !ok {
								isPreviousUnary = false
							}
						}
					}
				} else if orderOfOperations[operation] > orderOfOperations[topOperationInTempOpsStack] {
					tempOpsStack = append(tempOpsStack, operation)
				} else if orderOfOperations[operation] == orderOfOperations[topOperationInTempOpsStack] {
					if operation == pow {
						tempOpsStack = append(tempOpsStack, operation)
					} else {
						inputType[topIndexInPostfixStack+1] = args.Operation
						postfixStack, tempOpsStack = append(postfixStack, topOperationInTempOpsStack), tempOpsStack[:topIndexInTempOpsStack]
						tempOpsStack = append(tempOpsStack, operation)
					}
				}
			}
		case int, int32, int64, float32, float64, complex64, complex128, gcv.Value, v.Vector, m.Matrix:
			postfixStack = append(postfixStack, args.MakeConst(inputs[i]))
			inputType[topIndexInPostfixStack+1] = args.Constant
		case args.Const:
			postfixStack = append(postfixStack, n)
			inputType[topIndexInPostfixStack+1] = args.Constant
		case args.Var:
			if _, ok := varNum[n.(args.Var)]; !ok {
				return nil, fmt.Errorf("Variable at index %d, was not registered", i)

			}
			postfixStack = append(postfixStack, n)
			inputType[topIndexInPostfixStack+1] = args.Variable
		default:
			return nil, errors.New("Input type not supported")
		}
	}

	for len(tempOpsStack) > 0 {
		topIndexInTempOpsStack := len(tempOpsStack) - 1
		topIndexInPostfixStack := len(postfixStack) - 1
		var operation string
		operation, tempOpsStack = tempOpsStack[topIndexInTempOpsStack], tempOpsStack[:topIndexInTempOpsStack]
		if operation == "(" {
			return nil, errors.New("Mismatch of Parentheses found")
		}
		inputType[topIndexInPostfixStack+1] = args.Operation
		postfixStack = append(postfixStack, operation)
	}

	function.inputTypes = inputType
	function.numVars = numVars
	function.varNum = varNum
	function.Args = postfixStack
	return function, nil
}

// MakeFuncPanic will the same as MakeFunc but will panic
func MakeFuncPanic(regVars []args.Var, inputs ...interface{}) *Function {
	function, err := MakeFunc(regVars, inputs...)
	if err != nil {
		panic(err)
	}
	return function
}
