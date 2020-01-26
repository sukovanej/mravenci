package src

import "fmt"

type View interface {
	RenderGame() string
}

type view struct {
	game Game
}

func NewView(game Game) View {
	return &view{game}
}

func (view *view) renderMaterial(material Material) string {
	switch material {
	case Bricks:
		return "Bricks"
	case Weapons:
		return "Weapons"
	case Crystals:
		return "Crystals"
	}

	panic("Unknown material")
}

func (view *view) renderPlayer(player Player) string {
	return fmt.Sprintf(" Builders: (+%d, %d)\n Soldiers: (+%d, %d)\n Mages: (+%d, %d)\n Castle: %d\n Fence: %d",
		player.GetBuilders(), player.GetBricks(),
		player.GetSoldiers(), player.GetWeapons(),
		player.GetMages(), player.GetCrystals(),
		player.GetCastle(), player.GetFence(),
	)
}

func (view *view) RenderGame() string {
	currentPlayer := view.game.GetCurrentPlayer()
	redPlayer, blackPlayer := view.game.GetRedPlayer(), view.game.GetBlackPlayer()

	board := fmt.Sprintf("Red player: \n%s\n\n---------\n\nBlack player:\n%s\n\n",
		view.renderPlayer(redPlayer),
		view.renderPlayer(blackPlayer))

	if currentPlayer == blackPlayer {
		board += "Black player's move:"
	} else {
		board += "Red player's move:"
	}

	board += "\n\n----------\n\nCards:\n"
	for i, card := range currentPlayer.GetCards() {
		consumption := card.GetConsumption()

		formatStart := "\033[0m"
		formatEnd := "\033[0m"

		if !view.game.IsValidToPlayNthCard(i) {
			formatStart = "\033[2m"
		}

		board += fmt.Sprintf("%s(%d) %s [%d %s]%s\n",
			formatStart, i, card.Render(), consumption.Amount, view.renderMaterial(consumption.Material), formatEnd)
	}

	return board
}
