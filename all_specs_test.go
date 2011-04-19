package linear_test

import (
  "gospec"
  "testing"
)


func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  r.AddSpec(BasicOperationsSpec)
  r.AddSpec(BasicPropertiesSpec)
  r.AddSpec(AnglesSpec)
  r.AddSpec(SegmentsSpec)
  r.AddSpec(PolySpec1)
  r.AddSpec(PolySpec2)
  gospec.MainGoTest(r, t)
}

