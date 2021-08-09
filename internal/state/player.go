package state

type Player struct {
    location Location
    velocity Velocity
    geometry Geometry
    score int
    alive bool

    // If the player has been killed, this
    // variable should contain the time (in game time)
    // at which they wil reappear on the field
    respawnTime float32
  }
