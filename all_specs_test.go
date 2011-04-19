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
  gospec.MainGoTest(r, t)
}

