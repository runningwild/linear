package linear_test

import (
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
	"github.com/runningwild/linear"
	"math"
)

type gospecEval func(interface{}, interface{}) (bool, gospec.Message, gospec.Message, error)

func VecExpect(c gospec.Context, a linear.Vec2, eval gospec.Matcher, b linear.Vec2) {
	c.Expect(a.X, eval, b.X)
	c.Expect(a.Y, eval, b.Y)
}

func BasicOperationsSpec(c gospec.Context) {
	a := linear.Vec2{3, 4}
	b := linear.Vec2{5, 6}
	c.Specify("Make sure adding vectors works.", func() {
		VecExpect(c, a.Add(b), Equals, linear.Vec2{8, 10})
	})
	c.Specify("Make sure subtracting vectors works.", func() {
		VecExpect(c, a.Sub(b), Equals, linear.Vec2{-2, -2})
	})
	c.Specify("Make sure dotting vectors works.", func() {
		c.Expect(a.Dot(b), IsWithin(1e-9), 39.0)
	})
	c.Specify("Make sure crossing vectors works.", func() {
		VecExpect(c, a.Cross(), Equals, linear.Vec2{-4, 3})
	})
	c.Specify("Make sure taking the magnitude of vectors works.", func() {
		c.Expect(a.Mag(), IsWithin(1e-9), 5.0)
		c.Expect(a.Mag2(), IsWithin(1e-9), 25.0)
	})
	c.Specify("Make sure scaling vectors works.", func() {
		VecExpect(c, a.Scale(3), Equals, linear.Vec2{9, 12})
	})
}

func BasicPropertiesSpec(c gospec.Context) {
	a := linear.Vec2{3, 4}
	b := linear.Vec2{5, 6}
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
		c.Expect(a.Scale(3.5).Mag(), IsWithin(1e-9), a.Mag()*3.5)
	})
	c.Specify("Check that a-(a-b) == b.", func() {
		VecExpect(c, a.Sub(a.Sub(b)), IsWithin(1e-9), b)
	})
}

func ComplexOperationsSpec(c gospec.Context) {
	c.Specify("Vec2.DistToLine() works.", func() {
		centers := []linear.Vec2{
			linear.Vec2{10, 12},
			linear.Vec2{1, -9},
			linear.Vec2{-100, -42},
			linear.Vec2{0, 1232},
		}
		radiuses := []float64{3, 10.232, 435, 1}
		thetas := []float64{0.001, 0.1, 1, 1.01, 0.034241, 0.789, 90, 179, 180}
		angles := []float64{1.01, 1.0, 1.11111, 930142}
		for _, center := range centers {
			for _, radius := range radiuses {
				for _, angle := range angles {
					for _, theta := range thetas {
						a := linear.Vec2{math.Cos(angle), math.Sin(angle)}
						b := linear.Vec2{math.Cos(angle + theta), math.Sin(angle + theta)}
						seg := linear.Seg2{a.Scale(radius).Add(center), b.Scale(radius).Add(center)}
						dist := center.DistToLine(seg)
						real_dist := radius * math.Cos(theta/2)
						if real_dist < 0 {
							real_dist = -real_dist
						}
						c.Expect(dist, IsWithin(1e-9), real_dist)
					}
				}
			}
		}
	})
}

