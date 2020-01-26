package tests

import "testing"
import . "github.com/sukovanej/mravenci/src"

func TestSorcererCard(t *testing.T) {
	cards := NewCardsPackage()
	me, opponent := NewPlayer(cards), NewPlayer(cards)

	card := SorcererCard{}
	card.Play(me, opponent)

	if me.GetMages() != 3 {
		t.Errorf("Unexpected result after Sorcerer playe.")
	}
}
