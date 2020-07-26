package point

import (
	_ "fmt"
	"image"
	"strconv"
)

// MovingPoint is a Point that has a Point and accelleration assigned to it.
type MovingPoint struct {
	p Point
	s Point
	a Point
}

func (m *MovingPoint) next() {
	m.p.Add(&m.s)
	m.s.Add(&m.a)
}

type Field struct {
	image.Rectangle
	c Point
}

func Fld(r image.Rectangle) Field {
	c := Point{float64(r.Min.X+r.Max.X) / 2, float64(r.Min.Y+r.Max.Y) / 2}
	return Field{r, c}
}

type Point struct {
	x, y float64
}

func (p1 *Point) Add(p2 *Point) {
	p1.x += p2.x
	p1.y += p2.y
}

func (p *Point) String() string {
	return "(" + strconv.FormagFloat(p.x, 'e', 2, 64) + " , " + strconv.FormagFloat(p.y, 'e', 2, 64) + ")"
}
