package point

import (
	"strconv"
	"math/rand"
)

const W = 200
const H = 80
const FX = 20
const FY = 20
const N = 2000

// MovingPoint is a Point that has position, speed, accelleration and friction assigned to it.
type MovingPoint struct {
	p Point
	s Point
	a Point
	f float64
	in bool
}

func (m *MovingPoint) String() string {
	return "p: " + m.p.String() + " | s: " + m.s.String() + " | a: " + m.a.String() + " | f: " + strconv.FormatFloat(m.f, 'f', 2, 64)
}

// Next updates the position and speed of the MovingPoint based on its speed and accelleration.
func (m *MovingPoint) Next() {
	m.p.Add(&m.s)
	m.s.Mul(m.f)
	m.s.Add(&m.a)
	m.a.X = 0
	m.a.Y = 0
}

// SetAccelleration sets the accelleration of a MovingPoint based on the center of the Field that it is in. 
func (m *MovingPoint) SetAccelleration(g [][]Field) {
	y := int(m.p.Y)/(H/FY)
	x := int(m.p.X)/(W/FX)
	if x < len(g[0]) && x >= 0 && y < len(g) && y >= 0 {
		m.in = true
		f := g[y][x]
		a := Point{f.C.X, f.C.Y}
		a.Sub(&m.p)
		a.Div(1000)
		a.Mul(10*(f.A - float64(f.N)))
		m.a = a
	} else if m.in {
		m.in = false
		m.s.X *= -1
		m.s.Y *= -1
	}
}

// Point is a 2d vector, with float values as its coordinates.
type Point struct {
	X, Y float64
}

// Add adds 2 Points.
func (p *Point) Add(q *Point) {
	p.X += q.X
	p.Y += q.Y
}

// Sub adds 2 Points.
func (p *Point) Sub(q *Point) {
	p.X -= q.X
	p.Y -= q.Y
}

// Mul scales a point.
func (p *Point) Mul(f float64) {
	p.X *= f
	p.Y *= f
}

// Div scales a point.
func (p *Point) Div(f float64) {
	p.X /= f
	p.Y /= f
}

// String converts a point into a string.
func (p *Point) String() string {
	return "(" + strconv.FormatFloat(p.X, 'f', 2, 64) + " , " + strconv.FormatFloat(p.Y, 'f', 2, 64) + ")"
}

// Print sets a grid that will be pritned.
func (m *MovingPoint) Print(g *[H][W]int) {
	y := int(m.p.Y)
	x := int(m.p.X)
	if x < W && x >= 0 && y < H && y >= 0 {
		g[y][x] = 1
	}
}

// Field contains the center of a grid of a screen, and the quantity of points contained.
type Field struct {
	C Point
	N int
	A float64
}

// FieldGrid initializes a grid of fields.
func FieldGrid(lx int, ly int, w int, h int) [][]Field {
	var grid = make([][]Field, ly)
	for y := range grid {
		grid[y] = make([]Field, lx)
		for x := range grid[y] {
			cx := (float64(w)/(2 * float64(lx))) * (float64(x)*2 + 1)
			cy := (float64(h)/(2 * float64(ly))) * (float64(y)*2 + 1)
			grid[y][x] = Field {Point{cx, cy}, 0, 0}
		}
	}
	return grid
}

// ResetField resets the quantity of points contained.
func ResetField(f [][]Field) {
	for y := range f {
		for x := range f[y] {
			f[y][x].N = 0
		}
	}
}

// SetBallsNumber sets the N of every Field.
func SetBallsNumber(ms []MovingPoint, f [][]Field) {
	for i := range ms {
		x := int(ms[i].p.X)/(W/FX)
		y := int(ms[i].p.Y)/(H/FY)
		if x < FX && x >= 0 && y < FY && y >= 0 {
			f[y][x].N++
		}
	}
}

// SetAverage sets the average of every Field.
func SetAverage(f [][]Field, n int) {
	a := float64(n)/float64((len(f)*len(f[0])))
	for y := range f {
		for x := range f[y] {
			f[y][x].A = a
		}
	}
}

// RandomMovingPoints initializes n amount of MovingPoint randomly.
func RandomMovingPoints(n int, lx int, ly int) []MovingPoint {
	l := make([]MovingPoint, n)
	for i := range l {
		m := &l[i]
		m.p = Point{rand.Float64() * float64(lx), rand.Float64() * float64(ly)}
		//TODO
		m.s = Point{3,4}//Point{(rand.Float64() * 5) - 2.5, (rand.Float64() * 5) - 2.5}
		m.a = Point{0,0}
		m.f = 0.99
		m.in = true
	}
	return l
}