package mops

import (
	"testing"

	m "github.com/NumberXNumbers/types/gc/matrices"
	gcv "github.com/NumberXNumbers/types/gc/values"
	v "github.com/NumberXNumbers/types/gc/vectors"
)

func BenchmarkSMult(b *testing.B) {
	testVectorAa := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testVectorAb := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(2))
	testMatrix := m.MakeMatrix(testVectorAa, testVectorAb)

	testScalarA := gcv.MakeValue(2.0)

	for n := 0; n < b.N; n++ {
		SMult(testScalarA, testMatrix)
	}
}

func BenchmarkVMMult(b *testing.B) {
	testVectorA := v.MakeVector(v.RowSpace, gcv.MakeValue(1), gcv.MakeValue(1))
	testMatrix := m.MakeMatrix(testVectorA, testVectorA)

	for n := 0; n < b.N; n++ {
		VMMult(testVectorA, testMatrix)
	}
}
