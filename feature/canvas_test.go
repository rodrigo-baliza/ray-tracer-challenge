package feature_test

import (
	"errors"
	"ray-tracer/feature"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		err    error
	}{
		{
			name:   "valid",
			width:  1,
			height: 1,
			err:    nil,
		},
		{
			name:   "invalid 1",
			width:  0,
			height: 1,
			err:    feature.ErrInvalidCanvasSize,
		},
		{
			name:   "invalid 2",
			width:  -1,
			height: 1,
			err:    feature.ErrInvalidCanvasSize,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := feature.NewCanvas(test.width, test.height)

			if !errors.Is(err, test.err) {
				t.Errorf("%q: got error %v, expected error %v", test.name, err, test.err)
			}
			if err != nil {
				return
			}

			if got.Size() != test.width*test.height {
				t.Errorf("%q: got a canvas with size %d, expected a canvas with size %d", test.name, got.Size(), test.width*test.height)
			}
			for i := range test.width - 1 {
				for j := range test.height - 1 {
					pixel, err := got.Pixel(i, j)
					if err != nil {
						t.Errorf("%q: expected no error iterating width and height but got %v", test.name, err)
					}

					if !pixel.IsEqual(feature.ColorBlack) {
						t.Errorf("%q: expected pixel with black color iterating width and height but got %+v", test.name, pixel)
					}
				}
			}

			if _, err := got.Pixel(test.width, test.height); err == nil {
				t.Errorf("%q: expected error getting a pixel out of the canvas bounds but got no error", test.name)
			}
		})
	}
}

func TestWritePixel(t *testing.T) {
	tests := []struct {
		name  string
		x     int
		y     int
		pixel feature.Tuple
		want  feature.Tuple
		err   error
	}{
		{
			name:  "valid 1",
			x:     1,
			y:     1,
			pixel: feature.ColorRed,
			want:  feature.ColorRed,
			err:   nil,
		},
		{
			name:  "valid 2",
			x:     2,
			y:     1,
			pixel: feature.ColorBlue,
			want:  feature.ColorBlue,
			err:   nil,
		},
		{
			name:  "invalid 1",
			x:     -1,
			y:     3,
			pixel: feature.ColorRed,
			want:  feature.Tuple{},
			err:   feature.ErrInvalidCanvasPoint,
		},
		{
			name:  "invalid 2",
			x:     3,
			y:     3,
			pixel: feature.ColorRed,
			want:  feature.Tuple{},
			err:   feature.ErrInvalidCanvasPoint,
		},
		{
			name:  "invalid 3",
			x:     9,
			y:     3,
			pixel: feature.ColorRed,
			want:  feature.Tuple{},
			err:   feature.ErrInvalidCanvasPoint,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			canvas, err := feature.NewCanvas(3, 3)
			if err != nil {
				t.Fatalf("error creating a new canvas: %v", err)
			}

			err = canvas.WritePixel(test.x, test.y, test.pixel)
			if !errors.Is(err, test.err) {
				t.Errorf("%q: WritePixel got error %v, expected error %v", test.name, err, test.err)
			}

			got, err := canvas.Pixel(test.x, test.y)
			if !errors.Is(test.err, err) {
				t.Errorf("%q: Pixel got error %v, expected error %v", test.name, err, test.err)
			}

			if !got.IsEqual(test.want) {
				t.Errorf("%q: expected pixel %v from the canvas but got %+v", test.name, test.want, got)
			}
		})
	}
}
