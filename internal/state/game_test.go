package state

import "testing"

func TestANewGameShouldHaveAnEmptyPlayerList(t *testing.T){
  game := new(Game)
  if !(len(game.players)==0) {
    t.Errorf("New game did not have a zero-length list of players")
  }
}

func TestANewGameShouldHaveAnEmptyEnemyList(t *testing.T){
  game := new(Game)
  if !(len(game.enemies)==0) {
    t.Errorf("New game did not have a zero-length list of enemies")
  }
}

func TestANewGameShouldHaveAnEmptyGoalsList(t *testing.T){
  game := new(Game)
  if !(len(game.goals)==0) {
    t.Errorf("New game did not have a zero-length list of goals")
  }
}
