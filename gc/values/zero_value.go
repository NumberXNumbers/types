package values

type zeroValue struct{}

const (
	zeroReal    = 0.0
	zeroComplex = 0 + 0i
	zeroType    = Real
	isZero      = true
)

var zero *zeroValue

func init() {
	zero = new(zeroValue)
}

func (v *zeroValue) Real() float64 { return zeroReal }

func (v *zeroValue) Imag() float64 { return zeroReal }

func (v *zeroValue) Complex() complex128 { return zeroComplex }

func (v *zeroValue) Type() Type { return zeroType }

// IsZero will always return true
func (v *zeroValue) IsZero() bool { return isZero }

// Zero will return the zeroValue type as a Value
func Zero() Value {
	return zero
}