func isLess(a_, b_ interface{}) (match bool, pos Message, neg Message, err error) {
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
		linear.Vec2{-2, -1},
		linear.Vec2{-2, -2},
		linear.Vec2{-1, -2},
		linear.Vec2{0, -2},
		linear.Vec2{1, -2},
		linear.Vec2{2, -2},
		linear.Vec2{2, -1},
		linear.Vec2{2, 0},
		linear.Vec2{2, 1},
		linear.Vec2{2, 2},
		linear.Vec2{1, 2},
		linear.Vec2{0, 2},
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
	s1 := linear.MakeSeg2(0, 0, 9, 9)
	s2 := linear.MakeSeg2(0, 2, 7, 2)
	s3 := linear.MakeSeg2(10, 0, -10, -1)
	s4 := linear.MakeSeg2(-1, 0, 1, -1)
	s5 := linear.MakeSeg2(1000, 1000, 999, 1001)
	s6 := linear.MakeSeg2(1, 0, 1, 3)
	s7 := linear.MakeSeg2(1, 1, 2, 1)
	c.Specify("Check that intersections are computed correctly.", func() {
		i12 := s1.Isect(s2)
		VecExpect(c, i12, IsWithin(1e-9), linear.Vec2{2, 2})
		i34 := s3.Isect(s4)
		VecExpect(c, i34, IsWithin(1e-9), linear.Vec2{0, -0.5})
	})
	c.Specify("Check DoesIsect().", func() {
		c.Expect(s6.DoesIsect(s7), Equals, false)
		c.Expect(s7.DoesIsect(s6), Equals, false)
	})
	c.Specify("Check DistFromOrigin().", func() {
		c.Expect(s1.DistFromOrigin(), IsWithin(1e-9), 0.0)
		c.Expect(s2.DistFromOrigin(), IsWithin(1e-9), 2.0)
		c.Expect(s3.DistFromOrigin(), IsWithin(1e-9), 0.499376169)
		c.Expect(s4.DistFromOrigin(), IsWithin(1e-9), 0.447213595)
	})
	c.Specify("Check Left() and Right().", func() {
		c.Expect(s1.Left(linear.Vec2{0, 1}), Equals, true)
		c.Expect(s1.Right(linear.Vec2{0, 1}), Equals, false)
		c.Expect(s1.Left(linear.Vec2{1000, 10000}), Equals, true)
		c.Expect(s1.Right(linear.Vec2{1000, 10000}), Equals, false)
		c.Expect(s1.Left(linear.Vec2{0, -0.001}), Equals, false)
		c.Expect(s1.Right(linear.Vec2{0, -0.001}), Equals, true)

		c.Expect(s4.Left(linear.Vec2{0, 0}), Equals, true)
		c.Expect(s4.Right(linear.Vec2{0, 0}), Equals, false)
		c.Expect(s4.Left(linear.Vec2{0, -1000000}), Equals, false)
		c.Expect(s4.Right(linear.Vec2{0, -1000000}), Equals, true)

		c.Expect(s5.Left(linear.Vec2{999.5, 1000.5001}), Equals, false)
		c.Expect(s5.Right(linear.Vec2{999.5, 1000.5001}), Equals, true)
		c.Expect(s5.Left(linear.Vec2{999.5, 1000.4999}), Equals, true)
		c.Expect(s5.Right(linear.Vec2{999.5, 1000.4999}), Equals, false)
	})
}

