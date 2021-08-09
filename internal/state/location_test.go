package state

import "testing"

func TestSetAndRetrieveLocation(t *testing.T){
  loc := new(Location)
  loc.x = 1.0
  loc.y = -1.5

  if (loc.x != 1.0) {
    t.Errorf("X not retrieved correctly got %v, wanted %v ", loc.x, 1.0)
  }
  if (loc.y != -1.5) {
    t.Errorf("Y not retrieved correctly got %v, wanted %v ", loc.x, 1.0)
  }
}
