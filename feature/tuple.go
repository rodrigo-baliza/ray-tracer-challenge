package feature

import (
	"errors"
	"math"
)

var (
	ErrAddTwoPoints       = errors.New("can't add two points")
	ErrSubPointFromVector = errors.New("can't sub point from vector")
	ErrDivByZero          = errors.New("can't div by zero")
	ErrNotVector          = errors.New("it's not a vector")
)

// Tuple is the primitive type for any ray tracing operations.
// A Tuple can be a point if it has w = 1.0 or a vector if it has w = 0.0.
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewPoint creates a new point tuple.
func NewPoint(x, y, z float64) Tuple {
	return new(x, y, z, 1.0)
}

// NewVector creates a new vector tuple.
func NewVector(x, y, z float64) Tuple {
	return new(x, y, z, 0.0)
}

func new(x, y, z, w float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// IsPoint returns if the Tuple is a point.
func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

// IsVector returns if the Tuple is a vector.
func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

// IsVector returns if the Tuple is a vector.
func (t Tuple) IsEqual(o Tuple) bool {
	return isEqual(t.X, o.X) && isEqual(t.Y, o.Y) && isEqual(t.Z, o.Z) && isEqual(t.W, o.W)
}

// Add sums two Tuple values and returns a new Tuple.
// It can return an error if both Tuples are points.
func (t Tuple) Add(o Tuple) (Tuple, error) {
	var r Tuple

	if t.IsPoint() && o.IsPoint() {
		return r, ErrAddTwoPoints
	}

	r.X = t.X + o.X
	r.Y = t.Y + o.Y
	r.Z = t.Z + o.Z
	r.W = t.W + o.W

	return r, nil
}

// Sub subtracts two Tuple values and returns a new Tuple.
// It can return an error if both Tuples are vectors.
func (t Tuple) Sub(o Tuple) (Tuple, error) {
	var r Tuple

	if t.IsVector() && o.IsPoint() {
		return r, ErrSubPointFromVector
	}

	r.X = t.X - o.X
	r.Y = t.Y - o.Y
	r.Z = t.Z - o.Z
	r.W = t.W - o.W

	return r, nil
}

// Neg is the opposite of a tuple. It subtracts the vector from (0, 0, 0, 0).
func (t Tuple) Neg() Tuple {
	return Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}

// Mul is the multiplication of a tuple by a scalar.
func (t Tuple) Mul(s float64) Tuple {
	return Tuple{
		X: t.X * s,
		Y: t.Y * s,
		Z: t.Z * s,
		W: t.W * s,
	}
}

// Div is the division of a tuple by a scalar.
func (t Tuple) Div(s float64) (Tuple, error) {
	var r Tuple

	if s == 0.0 {
		return r, ErrDivByZero
	}

	r.X = t.X / s
	r.Y = t.Y / s
	r.Z = t.Z / s
	r.W = t.W / s

	return r, nil
}

// DotProduct is the multiplication of a vector by other vector and returns a scalar.
func (t Tuple) DotProduct(o Tuple) (float64, error) {
	var s float64

	if !t.IsVector() || !o.IsVector() {
		return s, ErrNotVector
	}

	s = t.X*o.X + t.Y*o.Y + t.Z*o.Z + t.W*o.W

	return s, nil
}

// CrossProduct is the multiplication of a vector by other vector and returns a vector.
func (t Tuple) CrossProduct(o Tuple) (Tuple, error) {
	var c Tuple

	if !t.IsVector() || !o.IsVector() {
		return c, ErrNotVector
	}

	c.X = t.Y*o.Z - t.Z*o.Y
	c.Y = t.Z*o.X - t.X*o.Z
	c.Z = t.X*o.Y - t.Y*o.X

	return c, nil
}

// Magnitude is the distance represented by a vector.
func (t Tuple) Magnitude() (float64, error) {
	var m float64

	if !t.IsVector() {
		return m, ErrNotVector
	}

	m = math.Sqrt(math.Pow(t.X, 2.0) + math.Pow(t.Y, 2.0) + math.Pow(t.Z, 2.0) + math.Pow(t.W, 2.0))

	return m, nil
}

// Normalize is the normalization of a vector.
func (t Tuple) Normalize() (Tuple, error) {
	var n Tuple

	if !t.IsVector() {
		return n, ErrNotVector
	}

	mag, err := t.Magnitude()
	if err != nil {
		return n, err
	}

	n.X = t.X / mag
	n.Y = t.Y / mag
	n.Z = t.Z / mag
	n.W = t.W / mag

	return n, nil
}
