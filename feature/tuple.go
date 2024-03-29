package feature

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrAddTwoPoints       = errors.New("can't add two points")
	ErrSubPointFromVector = errors.New("can't sub point from vector")
	ErrDivByZero          = errors.New("can't div by zero")
	ErrNotVector          = errors.New("it's not a vector")

	ColorBlack = NewColor(0, 0, 0)
	ColorWhite = NewColor(1, 1, 1)
	ColorRed   = NewColor(1, 0, 0)
	ColorGreen = NewColor(0, 1, 0)
	ColorBlue  = NewColor(0, 0, 1)
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
	return newTuple(x, y, z, 1.0)
}

// NewVector creates a new vector tuple.
func NewVector(x, y, z float64) Tuple {
	return newTuple(x, y, z, 0.0)
}

// NewColor creates a new color tuple.
func NewColor(r, g, b float64) Tuple {
	return newTuple(r, g, b, 0.0)
}

func newTuple(x, y, z, w float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// String is the string representation of the Tuple.
func (t Tuple) String() string {
	return fmt.Sprintf("proj new position: x[%f] y[%f] z[%f] w [%f]", t.X, t.Y, t.Z, t.W)
}

// ColorString is the string representation of the Color.
func (t Tuple) ColorString(max int) (string, string, string) {
	x := clamp(t.X, max)
	y := clamp(t.Y, max)
	z := clamp(t.Z, max)

	return fmt.Sprintf("%d", x), fmt.Sprintf("%d", y), fmt.Sprintf("%d", z)
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

// HadamardProduct is the multiplication of a color by other color and returns a color.
func (t Tuple) HadamardProduct(o Tuple) Tuple {
	return newTuple(
		t.X*o.X,
		t.Y*o.Y,
		t.Z*o.Z,
		0.0,
	)
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

func isEqual(a, b float64) bool {
	const EPSILON = 0.00001

	return math.Abs(a-b) < EPSILON
}

func clamp(value float64, max int) int {
	c := value * float64(max)
	c = math.Round(c)

	if c > 255 {
		c = 255
	} else if c < 0 {
		c = 0
	}

	return int(c)
}
