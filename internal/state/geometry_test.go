package state

import "testing"

func TestSetAndRetrieveGeometry(t *testing.T){
  geo := new(Geometry)
  geo.radius = 1.0

  if (geo.radius != 1.0) {
    t.Errorf("Radius not retrieved correctly got %v, wanted %v ", geo.radius, 1.0)
  }
}
