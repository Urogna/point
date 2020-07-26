package point

import (
	"image"
	"strconv"
)

// MovingPoint is a Point that has position, speed and accelleration assigned to it.
type MovingPoint struct {
	p Point
	s Point
	a Point
}

// next() updates the position and speed of the MovingPoint based on its speed and accelleration.
func (m *MovingPoint) next() {
	m.p.Add(&m.s)
	m.s.Add(&m.a)
}

// Field is a rectangle with the center position.
type Field struct {
	image.Rectangle
	c Point
}

// Fld initializes a Field.
func Fld(r image.Rectangle) Field {
	c := Point{float64(r.Min.X+r.Max.X) / 2, float64(r.Min.Y+r.Max.Y) / 2}
	return Field{r, c}
}

// Point is a 2d vector, with float values as its coordinates.
type Point struct {
	x, y float64
}

// Add adds 2 Points.
func (p *Point) Add(q *Point) {
	p.x += q.x
	p.y += q.y
}

// String converts a point into a string.
func (p *Point) String() string {
	return "(" + strconv.FormatFloat(p.x, 'e', 2, 64) + " , " + strconv.FormatFloat(p.y, 'e', 2, 64) + ")"
}
