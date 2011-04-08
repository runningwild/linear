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
func (u Vec2) Cross() Vec2 {
  return Vec2{ -u[1], u[0] }
}
func (u Vec2) Norm() Vec2 {
  mag := u.Mag()
  return Vec2{ u[0] / mag, u[1] / mag }
}
func (u Vec2) Mag() float64 {
  return math.Sqrt(u[0]*u[0] + u[1]*u[1])
}
func (u Vec2) Mag2() float64 {
  return u[0]*u[0] + u[1]*u[1]
}
func (u Vec2) Scale(scale float64) Vec2 {
  return Vec2{ u[0] * scale, u[1] * scale }
}
func (u Vec2) Angle() float64 {
  return math.Atan2(u[1], u[0])
}

