package linear

import "math"

type Vec2 struct {
	X, Y float64
}

func MakeVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

// Just for the sake of testing
func (u Vec2) Equals(other interface{}) bool {
	v := other.(Vec2)
	return u.X == v.X && u.Y == v.Y
}

func (u Vec2) Add(v Vec2) Vec2 {
	return Vec2{u.X + v.X, u.Y + v.Y}
}
func (u Vec2) Sub(v Vec2) Vec2 {
	return Vec2{u.X - v.X, u.Y - v.Y}
}
func (u Vec2) Dot(v Vec2) float64 {
	return u.X*v.X + u.Y*v.Y
}

// u.Cross() gives a vector that is perpendicular to u
// u.Cross() of the zero vector gives the zero vector
func (u Vec2) Cross() Vec2 {
	return Vec2{-u.Y, u.X}
}

// Returns a vector with the same angle as u that has a magnitude of 1
func (u Vec2) Norm() Vec2 {
	mag := u.Mag()
	return Vec2{u.X / mag, u.Y / mag}
}
func (u Vec2) Mag() float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y)
}

// Returns the squared magnitude of the vector, faster than calling u.Mag()
func (u Vec2) Mag2() float64 {
	return u.X*u.X + u.Y*u.Y
}
func (u Vec2) Scale(scale float64) Vec2 {
	return Vec2{u.X * scale, u.Y * scale}
}
func (u Vec2) Angle() float64 {
	return math.Atan2(u.Y, u.X)
}
func (u Vec2) DistToLine(s Seg2) float64 {
	q := s.Ray().Cross()
	t := Seg2{u, q.Add(u)}
	sect := s.Isect(t)
	return sect.Sub(u).Mag()
}
func (u Vec2) Rotate(angle float64) Vec2 {
	return Vec2{
		X: u.X*math.Cos(angle) - u.Y*math.Sin(angle),
		Y: u.X*math.Sin(angle) + u.Y*math.Cos(angle),
	}
}
func (u Vec2) RotateAround(v Vec2, angle float64) Vec2 {
	return u.Sub(v).Rotate(angle).Add(v)
}

type Seg2 struct {
	P, Q Vec2
}

func MakeSeg2(x, y, x2, y2 float64) Seg2 {
	return Seg2{Vec2{x, y}, Vec2{x2, y2}}
}

// Just for the sake of testing
func (a Seg2) Equals(other interface{}) bool {
	b := other.(Seg2)
	return a.P.Equals(b.P) && a.Q.Equals(b.Q)
}

func (a Seg2) Ray() Vec2 {
	return a.Q.Sub(a.P)
}

// Returns a Vec2 indicating the intersection point of the lines passing
// through segments a and b
func (a Seg2) Isect(b Seg2) Vec2 {
	by := b.Q.Y - b.P.Y
	bx := b.P.X - b.Q.X
	n := (b.P.X-a.P.X)*by + (b.P.Y-a.P.Y)*bx
	d := (a.Q.X-a.P.X)*by + (a.Q.Y-a.P.Y)*bx
	f := n / d
	return Vec2{a.P.X + (a.Q.X-a.P.X)*f, a.P.Y + (a.Q.Y-a.P.Y)*f}
}

// Returns a value V indicating where along a that b isects it.
// V == 0: Intersection at a[0]
// V == 1: Intersection at a[1]
// V in (0, 1): Intersection between a[0] and a[1]
// otherwise: No intersection
func (a Seg2) RelIsect(b Seg2) float64 {
	by := b.Q.Y - b.P.Y
	bx := b.P.X - b.Q.X
	n := (b.P.X-a.P.X)*by + (b.P.Y-a.P.Y)*bx
	d := (a.Q.X-a.P.X)*by + (a.Q.Y-a.P.Y)*bx
	return n / d
}

func (a Seg2) DoesIsect(b Seg2) bool {
	return (b.Left(a.P) && b.Right(a.Q) || b.Left(a.Q) && b.Right(a.P)) &&
		(a.Left(b.P) && a.Right(b.Q) || a.Left(b.Q) && a.Right(b.P))
}

func (a Seg2) DistFromOrigin() float64 {
	r := Seg2{a.Ray().Cross(), Vec2{0, 0}}.Isect(a)
	return r.Mag()
}

// Returns true iff u lies to the left of a
func (a Seg2) Left(u Vec2) bool {
	return a.Ray().Cross().Dot(u.Sub(a.P)) > 0
}

// Returns true iff u lies to the right of a
func (a Seg2) Right(u Vec2) bool {
	return a.Ray().Cross().Dot(u.Sub(a.P)) < 0
}

// The vertices of a polygon should be in clockwise order
type Poly []Vec2

func (p Poly) Seg(i int) Seg2 {
	return Seg2{p[i], p[(i+1)%len(p)]}
}

func (p Poly) visibility(u Vec2, f func(Seg2, Vec2) bool) []Seg2 {
	segs := make([]Seg2, len(p))[0:0]
	for i := 1; i < len(p); i++ {
		s := Seg2{p[i-1], p[i]}
		if f(s, u) {
			segs = append(segs, s)
		}
	}
	s := Seg2{p[len(p)-1], p[0]}
	if f(s, u) {
		segs = append(segs, s)
	}
	return segs
}

// Returns the set of line segments of p that might be visible from u,
// assuming that u does not lie within p.
func (p Poly) VisibleExterior(u Vec2) []Seg2 {
	return p.visibility(u, func(s Seg2, v Vec2) bool { return s.Left(v) })
}

// Returns the set of line segments of p that might be visible from u,
// assuming that u lies within p.
func (p Poly) VisibleInterior(u Vec2) []Seg2 {
	return p.visibility(u, func(s Seg2, v Vec2) bool { return s.Right(v) })
}

func AreaOfPGram(v0, v1, v2 Vec2) float64 {
	return (v2.X-v1.X)*(v1.Y-v0.Y) - (v2.Y-v1.Y)*(v1.X-v0.X)
}

func (p Poly) Area() float64 {
	var sum float64
	for i := range p {
		j := (i + 1) % len(p)
		sum += p[i].X*p[j].Y - p[i].Y*p[j].X
	}
	return sum / -2
}

func convexPolysDisjointOneWay(a, b Poly) bool {
	for i := range a {
		disjoint := true
		for j := range b {
			if a.Seg(i).Right(b[j]) {
				disjoint = false
				break
			}
		}
		if disjoint {
			return true
		}
	}
	return false
}

func ConvexPolysOverlap(a, b Poly) bool {
	return !convexPolysDisjointOneWay(a, b) && !convexPolysDisjointOneWay(b, a)
}

func VecInsideConvexPoly(vec Vec2, poly Poly) bool {
	for i := range poly {
		seg := poly.Seg(i)
		if seg.Left(vec) {
			return false
		}
	}
	return true
}
