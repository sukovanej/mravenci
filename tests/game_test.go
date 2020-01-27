package tests

import "testing"
import . "github.com/sukovanej/mravenci/src"

func TestGameStart(t *testing.T) {
	game := NewGame()
	me := game.GetCurrentPlayer()

	initialBricks, initialWeapons, initialCrystals := me.GetBricks(), me.GetWeapons(), me.GetCrystals()

	if initialBricks != 5 || initialWeapons != 5 || initialCrystals != 5 {
		t.Errorf("Initial material should be 5 each, [%d, %d, %d] found", initialBricks, initialWeapons, initialCrystals)
	}

	game.StartRound()
	game.EndRound()

	if me != game.GetOpponentPlayer() {
		t.Errorf("After first round players must change")
	}

	if me.GetBricks() != initialBricks || me.GetWeapons() != initialWeapons || me.GetCrystals() != initialCrystals {
		t.Errorf("After first round 1st player's materials must remain the same")
	}
}

func TestGameSecondRound(t *testing.T) {
	game := NewGame()
	me, opponent := game.GetCurrentPlayer(), game.GetOpponentPlayer()

	initialBricks, initialWeapons, initialCrystals := me.GetBricks(), me.GetWeapons(), me.GetCrystals()

	game.StartRound()
	game.EndRound()

	game.StartRound()

	if opponent.GetBricks() == initialBricks || opponent.GetWeapons() == initialWeapons || opponent.GetCrystals() == initialCrystals {
		t.Errorf("When second round starts 2nd player's materials must change")
	}
}
