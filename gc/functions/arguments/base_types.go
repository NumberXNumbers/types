package arguments

// Type is the type of the constant
type Type int

const (
	// Value is for gcv.Value and is used for constant constants, i.e 5
	Value Type = iota
	// Vector is for v.Vector
	Vector
	// Matrix is for m.Matrix
	Matrix
	// Constant is for Const
	Constant
	// Variable is for Vars
	Variable
	// Operation is for operations
	Operation
)
