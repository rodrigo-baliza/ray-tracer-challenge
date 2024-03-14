package feature_test

import (
	"errors"
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
			if diff := pretty.Diff(got, test.want); len(diff) != 0 {
				t.Errorf("\n%s", strings.Join(diff, "\n"))
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
			if diff := pretty.Diff(got, test.want); len(diff) != 0 {
				t.Errorf("\n%s", strings.Join(diff, "\n"))
			}
		})
	}
}
