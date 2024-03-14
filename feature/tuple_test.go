package feature_test

import (
	"errors"
	"math"
	"ray-tracer/feature"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

func TestNewPoint(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		y    float64
		z    float64
		want feature.Tuple
	}{
		{
			name: "is a point",
			x:    4.3,
			y:    -4.2,
			z:    3.1,
			want: feature.Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 1.0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := feature.NewPoint(test.x, test.y, test.z)

			if !got.IsPoint() {
				t.Error("expected IsPoint() = true, but got IsPoint() = false")
			}
			if got.IsVector() {
				t.Error("expected IsVector() = false, but got IsVector() = true")
			}
			if diff := pretty.Diff(got, test.want); len(diff) != 0 {
				t.Errorf("\n%s", strings.Join(diff, "\n"))
			}
		})
	}
}

func TestNewVector(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		y    float64
		z    float64
		want feature.Tuple
	}{
		{
			name: "is a vector",
			x:    4.3,
			y:    -4.2,
			z:    3.1,
			want: feature.Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 0.0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := feature.NewVector(test.x, test.y, test.z)

			if !got.IsVector() {
				t.Error("expected IsVector() = true, but got IsVector() = false")
			}
			if got.IsPoint() {
				t.Error("expected IsPoint() = false, but got IsPoint() = true")
			}
			if diff := pretty.Diff(got, test.want); len(diff) != 0 {
				t.Errorf("\n%s", strings.Join(diff, "\n"))
			}
		})
	}
}

