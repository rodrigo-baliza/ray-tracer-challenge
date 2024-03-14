package feature

import (
	"errors"
	"fmt"
	"strings"
)

const (
	IdentifierP3 = "P3"
	MaxColor     = 255
)

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
	for i := range width * height {
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

// Size returns the Canvas size.
func (c *Canvas) Fill(color Tuple) {
	for i := range c.pixels {
		c.pixels[i] = color
	}
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

// ToPPM returns an PPM version of the canvas
func (c *Canvas) ToPPM(identifier string, maxColor int) string {
	const maxSpace = 70

	ppm := strings.Builder{}

	// Header
	ppm.WriteString(fmt.Sprintf("%s\n", identifier))
	ppm.WriteString(fmt.Sprintf("%d %d\n", c.width, c.height))
	ppm.WriteString(fmt.Sprintf("%d\n", maxColor))

	// Data
	line := ""
	for _, pixel := range c.pixels {
		r, g, b := pixel.ColorString(maxColor)

		if line == "" {
			line = r
		} else if len(line)+len(r)+1 > maxSpace {
			ppm.WriteString(fmt.Sprintf("%s\n", line))
			line = r
		} else {
			line = fmt.Sprintf("%s %s", line, r)
		}

		if line == "" {
			line = g
		} else if len(line)+len(g)+1 > maxSpace {
			ppm.WriteString(fmt.Sprintf("%s\n", line))
			line = g
		} else {
			line = fmt.Sprintf("%s %s", line, g)
		}

		if line == "" {
			line = b
		} else if len(line)+len(b)+1 > maxSpace {
			ppm.WriteString(fmt.Sprintf("%s\n", line))
			line = b
		} else {
			line = fmt.Sprintf("%s %s", line, b)
		}
	}
	if line != "" {
		ppm.WriteString(line)
	}

	// Ending new line
	ppm.WriteString("\n")

	return ppm.String()
}

func (c *Canvas) xy2pos(x, y int) (int, error) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return 0, ErrInvalidCanvasPoint
	}

	pos := x*c.width + y

	return pos, nil
}
