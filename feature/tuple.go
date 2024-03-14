package feature

import "errors"

var (
	ErrAddTwoPoints = errors.New("can't add two points")
	ErrSubPointFromVector = errors.New("can't sub point from vector")
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