func TestNewColor(t *testing.T) {
	tests := []struct {
		name string
		r    float64
		g    float64
		b    float64
		want feature.Tuple
	}{
		{
			name: "is a color",
			r:    -0.5,
			g:    0.4,
			b:    1.7,
			want: feature.Tuple{
				X: -0.5,
				Y: 0.4,
				Z: 1.7,
				W: 0.0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := feature.NewColor(test.r, test.g, test.b)

			if diff := pretty.Diff(got, test.want); len(diff) != 0 {
				t.Errorf("\n%s", strings.Join(diff, "\n"))
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		first  feature.Tuple
		second feature.Tuple
		want   feature.Tuple
		err    error
	}{
		{
			name: "add one point and one vector",
			first: feature.Tuple{
				X: 3.0,
				Y: -2.0,
				Z: 5.0,
				W: 1.0,
			},
			second: feature.Tuple{
				X: -2.0,
				Y: 3.0,
				Z: 1.0,
				W: 0.0,
			},
			want: feature.Tuple{
				X: 1.0,
				Y: 1.0,
				Z: 6.0,
				W: 1.0,
			},
			err: nil,
		},
		{
			name: "add two vectors",
			first: feature.Tuple{
				X: 3.0,
				Y: -2.0,
				Z: 5.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: -2.0,
				Y: 3.0,
				Z: 1.0,
				W: 0.0,
			},
			want: feature.Tuple{
				X: 1.0,
				Y: 1.0,
				Z: 6.0,
				W: 0.0,
			},
			err: nil,
		},
		{
			name: "add two points",
			first: feature.Tuple{
				X: 3.0,
				Y: -2.0,
				Z: 5.0,
				W: 1.0,
			},
			second: feature.Tuple{
				X: -2.0,
				Y: 3.0,
				Z: 1.0,
				W: 1.0,
			},
			want: feature.Tuple{},
			err:  feature.ErrAddTwoPoints,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.first.Add(test.second)

			if !errors.Is(err, test.err) {
				t.Errorf("%q: got error %v, expected error %v", test.name, err, test.err)
			}
			if !test.want.IsEqual(got) {
				t.Errorf("%s wants %+v and got %+v", test.name, test.want, got)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name   string
		first  feature.Tuple
		second feature.Tuple
		want   feature.Tuple
		err    error
	}{
		{
			name: "sub two points",
			first: feature.Tuple{
				X: 3.0,
				Y: 2.0,
				Z: 1.0,
				W: 1.0,
			},
			second: feature.Tuple{
				X: 5.0,
				Y: 6.0,
				Z: 7.0,
				W: 1.0,
			},
			want: feature.Tuple{
				X: -2.0,
				Y: -4.0,
				Z: -6.0,
				W: 0.0,
			},
			err: nil,
		},
		{
			name: "sub one vector from a point",
			first: feature.Tuple{
				X: 3.0,
				Y: 2.0,
				Z: 1.0,
				W: 1.0,
			},
			second: feature.Tuple{
				X: 5.0,
				Y: 6.0,
				Z: 7.0,
				W: 0.0,
			},
			want: feature.Tuple{
				X: -2.0,
				Y: -4.0,
				Z: -6.0,
				W: 1.0,
			},
			err: nil,
		},
		{
			name: "sub two vectors",
			first: feature.Tuple{
				X: 3.0,
				Y: 2.0,
				Z: 1.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: 5.0,
				Y: 6.0,
				Z: 7.0,
				W: 0.0,
			},
			want: feature.Tuple{
				X: -2.0,
				Y: -4.0,
				Z: -6.0,
				W: 0.0,
			},
			err: nil,
		},
		{
			name: "sub one point from a vector",
			first: feature.Tuple{
				X: 3.0,
				Y: 2.0,
				Z: 1.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: 5.0,
				Y: 6.0,
				Z: 7.0,
				W: 1.0,
			},
			want: feature.Tuple{},
			err:  feature.ErrSubPointFromVector,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.first.Sub(test.second)

			if !errors.Is(err, test.err) {
				t.Errorf("%q: got error %v, expected error %v", test.name, err, test.err)
			}
			if !test.want.IsEqual(got) {
				t.Errorf("%s wants %+v and got %+v", test.name, test.want, got)
			}
		})
	}
}

func TestNeg(t *testing.T) {
	tests := []struct {
		name  string
		tuple feature.Tuple
		want  feature.Tuple
	}{
		{
			name: "tuple",
			tuple: feature.Tuple{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: -4.0,
			},
			want: feature.Tuple{
				X: -1.0,
				Y: 2.0,
				Z: -3.0,
				W: 4.0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.tuple.Neg()

			if !test.want.IsEqual(got) {
				t.Errorf("%s wants %+v and got %+v", test.name, test.want, got)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name   string
		tuple  feature.Tuple
		scalar float64
		want   feature.Tuple
	}{
		{
			name: "gt one",
			tuple: feature.Tuple{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: -4.0,
			},
			scalar: 3.5,
			want: feature.Tuple{
				X: 3.5,
				Y: -7.0,
				Z: 10.5,
				W: -14.0,
			},
		},
		{
			name: "lt one",
			tuple: feature.Tuple{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: -4.0,
			},
			scalar: 0.5,
			want: feature.Tuple{
				X: 0.5,
				Y: -1.0,
				Z: 1.5,
				W: -2.0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.tuple.Mul(test.scalar)

			if !test.want.IsEqual(got) {
				t.Errorf("%s wants %+v and got %+v", test.name, test.want, got)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name   string
		tuple  feature.Tuple
		scalar float64
		want   feature.Tuple
		err    error
	}{
		{
			name: "gt one",
			tuple: feature.Tuple{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: -4.0,
			},
			scalar: 2.0,
			want: feature.Tuple{
				X: 0.5,
				Y: -1.0,
				Z: 1.5,
				W: -2.0,
			},
			err: nil,
		},
		{
			name: "lt one",
			tuple: feature.Tuple{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: -4.0,
			},
			scalar: 0.5,
			want: feature.Tuple{
				X: 2.0,
				Y: -4.0,
				Z: 6.0,
				W: -8.0,
			},
			err: nil,
		},
		{
			name: "by zero",
			tuple: feature.Tuple{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: -4.0,
			},
			scalar: 0.0,
			want:   feature.Tuple{},
			err:    feature.ErrDivByZero,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.tuple.Div(test.scalar)

			if !errors.Is(err, test.err) {
				t.Errorf("%q: got error %v, expected error %v", test.name, err, test.err)
			}
			if !test.want.IsEqual(got) {
				t.Errorf("%s wants %+v and got %+v", test.name, test.want, got)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		name  string
		tuple feature.Tuple
		want  float64
		err   error
	}{
		{
			name: "x 1",
			tuple: feature.Tuple{
				X: 1.0,
				Y: 0.0,
				Z: 0.0,
				W: 0.0,
			},
			want: 1.0,
			err:  nil,
		},
		{
			name: "y 1",
			tuple: feature.Tuple{
				X: 0.0,
				Y: 1.0,
				Z: 0.0,
				W: 0.0,
			},
			want: 1.0,
			err:  nil,
		},
		{
			name: "z 1",
			tuple: feature.Tuple{
				X: 0.0,
				Y: 0.0,
				Z: 1.0,
				W: 0.0,
			},
			want: 1.0,
			err:  nil,
		},
		{
			name: "positive vector",
			tuple: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			want: math.Sqrt(14.0),
			err:  nil,
		},
		{
			name: "negative vector",
			tuple: feature.Tuple{
				X: -1.0,
				Y: -2.0,
				Z: -3.0,
				W: 0.0,
			},
			want: math.Sqrt(14.0),
			err:  nil,
		},
		{
			name: "not a vector",
			tuple: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 1.0,
			},
			want: 0.0,
			err:  feature.ErrNotVector,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.tuple.Magnitude()

			if !errors.Is(err, test.err) {
				t.Errorf("%q: got error %v, expected error %v", test.name, err, test.err)
			}
			if test.want != got {
				t.Errorf("%s wants %f and got %f", test.name, test.want, got)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name  string
		tuple feature.Tuple
		want  feature.Tuple
		err   error
	}{
		{
			name: "case 1",
			tuple: feature.Tuple{
				X: 4.0,
				Y: 0.0,
				Z: 0.0,
				W: 0.0,
			},
			want: feature.Tuple{
				X: 1.0,
				Y: 0.0,
				Z: 0.0,
				W: 0.0,
			},
			err: nil,
		},
		{
			name: "case 2",
			tuple: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			want: feature.Tuple{
				X: 0.26726,
				Y: 0.53452,
				Z: 0.80178,
				W: 0.0,
			},
			err: nil,
		},
		{
			name: "not a vector",
			tuple: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 1.0,
			},
			want: feature.Tuple{},
			err:  feature.ErrNotVector,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.tuple.Normalize()

			if !errors.Is(err, test.err) {
				t.Errorf("%q: got error %v, expected error %v", test.name, err, test.err)
			}
			if !test.want.IsEqual(got) {
				t.Errorf("%s wants %+v and got %+v", test.name, test.want, got)
			}
		})
	}
}

func TestDotProduct(t *testing.T) {
	tests := []struct {
		name   string
		first  feature.Tuple
		second feature.Tuple
		want   float64
		err    error
	}{
		{
			name: "case 1",
			first: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: 2.0,
				Y: 3.0,
				Z: 4.0,
				W: 0.0,
			},
			want: 20.0,
			err:  nil,
		},
		{
			name: "case 2",
			first: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: -2.0,
				Y: -3.0,
				Z: -4.0,
				W: 0.0,
			},
			want: -20.0,
			err:  nil,
		},
		{
			name: "case 3",
			first: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: 2.0,
				Y: 3.0,
				Z: -4.0,
				W: 0.0,
			},
			want: -4.0,
			err:  nil,
		},
		{
			name: "invalid 1",
			first: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 1.0,
			},
			second: feature.Tuple{
				X: 2.0,
				Y: 3.0,
				Z: 4.0,
				W: 0.0,
			},
			want: 0.0,
			err:  feature.ErrNotVector,
		},
		{
			name: "invalid 2",
			first: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: 2.0,
				Y: 3.0,
				Z: 4.0,
				W: 1.0,
			},
			want: 0.0,
			err:  feature.ErrNotVector,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.first.DotProduct(test.second)

			if !errors.Is(err, test.err) {
				t.Errorf("%q: got error %v, expected error %v", test.name, err, test.err)
			}
			if test.want != got {
				t.Errorf("%s wants %f and got %f", test.name, test.want, got)
			}
		})
	}
}

func TestCrossProduct(t *testing.T) {
	tests := []struct {
		name   string
		first  feature.Tuple
		second feature.Tuple
		want   feature.Tuple
		err    error
	}{
		{
			name: "case 1",
			first: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: 2.0,
				Y: 3.0,
				Z: 4.0,
				W: 0.0,
			},
			want: feature.Tuple{
				X: -1.0,
				Y: 2.0,
				Z: -1.0,
				W: 0.0,
			},
			err: nil,
		},
		{
			name: "case 2",
			first: feature.Tuple{
				X: 2.0,
				Y: 3.0,
				Z: 4.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			want: feature.Tuple{
				X: 1.0,
				Y: -2.0,
				Z: 1.0,
				W: 0.0,
			},
			err: nil,
		},
		{
			name: "invalid 1",
			first: feature.Tuple{
				X: 2.0,
				Y: 3.0,
				Z: 4.0,
				W: 1.0,
			},
			second: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 0.0,
			},
			want: feature.Tuple{},
			err:  feature.ErrNotVector,
		},
		{
			name: "invalid 2",
			first: feature.Tuple{
				X: 2.0,
				Y: 3.0,
				Z: 4.0,
				W: 0.0,
			},
			second: feature.Tuple{
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
				W: 1.0,
			},
			want: feature.Tuple{},
			err:  feature.ErrNotVector,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.first.CrossProduct(test.second)

			if !errors.Is(err, test.err) {
				t.Errorf("%q: got error %v, expected error %v", test.name, err, test.err)
			}
			if !test.want.IsEqual(got) {
				t.Errorf("%s wants %+v and got %+v", test.name, test.want, got)
			}
		})
	}
}

func TestHadamardProduct(t *testing.T) {
	tests := []struct {
		name   string
		first  feature.Tuple
		second feature.Tuple
		want   feature.Tuple
	}{
		{
			name: "case 1",
			first: feature.Tuple{
				X: 1.0,
				Y: 0.2,
				Z: 0.4,
			},
			second: feature.Tuple{
				X: 0.9,
				Y: 1.0,
				Z: 0.1,
			},
			want: feature.Tuple{
				X: 0.9,
				Y: 0.2,
				Z: 0.04,
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.first.HadamardProduct(test.second)

			if !test.want.IsEqual(got) {
				t.Errorf("%s wants %+v and got %+v", test.name, test.want, got)
			}
		})
	}
}