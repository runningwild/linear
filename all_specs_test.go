package linear_test

import (
	"github.com/orfjackal/gospec/src/gospec"
	"testing"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(BasicOperationsSpec)
	r.AddSpec(ComplexOperationsSpec)
	r.AddSpec(BasicPropertiesSpec)
	r.AddSpec(AnglesSpec)
	r.AddSpec(SegmentsSpec)
	r.AddSpec(SegmentsSpec2)
	r.AddSpec(PolySpec1)
	r.AddSpec(PolySpec2)
	r.AddSpec(TriangleSpec)
	r.AddSpec(PolyOverlapSpec)
	gospec.MainGoTest(r, t)
}
