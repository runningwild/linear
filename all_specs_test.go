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
  gospec.MainGoTest(r, t)
}

