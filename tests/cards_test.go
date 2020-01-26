package tests

import "testing"
import . "github.com/sukovanej/mravenci/src"

func TestSorcererCard(t *testing.T) {
	cards := NewCardsPackage()
	me, opponent := NewPlayer(cards), NewPlayer(cards)

	card := SorcererCard{}
	card.Play(me, opponent)

	if me.GetMages() != 3 {
		t.Errorf("Unexpected result after Sorcerer played.")
	}
}

func TestThiefCard(t *testing.T) {
	cards := NewCardsPackage()
	me, opponent := NewPlayer(cards), NewPlayer(cards)

	me.SetWeapons(15)
	me.SetBricks(15)
	me.SetCrystals(15)

	opponent.SetWeapons(15)
	opponent.SetBricks(15)
	opponent.SetCrystals(15)

	card := ThiefCard{}
	card.Play(me, opponent)

	if me.GetBricks() != 20 {
		t.Errorf("Me: Expected 20, got %d.", me.GetBricks())
	}
	if me.GetWeapons() != 20 {
		t.Errorf("Me: Expected 20, got %d.", me.GetWeapons())
	}
	if me.GetCrystals() != 20 {
		t.Errorf("Me: Expected 20, got %d.", me.GetCrystals())
	}
	if opponent.GetBricks() != 10 {
		t.Errorf("Opponent: Expected 20, got %d.", opponent.GetBricks())
	}
	if opponent.GetWeapons() != 10 {
		t.Errorf("Opponent: Expected 10, got %d.", opponent.GetWeapons())
	}
	if opponent.GetCrystals() != 10 {
		t.Errorf("Opponent: Expected 10, got %d.", opponent.GetCrystals())
	}

	// Test not enough materials
	me.SetWeapons(15)
	opponent.SetWeapons(3)

	card.Play(me, opponent)

	if me.GetWeapons() != 18 {
		t.Errorf("Me: Expected 18, got %d.", me.GetWeapons())
	}

	if opponent.GetWeapons() != 0 {
		t.Errorf("Opponent: Expected 0, got %d.", opponent.GetWeapons())
	}

}
