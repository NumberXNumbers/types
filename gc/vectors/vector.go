package vectors

import (
	"errors"

	gcv "github.com/NumberXNumbers/types/gc/values"
	gcvops "github.com/NumberXNumbers/types/gc/values/ops"
)

// Space is the space in which a vector lives. Either Column or Row space.
type Space int

const (
	// RowSpace is a space Row Vector
	RowSpace Space = iota
	// ColSpace is a space Column Vector
	ColSpace
)

// Vector is the main vector interface for real vectors
type Vector interface {
	// Returns the index of vect. If vect is not in gcv it returns -1.
	IndexOf(val gcv.Value) int

	// Returns the Value at index
	Get(index int) gcv.Value

	// Set the Value val at index
	Set(index int, val gcv.Value)

	// Return norm of vector
	Norm() gcv.Value

	// return the normalized unit vector of Vector. will return error if norm is equal to 0.
	Unit() (Vector, error)

	// Returns a new Copy of vector
	Copy() Vector

	// Returns Space of vector. Either Column or Row vector
	Space() Space

	// Returns the core value type of the vector
	Type() gcv.Type

	// Returns the len of the gcv
	Len() int

	// Transpose of vector. i.e changes a row into a column vector and vice versa
	Trans()

	// Conjugate of a vector
	Conj()

	// Conjugate Transpose of vector. i.e changes a row into a column vector and vice versa
	ConjTrans()

	// Append a value to vector
	Append(val gcv.Value)

	// Returns the elements in the vector
	Elements() gcv.Values
}

type vector struct {
	space    Space
	coreType gcv.Type
	elements gcv.Values
}

// implementation of Len method
func (v *vector) Len() int { return v.elements.Len() }

// implementation of Len method
func (v *vector) Get(index int) gcv.Value { return v.elements.Get(index) }

// implementation of Set method
func (v *vector) Set(index int, val gcv.Value) {
	if v.coreType < val.Type() {
		v.coreType = val.Type()
	}
	v.elements.Set(index, val)
}

// implementation of Space method
func (v *vector) Space() Space { return v.space }

// implementation of Type method
func (v *vector) Type() gcv.Type { return v.coreType }

// implementation of Copy method
func (v *vector) Copy() Vector { return MakeVectorAlt(v.Space(), v.Elements()) }

// implementation of Elements method
func (v *vector) Elements() gcv.Values { return v.elements }

// implementation of Append method
func (v *vector) Append(val gcv.Value) { v.Elements().Append(val) }

// implementation of Trans method
func (v *vector) Trans() {
	if v.Space() == ColSpace {
		v.space = RowSpace
	} else {
		v.space = ColSpace
	}
}

// implementation of Conj method
func (v *vector) Conj() {
	if v.Type() == gcv.Complex {
		for i := 0; i < v.Len(); i++ {
			value := v.Get(i)
			v.Set(i, gcvops.Conj(value))
		}
	}
}

// implementation of ConjTrans method
func (v *vector) ConjTrans() {
	v.Conj()
	v.Trans()
}

// implementation of Norm method
func (v *vector) Norm() gcv.Value {
	dotProduct := gcv.Zero()
	conjVector := MakeConjTransVector(v)
	for i := 0; i < v.Len(); i++ {
		dotProduct = gcvops.Add(dotProduct, gcvops.Mult(v.Get(i), conjVector.Get(i)))
	}
	return gcvops.Sqrt(dotProduct)
}

// implementation of Unit method
func (v *vector) Unit() (Vector, error) {
	norm := v.Norm()
	if norm.Complex() == 0 {
		return nil, errors.New("Norm equal to zero")
	}

	unitVector := NewVector(v.Space(), v.Len())
	for index, value := range gcv.RetrieveValues(v.Elements()) {
		unitVector.Set(index, gcvops.Div(value, norm))
	}
	return unitVector, nil
}

// implementation of IndexOf method
func (v *vector) IndexOf(val gcv.Value) int { return v.Elements().IndexOf(val) }

// NewVector returns zero vector of size length
func NewVector(space Space, length int) Vector {
	vector := new(vector)
	vector.space = space
	vector.coreType = gcv.Real
	vector.elements = gcv.NewValues(length)
	return vector
}

// MakeVectorAlt returns Vector, requires a framework gcv type
func MakeVectorAlt(space Space, elements gcv.Values) Vector {
	vector := new(vector)
	vector.space = space
	vector.coreType = elements.Type()
	vector.elements = elements.Copy()
	return vector
}

// MakeVector returns Vector, takes in a slice of interfaces.
// in an interface is not a supported type, that interface will be forced to the zero
// Value
func MakeVector(space Space, elements ...interface{}) Vector {
	values := gcv.MakeValues(elements...)
	vector := MakeVectorAlt(space, values)
	return vector
}

// MakeTransVector returns a new transpose vector of vector
func MakeTransVector(v Vector) Vector {
	transVector := v.Copy()
	transVector.Trans()
	return transVector
}

// MakeConjVector returns a new conjugate vector of vector
func MakeConjVector(v Vector) Vector {
	conjVector := v.Copy()
	conjVector.Conj()
	return conjVector
}

// MakeConjTransVector returns a new conjugate transpose vector of vector
func MakeConjTransVector(v Vector) Vector {
	conjTransVector := v.Copy()
	conjTransVector.ConjTrans()
	return conjTransVector
}
