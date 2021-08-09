package state

import "testing"

func TestEnemyHasDefaultLocation(t *testing.T) {
  enemy := new(Enemy)

  gotx := enemy.location.x
  goty := enemy.location.y

  if (gotx != 0.0) {
    t.Errorf("Initialized enemy did not have default x coord 0.0")
  }
  if (goty != 0.0) {
    t.Errorf("Initialized enemy did not have default y coord 0.0")
  }
}

func TestEnemyHasDefaultVelocity(t *testing.T) {
  enemy := new(Enemy)

  gotx := enemy.velocity.x
  goty := enemy.velocity.y

  if (gotx != 0.0) {
    t.Errorf("Initialized enemy did not have default x velocity 0.0")
  }
  if (goty != 0.0) {
    t.Errorf("Initialized enemy did not have default y velocity 0.0")
  }
}
