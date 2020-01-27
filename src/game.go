package src

type Game interface {
	IsValidToPlayNthCard(card_index int) bool

	PlayNthCard(card_index int)
	DiscardNthCard(card_index int)

	StartRound()
	EndRound()

	GetCurrentPlayer() Player
	GetOpponentPlayer() Player
	GetBlackPlayer() Player
	GetRedPlayer() Player
}

type game struct {
	blackPlayer, redPlayer        Player
	currentPlayer, opponentPlayer Player
	cardsPackage                  CardsPackage
}

func NewGame() Game {
	game := &game{}
	game.cardsPackage = NewCardsPackage()
	game.blackPlayer = NewPlayer(game.cardsPackage)
	game.redPlayer = NewPlayer(game.cardsPackage)
	game.currentPlayer = game.blackPlayer
	game.opponentPlayer = game.redPlayer
	return game
}

func consumeMaterial(player Player, consumption CardConsumption) {
	switch consumption.Material {
	case Bricks:
		player.SetBricks(player.GetBricks() - consumption.Amount)
	case Weapons:
		player.SetWeapons(player.GetWeapons() - consumption.Amount)
	case Crystals:
		player.SetCrystals(player.GetCrystals() - consumption.Amount)
	}
	return
}

func (game *game) IsValidToPlayNthCard(card_index int) bool {
	player := game.currentPlayer
	card := player.GetCards()[card_index]
	consumption := card.GetConsumption()

	if consumption.Material == Bricks && player.GetBricks() < consumption.Amount {
		return false
	}

	if consumption.Material == Weapons && player.GetWeapons() < consumption.Amount {
		return false
	}

	if consumption.Material == Crystals && player.GetCrystals() < consumption.Amount {
		return false
	}

	return true
}

func (game *game) PlayNthCard(card_index int) {
	card := game.currentPlayer.GetCards()[card_index]
	card.Play(game.currentPlayer, game.opponentPlayer)
	consumeMaterial(game.currentPlayer, card.GetConsumption())
	game.currentPlayer.ReplaceCard(card_index, game.cardsPackage.GetNextCard())
}

func (game *game) DiscardNthCard(card_index int) {
	game.currentPlayer.ReplaceCard(card_index, game.cardsPackage.GetNextCard())
}

func (game *game) StartRound() {
	me := game.currentPlayer
	me.SetBricks(me.GetBricks() + me.GetBuilders())
	me.SetWeapons(me.GetWeapons() + me.GetSoldiers())
	me.SetCrystals(me.GetCrystals() + me.GetMages())
}

func (game *game) EndRound() {
	game.currentPlayer, game.opponentPlayer = game.opponentPlayer, game.currentPlayer
}

func (game *game) GetCurrentPlayer() Player  { return game.currentPlayer }
func (game *game) GetOpponentPlayer() Player { return game.opponentPlayer }
func (game *game) GetBlackPlayer() Player    { return game.blackPlayer }
func (game *game) GetRedPlayer() Player      { return game.redPlayer }
