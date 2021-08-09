package state

import "testing"

func TestSetAndRetrieveVelocity(t *testing.T){
  vel := new(Velocity)
  vel.x = 1.0
  vel.y = -1.5

  if (vel.x != 1.0) {
    t.Errorf("X not retrieved correctly got %v, wanted %v ", vel.x, 1.0)
  }
  if (vel.y != -1.5) {
    t.Errorf("Y not retrieved correctly got %v, wanted %v ", vel.x, 1.0)
  }
}
