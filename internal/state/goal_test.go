package state

import "testing"

func TestGoalHasDefaultLocation(t *testing.T) {
  goal := new(Goal)

  gotx := goal.location.x
  goty := goal.location.y

  if (gotx != 0.0) {
    t.Errorf("Initialized goal did not have default x coord 0.0")
  }
  if (goty != 0.0) {
    t.Errorf("Initialized goal did not have default y coord 0.0")
  }
}
