package state

import "testing"

func TestSetAndRetrieveGeometry(t *testing.T){
  geo := new(Geometry)
  geo.radius = 3.0

  if (geo.radius != 3.0) {
    t.Errorf("Radius not retrieved correctly got %v, wanted %v ", geo.radius, 3.0)
  }
}