func SegmentsSpec2(c gospec.Context) {
	s1 := linear.MakeSeg2(0, 0, 9, 9)
	s2 := linear.MakeSeg2(0, 5, 10, 5)
	s3 := linear.MakeSeg2(-10, 10, -20, 10)
	s4 := linear.MakeSeg2(-15, 10000, -15, -10000)
	s5 := linear.MakeSeg2(0, 1, 1, 0)
	s6 := linear.MakeSeg2(0.4, 0.4, 0.6, 0.6)
	s7 := linear.MakeSeg2(0.5, 0.5, 0.6, 0.6)
	s8 := linear.MakeSeg2(0.4, 0.4, 0.5, 0.5)
	c.Specify("Check function that determines whether or not segments intersect.", func() {
		c.Expect(s1.DoesIsect(s2), Equals, true)
		c.Expect(s2.DoesIsect(s1), Equals, true)
		c.Expect(s3.DoesIsect(s4), Equals, true)
		c.Expect(s4.DoesIsect(s3), Equals, true)
		c.Expect(s1.DoesIsect(s3), Equals, false)
		c.Expect(s1.DoesIsect(s4), Equals, false)
		c.Expect(s2.DoesIsect(s3), Equals, false)
		c.Expect(s2.DoesIsect(s4), Equals, false)
		c.Expect(s3.DoesIsect(s1), Equals, false)
		c.Expect(s3.DoesIsect(s2), Equals, false)
		c.Expect(s4.DoesIsect(s1), Equals, false)
		c.Expect(s4.DoesIsect(s2), Equals, false)
		c.Expect(s5.DoesIsect(s6), Equals, true)
		c.Expect(s5.DoesIsect(s7), Equals, false)
		c.Expect(s5.DoesIsect(s8), Equals, false)
		c.Expect(s6.DoesIsect(s5), Equals, true)
		c.Expect(s7.DoesIsect(s5), Equals, false)
		c.Expect(s8.DoesIsect(s5), Equals, false)
	})
	t0 := linear.MakeSeg2(0, 0, 5, 5)
	t1 := linear.MakeSeg2(0, 10, 10, 0)
	t2 := linear.MakeSeg2(0, 0, 4.99, 4.99)
	c.Specify("Check function that determines whether or not segments intersect or touch.", func() {
		// Should return true for everything that DoesIsect returns true for.
		c.Expect(s1.DoesIsectOrTouch(s2), Equals, true)
		c.Expect(s2.DoesIsectOrTouch(s1), Equals, true)
		c.Expect(s3.DoesIsectOrTouch(s4), Equals, true)
		c.Expect(s4.DoesIsectOrTouch(s3), Equals, true)
		c.Expect(s5.DoesIsectOrTouch(s6), Equals, true)
		c.Expect(s6.DoesIsectOrTouch(s5), Equals, true)

		c.Expect(t0.DoesIsectOrTouch(t1), Equals, true)
		c.Expect(t2.DoesIsectOrTouch(t1), Equals, false)
	})
}

func PolySpec1(c gospec.Context) {
	p := linear.Poly{
		{0, 0},
		{-1, 2},
		{0, 1},
		{1, 2},
	}
	c.Specify("Check that exterior and interior segments of a polygon are correctly identified.", func() {
		visible_exterior := []linear.Seg2{
			linear.MakeSeg2(-1, 2, 0, 1),
			linear.MakeSeg2(0, 1, 1, 2),
		}
		visible_interior := []linear.Seg2{
			linear.MakeSeg2(0, 1, 1, 2),
			linear.MakeSeg2(1, 2, 0, 0),
			linear.MakeSeg2(0, 0, -1, 2),
		}
		c.Expect(p.VisibleExterior(linear.Vec2{0, 2}), ContainsExactly, visible_exterior)
		c.Expect(p.VisibleInterior(linear.Vec2{0.5, 1.4}), ContainsExactly, visible_interior)
	})
}

func PolySpec2(c gospec.Context) {
	p := linear.Poly{
		{-1, 0},
		{-3, 0},
		{0, 10},
		{3, 0},
		{1, 0},
		{2, 1},
		{-2, 1},
	}
	c.Specify("Check that exterior and interior segments of a polygon are correctly identified.", func() {
		visible_exterior := []linear.Seg2{
			linear.MakeSeg2(-1, 0, -3, 0),
			linear.MakeSeg2(2, 1, -2, 1),
			linear.MakeSeg2(3, 0, 1, 0),
		}
		visible_interior := []linear.Seg2{
			linear.MakeSeg2(2, 1, -2, 1),
			linear.MakeSeg2(-3, 0, 0, 10),
			linear.MakeSeg2(0, 10, 3, 0),
			linear.MakeSeg2(-1, 0, -3, 0),
			linear.MakeSeg2(3, 0, 1, 0),
		}
		c.Expect(p.VisibleExterior(linear.Vec2{0, -10}), ContainsExactly, visible_exterior)
		c.Expect(p.VisibleInterior(linear.Vec2{0, 5}), ContainsExactly, visible_interior)
	})
}

