package state

import "testing"

func TestPlayerHasDefaultLocation(t *testing.T) {
  player := new(Player)

  gotx := player.location.x
  goty := player.location.y

  if (gotx != 0.0) {
    t.Errorf("Initialized player did not have default x coord 0.0")
  }
  if (goty != 0.0) {
    t.Errorf("Initialized player did not have default y coord 0.0")
  }
}

func TestPlayerHasDefaultVelocity(t *testing.T) {
  player := new(Player)

  gotx := player.velocity.x
  goty := player.velocity.y

  if (gotx != 0.0) {
    t.Errorf("Initialized player did not have default x velocity 0.0")
  }
  if (goty != 0.0) {
    t.Errorf("Initialized player did not have default y velocity 0.0")
  }
}
