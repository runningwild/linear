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

func SegmentsSpec(c gospec.Context) {
  s1 := linear.Seg2{{0,0}, {9,9}}
  s2 := linear.Seg2{{0,2}, {7,2}}
  s3 := linear.Seg2{{10,0}, {-10,-1}}
  s4 := linear.Seg2{{-1,0}, {1,-1}}
  s5 := linear.Seg2{{1000,1000}, {999,1001}}
  c.Specify("Check that intersections are computed correctly.", func() {
    i12,_ := s1.Isect(s2)
    VecExpect(c, i12, IsWithin(1e-9), linear.Vec2{2,2})
    i34,_ := s3.Isect(s4)
    VecExpect(c, i34, IsWithin(1e-9), linear.Vec2{0,-0.5})
  })
  c.Specify("Check DistFromOrigin().", func() {
    c.Expect(s1.DistFromOrigin(), IsWithin(1e-9), 0.0)
    c.Expect(s2.DistFromOrigin(), IsWithin(1e-9), 2.0)
    c.Expect(s3.DistFromOrigin(), IsWithin(1e-9), 0.499376169)
    c.Expect(s4.DistFromOrigin(), IsWithin(1e-9), 0.447213595)
  })
  c.Specify("Check Left() and Right().", func() {
    c.Expect(s1.Left(linear.Vec2{0,1}), Equals, true)
    c.Expect(s1.Right(linear.Vec2{0,1}), Equals, false)
    c.Expect(s1.Left(linear.Vec2{1000,10000}), Equals, true)
    c.Expect(s1.Right(linear.Vec2{1000,10000}), Equals, false)
    c.Expect(s1.Left(linear.Vec2{0,-0.001}), Equals, false)
    c.Expect(s1.Right(linear.Vec2{0,-0.001}), Equals, true)

    c.Expect(s4.Left(linear.Vec2{0,0}), Equals, true)
    c.Expect(s4.Right(linear.Vec2{0,0}), Equals, false)
    c.Expect(s4.Left(linear.Vec2{0,-1000000}), Equals, false)
    c.Expect(s4.Right(linear.Vec2{0,-1000000}), Equals, true)

    c.Expect(s5.Left(linear.Vec2{999.5,1000.5001}), Equals, false)
    c.Expect(s5.Right(linear.Vec2{999.5,1000.5001}), Equals, true)
    c.Expect(s5.Left(linear.Vec2{999.5,1000.4999}), Equals, true)
    c.Expect(s5.Right(linear.Vec2{999.5,1000.4999}), Equals, false)
  })
}

func PolySpec1(c gospec.Context) {
  p := linear.Poly{
    {0,0},
    {-1,2},
    {0,1},
    {1,2},
  }
  c.Specify("Check that exterior and interior segments of a polygon are correctly identified.", func() {
    visible_exterior := []linear.Seg2{ {{-1,2},{0,1}}, {{0,1},{1,2}} }
    visible_interior := []linear.Seg2{ {{0,1},{1,2}}, {{1,2},{0,0}}, {{0,0},{-1,2}} }
    c.Expect(p.VisibleExterior(linear.Vec2{0,2}), ContainsExactly, visible_exterior)
    c.Expect(p.VisibleInterior(linear.Vec2{0.5,1.4}), ContainsExactly, visible_interior)
  })
}

func PolySpec2(c gospec.Context) {
  p := linear.Poly{
    {-1,0},
    {-3,0},
    {0,10},
    {3,0},
    {1,0},
    {2,1},
    {-2,1},
  }
  c.Specify("Check that exterior and interior segments of a polygon are correctly identified.", func() {
    visible_exterior := []linear.Seg2{ {{-1,0},{-3,0}}, {{2,1},{-2,1}}, {{3,0},{1,0}} }
    visible_interior := []linear.Seg2{ {{2,1},{-2,1}}, {{-3,0},{0,10}}, {{0,10},{3,0}}, {{-1,0},{-3,0}}, {{3,0},{1,0}} }
    c.Expect(p.VisibleExterior(linear.Vec2{0,-10}), ContainsExactly, visible_exterior)
    c.Expect(p.VisibleInterior(linear.Vec2{0,5}), ContainsExactly, visible_interior)
  })
}

func PolySpec3(c gospec.Context) {
  p := linear.Poly{
    {0,0},
    {0,1},
    {1,1},
    {1,0},
  }
  c.Specify("Check that Poly.Seg(i) returns the i-th segment.", func() {
    s0 := linear.Seg2{ {0,0}, {0,1} }
    VecExpect(c, p.Seg(0)[0], Equals, s0[0])
    VecExpect(c, p.Seg(0)[1], Equals, s0[1])
    s1 := linear.Seg2{ {0,1}, {1,1} }
    VecExpect(c, p.Seg(1)[0], Equals, s1[0])
    VecExpect(c, p.Seg(1)[1], Equals, s1[1])
    s2 := linear.Seg2{ {1,1}, {1,0} }
    VecExpect(c, p.Seg(2)[0], Equals, s2[0])
    VecExpect(c, p.Seg(2)[1], Equals, s2[1])
    s3 := linear.Seg2{ {1,0}, {0,0} }
    VecExpect(c, p.Seg(3)[0], Equals, s3[0])
    VecExpect(c, p.Seg(3)[1], Equals, s3[1])
  })
}

