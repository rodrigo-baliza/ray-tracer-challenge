package feature

import "errors"

var (
	ErrInvalidCanvasSize  = errors.New("invalid canvas size")
	ErrInvalidCanvasPoint = errors.New("invalid canvas point")
)

// Canvas is a rectangular (widht x height) grid of pixels, where each
// pixel has a color setting.
type Canvas struct {
	width  int
	height int
	pixels []Tuple
}

// NewCanvas creates a new Canvas with width and height sizes, where each
// pixel is initialized to black (0, 0, 0)
func NewCanvas(width, height int) (*Canvas, error) {
	if width <= 0 || height <= 0 {
		return nil, ErrInvalidCanvasSize
	}

	pixels := make([]Tuple, width*height)
	for i := range width*height{
		pixels[i] = ColorBlack
	}

	c := Canvas{
		width:  width,
		height: height,
		pixels: pixels,
	}

	return &c, nil
}

// Size returns the Canvas size.
func (c *Canvas) Size() int {
	return c.width * c.height
}

// Pixel returns the pixel color in the position x and y.
// It returns an error if any of the positions be invalid.
func (c *Canvas) Pixel(x, y int) (Tuple, error) {
	var p Tuple

	pos, err := c.xy2pos(x, y)
	if err != nil {
		return p, err
	}

	p = c.pixels[pos]

	return p, nil
}

// WritePixel cheanges the color of a point in the position x and y.
func (c *Canvas) WritePixel(x, y int, color Tuple) error {
	pos, err := c.xy2pos(x, y)
	if err != nil {
		return err
	}

	c.pixels[pos] = color

	return nil
}

func (c *Canvas) xy2pos(x, y int) (int, error) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return 0, ErrInvalidCanvasPoint
	}

	pos := x*c.width + y

	return pos, nil
}
