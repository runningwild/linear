package linear

import "math"

type Vec2 [2]float64

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

func (a Seg2) Ray() Vec2 {
  return a[1].Sub(a[0])
}

// Returns a bool indicating whether or not the two line segments intersect
// Returns a Vec2 indicating the intersection point if they intersect
func (a Seg2) Isect(b Seg2) (Vec2,bool) {
  a_to_b := b[0].Sub(a[0])
  bx := b.Ray().Cross()
  aray := a.Ray()
  h := a_to_b.Dot(bx) / aray.Dot(bx)
  return a[0].Add(aray.Scale(h)), true
}

func (a Seg2) DistFromOrigin() float64 {
  r,_ := Seg2{a.Ray().Cross(),Vec2{0,0}}.Isect(a)
  return r.Mag()
}

func (a Seg2) Left(u Vec2) bool {
  return a.Ray().Cross().Dot(u.Sub(a[0])) > 0
}

func (a Seg2) Right(u Vec2) bool {
  return a.Ray().Cross().Dot(u.Sub(a[0])) < 0
}