func PolySpec3(c gospec.Context) {
	p := linear.Poly{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 0},
	}
	c.Specify("Check that Poly.Seg(i) returns the i-th segment.", func() {
		s0 := linear.MakeSeg2(0, 0, 0, 1)
		VecExpect(c, p.Seg(0).P, Equals, s0.P)
		VecExpect(c, p.Seg(0).Q, Equals, s0.Q)
		s1 := linear.MakeSeg2(0, 1, 1, 1)
		VecExpect(c, p.Seg(1).P, Equals, s1.P)
		VecExpect(c, p.Seg(1).Q, Equals, s1.Q)
		s2 := linear.MakeSeg2(1, 1, 1, 0)
		VecExpect(c, p.Seg(2).P, Equals, s2.P)
		VecExpect(c, p.Seg(2).Q, Equals, s2.Q)
		s3 := linear.MakeSeg2(1, 0, 0, 0)
		VecExpect(c, p.Seg(3).P, Equals, s3.P)
		VecExpect(c, p.Seg(3).Q, Equals, s3.Q)
	})
}

func TriangleSpec(c gospec.Context) {
	c.Specify("Check that areaOfPGram() works properly.", func() {
		v0 := linear.Vec2{1, 0}
		v1 := linear.Vec2{0, 0}
		v2 := linear.Vec2{0, 1}
		c.Expect(linear.AreaOfPGram(v0, v1, v2), IsWithin(1e-9), 1.0)
		v3 := linear.Vec2{-1000, 1}
		c.Expect(linear.AreaOfPGram(v0, v1, v3), IsWithin(1e-9), 1.0)
		v4 := linear.Vec2{-10, 20}
		v5 := linear.Vec2{10, 10}
		v6 := linear.Vec2{-5, 10}
		c.Expect(linear.AreaOfPGram(v4, v5, v6), IsWithin(1e-9), 150.0)
	})
}

func PolyOverlapSpec(c gospec.Context) {
	c.Specify("Check that ConvexPolysOverlap() works properly.", func() {
		// p0 contains p1, p2 and p3
		// p1 is disjoint from everything else
		// p2 and p3 share a segment
		// p3 and p4 share a vertex
		// p4 and p5 intersect
		p0 := linear.Poly{
			linear.Vec2{0, 0},
			linear.Vec2{0, 10},
			linear.Vec2{10, 10},
			linear.Vec2{10, 0},
		}
		p1 := linear.Poly{
			linear.Vec2{1, 7},
			linear.Vec2{1, 9},
			linear.Vec2{3, 9},
			linear.Vec2{3, 7},
		}
		p2 := linear.Poly{
			linear.Vec2{1, 3},
			linear.Vec2{1, 5},
			linear.Vec2{3, 5},
			linear.Vec2{3, 3},
		}
		p3 := linear.Poly{
			linear.Vec2{3, 3},
			linear.Vec2{3, 5},
			linear.Vec2{5, 5},
			linear.Vec2{5, 3},
		}
		p4 := linear.Poly{
			linear.Vec2{5, 5},
			linear.Vec2{5, 6},
			linear.Vec2{6, 6},
		}
		p5 := linear.Poly{
			linear.Vec2{5, 1},
			linear.Vec2{6, 10},
			linear.Vec2{7, 2},
		}
		c.Expect(linear.ConvexPolysOverlap(p0, p1), Equals, true)
		c.Expect(linear.ConvexPolysOverlap(p0, p2), Equals, true)
		c.Expect(linear.ConvexPolysOverlap(p0, p3), Equals, true)
		c.Expect(linear.ConvexPolysOverlap(p0, p4), Equals, true)
		c.Expect(linear.ConvexPolysOverlap(p0, p5), Equals, true)
		c.Expect(linear.ConvexPolysOverlap(p1, p2), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p1, p3), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p1, p4), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p1, p5), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p2, p3), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p2, p4), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p2, p5), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p3, p4), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p3, p5), Equals, false)
		c.Expect(linear.ConvexPolysOverlap(p4, p5), Equals, true)
	})
}
