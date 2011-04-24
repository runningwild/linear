package linear

import "math"

type Vec2 [2]float64

// Just for the sake of testing
func (u Vec2) Equals(other interface{}) bool {
  v := other.(Vec2)
  return u[0] == v[0] && u[1] == v[1]
}

func (u Vec2) Add(v Vec2) Vec2 {
  return Vec2{ u[0]+v[0], u[1]+v[1] }
}
func (u Vec2) Sub(v Vec2) Vec2 {
  return Vec2{ u[0]-v[0], u[1]-v[1] }
}
func (u Vec2) Dot(v Vec2) float64 {
  return u[0]*v[0] + u[1]*v[1]
}

// u.Cross() gives a vector that is perpendicular to u
// u.Cross() of the zero vector gives the zero vector
func (u Vec2) Cross() Vec2 {
  return Vec2{ -u[1], u[0] }
}

// Returns a vector with the same angle as u that has a magnitude of 1
func (u Vec2) Norm() Vec2 {
  mag := u.Mag()
  return Vec2{ u[0] / mag, u[1] / mag }
}
func (u Vec2) Mag() float64 {
  return math.Sqrt(u[0]*u[0] + u[1]*u[1])
}

// Returns the squared magnitude of the vector, faster than calling u.Mag()
func (u Vec2) Mag2() float64 {
  return u[0]*u[0] + u[1]*u[1]
}
func (u Vec2) Scale(scale float64) Vec2 {
  return Vec2{ u[0] * scale, u[1] * scale }
}
func (u Vec2) Angle() float64 {
  return math.Atan2(u[1], u[0])
}

type Seg2 [2]Vec2

// Just for the sake of testing
func (a Seg2) Equals(other interface{}) bool {
  b := other.(Seg2)
  return a[0].Equals(b[0]) && a[1].Equals(b[1])
}

func (a Seg2) Ray() Vec2 {
  return a[1].Sub(a[0])
}

// Returns a bool indicating whether or not the two line segments intersect
// Returns a Vec2 indicating the intersection point if they intersect
func (a Seg2) Isect(b Seg2) (Vec2,bool) {
  by := b[1][1] - b[0][1]
  bx := b[0][0] - b[1][0]
  n := (b[0][0] - a[0][0]) * by + (b[0][1] - a[0][1]) * bx
  d := (a[1][0] - a[0][0]) * by + (a[1][1] - a[0][1]) * bx
  f := n/d
  return Vec2{ a[0][0] + (a[1][0] - a[0][0]) * f, a[0][1] + (a[1][1] - a[0][1]) * f}, true
}

func (a Seg2) DistFromOrigin() float64 {
  r,_ := Seg2{a.Ray().Cross(),Vec2{0,0}}.Isect(a)
  return r.Mag()
}

// Returns true iff u lies to the left of a
func (a Seg2) Left(u Vec2) bool {
  return a.Ray().Cross().Dot(u.Sub(a[0])) > 0
}

// Returns true iff u lies to the right of a
func (a Seg2) Right(u Vec2) bool {
  return a.Ray().Cross().Dot(u.Sub(a[0])) < 0
}


// The vertices of a polygon should be in clockwise order
type Poly []Vec2

func (p Poly) visibility(u Vec2, f func(Seg2,Vec2) bool) []Seg2 {
  segs := make([]Seg2, len(p))[0:0]
  for i := 1; i < len(p); i++ {
    s := Seg2{p[i-1],p[i]}
    if f(s,u) {
      segs = append(segs, s)
    }
  }
  s := Seg2{p[len(p)-1], p[0]}
  if f(s,u) {
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

