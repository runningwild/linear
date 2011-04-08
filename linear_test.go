package linear_test

import (
  . "gospec"
  "gospec"
  "linear"
  "os"
)

type gospecEval func(interface {}, interface {}) (bool, gospec.Message, gospec.Message, os.Error)

func VecExpect(c gospec.Context, a linear.Vec2, eval gospec.Matcher, b linear.Vec2) {
  c.Expect(a[0], eval, b[0])
  c.Expect(a[1], eval, b[1])
}

func BasicOperationsSpec(c gospec.Context) {
  a := linear.Vec2{3,4}
  b := linear.Vec2{5,6}
  c.Specify("Make sure adding vectors works.", func() {
    VecExpect(c, a.Add(b), Equals, linear.Vec2{8,10})
  })
  c.Specify("Make sure subtracting vectors works.", func() {
    VecExpect(c, a.Sub(b), Equals, linear.Vec2{-2,-2})
  })
  c.Specify("Make sure dotting vectors works.", func() {
    c.Expect(a.Dot(b), IsWithin(1e-9), 39.0)
  })
  c.Specify("Make sure crossing vectors works.", func() {
    VecExpect(c, a.Cross(), Equals, linear.Vec2{-4,3})
  })
  c.Specify("Make sure taking the magnitude of vectors works.", func() {
    c.Expect(a.Mag(), IsWithin(1e-9), 5.0)
    c.Expect(a.Mag2(), IsWithin(1e-9), 25.0)
  })
  c.Specify("Make sure scaling vectors works.", func() {
    VecExpect(c, a.Scale(3), Equals, linear.Vec2{9,12})
  })
}

func BasicPropertiesSpec(c gospec.Context) {
  a := linear.Vec2{3,4}
  b := linear.Vec2{5,6}
  c.Specify("Check that (cross a) dot a == 0.", func() {
    c.Expect(a.Cross().Dot(a), Equals, 0.0)
  })
  c.Specify("Check that a normalize vector's magnitude is 1.", func() {
    c.Expect(a.Norm().Mag(), IsWithin(1e-9), 1.0)
  })
  c.Specify("Check that v.Mag2() == v.Mag()*v.Mag()", func() {
    c.Expect(a.Mag2(), IsWithin(1e-9), a.Mag()*a.Mag())
  })
  c.Specify("Check that a scaled vector's magnitude is appropriately scaled.", func() {
    c.Expect(a.Scale(3.5).Mag(), IsWithin(1e-9), a.Mag() * 3.5)
  })
  c.Specify("Check that a-(a-b) == b.", func() {
    VecExpect(c, a.Sub(a.Sub(b)), IsWithin(1e-9), b)
  })
}

func isLess(a_,b_ interface{}) (match bool, pos Message, neg Message, err os.Error) {
  a := a_.(float64)
  b := b_.(float64)
  match = a < b
  pos = Messagef(a, "should be less than %v", b)
  neg = Messagef(a, "should NOT be less than %v", b)
  err = nil
  return
}

func AnglesSpec(c gospec.Context) {
  v := []linear.Vec2{
    linear.Vec2{-2,-1},
    linear.Vec2{-2,-2},
    linear.Vec2{-1,-2},
    linear.Vec2{ 0,-2},
    linear.Vec2{ 1,-2},
    linear.Vec2{ 2,-2},
    linear.Vec2{ 2,-1},
    linear.Vec2{ 2, 0},
    linear.Vec2{ 2, 1},
    linear.Vec2{ 2, 2},
    linear.Vec2{ 1, 2},
    linear.Vec2{ 0, 2},
    linear.Vec2{-1, 2},
    linear.Vec2{-2, 2},
    linear.Vec2{-2, 1},
  }
  c.Specify("Check that Vec2.Angle() is properly ordered.", func() {
    for i := 1; i < len(v); i++ {
      c.Expect(v[i-1].Angle(), isLess, v[i].Angle())
    }
  })
}

