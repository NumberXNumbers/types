package ops

import (
	"errors"

	m "github.com/NumberXNumbers/types/gc/matrices"
	gcv "github.com/NumberXNumbers/types/gc/values"
	gcvops "github.com/NumberXNumbers/types/gc/values/ops"
	v "github.com/NumberXNumbers/types/gc/vectors"
)

// SMult multiplies a real vector by a scalar
func SMult(scalar gcv.Value, vector v.Vector) v.Vector {
	newVector := v.NewVector(vector.Space(), vector.Len())
	for i := 0; i < vector.Len(); i++ {
		newVector.Set(i, gcv.MakeValue(scalar.Complex()*vector.Get(i).Complex()))
	}

	return newVector
}

// SDiv will divide a vector by a scalar
func SDiv(scalar gcv.Value, vector v.Vector) v.Vector {
	newVector := vector.Copy()
	for i := 0; i < newVector.Len(); i++ {
		newVector.Set(i, gcvops.Div(vector.Get(i), scalar))
	}
	return newVector
}

// Acos returns the angle theta between Vector A and Vector B using the dot product
func Acos(vectorA v.Vector, vectorB v.Vector) (gcv.Value, error) {
	normA := vectorA.Norm()
	normB := vectorB.Norm()
	var theta gcv.Value

	if normA.Real() == 0 || normB.Real() == 0 {
		return theta, errors.New("Either Vector A or Vector B is the zero vector")
	}

	dotProduct, err := InnerProduct(vectorA, vectorB)

	if err != nil {
		return theta, err
	}

	theta = gcvops.Acos(gcvops.Div(dotProduct, gcvops.Mult(normA, normB)))

	return theta, nil
}

// InnerProduct returns the inner product for real Vectors
func InnerProduct(vectorA v.Vector, vectorB v.Vector) (gcv.Value, error) {
	product := gcv.Zero()

	if vectorA.Len() != vectorB.Len() {
		return product, errors.New("Length of v does not match")
	}

	if vectorA.Space() != v.RowSpace || vectorB.Space() != v.ColSpace {
		return product, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	for i := 0; i < vectorA.Len(); i++ {
		product = gcvops.Add(product, gcvops.Mult(vectorA.Get(i), vectorB.Get(i)))
	}

	return product, nil
}

// OuterProduct returns the outer product for real Vectors
func OuterProduct(vectorA v.Vector, vectorB v.Vector) (m.Matrix, error) {
	if vectorA.Space() != v.ColSpace || vectorB.Space() != v.RowSpace {
		return nil, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	matrix := m.NewMatrix(vectorA.Len(), vectorB.Len())

	for i := 0; i < vectorA.Len(); i++ {
		for j := 0; j < vectorB.Len(); j++ {
			matrix.Set(i, j, gcvops.Mult(vectorA.Get(i), vectorB.Get(j)))
		}
	}

	return matrix, nil
}

// Add adds two real v together
func Add(vectorA v.Vector, vectorB v.Vector) (v.Vector, error) {
	if vectorA.Space() != vectorB.Space() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column v or row v")
	}

	if vectorA.Len() != vectorB.Len() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := v.NewVector(vectorA.Space(), vectorA.Len())

	for i := 0; i < vectorA.Len(); i++ {
		vector.Set(i, gcvops.Add(vectorA.Get(i), vectorB.Get(i)))
	}

	return vector, nil
}

// Sub subtracts two real v together
func Sub(vectorA v.Vector, vectorB v.Vector) (v.Vector, error) {
	if vectorA.Space() != vectorB.Space() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column v or row v")
	}

	if vectorA.Len() != vectorB.Len() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := v.NewVector(vectorA.Space(), vectorA.Len())
	for i := 0; i < vectorA.Len(); i++ {
		vector.Set(i, gcvops.Sub(vectorA.Get(i), vectorB.Get(i)))
	}

	return vector, nil
}
