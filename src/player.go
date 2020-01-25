package src

import "fmt"

type Player interface {
	GetBricks() int
	GetWeapons() int
	GetCrystals() int

	DiffCrystals(diff int)
	DiffBricks(diff int)
	DiffWeapons(diff int)

	DiffMagicians(diff int)
	DiffBuilders(diff int)
	DiffSoldiers(diff int)

	DiffCastle(diff int)
	DiffFence(diff int)

	DiffAllStats()

	Attack(diff int)

	GetCards() []Card
	ReplaceCard(pos int, newCard Card)

	Render() string
}

type player struct {
	cards []Card

	crystals, bricks, weapons     int
	magicians, builders, soldiers int
	castle, fence                 int
}

func NewPlayer(cardsPackage CardsPackage) Player {
	player := &player{}
	player.crystals = 2
	player.bricks = 2
	player.weapons = 2

	player.magicians = 2
	player.builders = 2
	player.soldiers = 2

	player.castle = 30
	player.fence = 10

	for i := 0; i < 10; i++ {
		player.cards = append(player.cards, cardsPackage.GetNextCard())
	}

	return player
}

func (p *player) DiffCrystals(diff int) {
	p.crystals += diff
}

func (p *player) DiffBricks(diff int) {
	p.bricks += diff
}

func (p *player) DiffWeapons(diff int) {
	p.weapons += diff
}

func (p *player) DiffMagicians(diff int) {
	p.magicians += diff
}

func (p *player) DiffBuilders(diff int) {
	p.builders += diff
}

func (p *player) DiffSoldiers(diff int) {
	p.soldiers += diff
}

func (p *player) DiffCastle(diff int) {
	p.castle += diff
}

func (p *player) DiffFence(diff int) {
	p.fence += diff
}

func (p *player) Attack(diff int) {
	p.fence -= diff

	if p.fence < 0 {
		p.castle += p.fence
		p.fence = 0
	}
}

func (p *player) GetCards() []Card {
	return p.cards
}

func (p *player) ReplaceCard(pos int, newCard Card) {
	p.cards[pos] = newCard
}

func (p *player) Render() string {
	return fmt.Sprintf(" Builders: (+%d, %d)\n Soldiers: (+%d, %d)\n Magicians: (+%d, %d)\n Castle: %d\n Fence: %d",
		p.builders, p.bricks,
		p.soldiers, p.weapons,
		p.magicians, p.crystals,
		p.castle, p.fence,
	)
}

func (p *player) DiffAllStats() {
	p.DiffCrystals(p.magicians)
	p.DiffBricks(p.builders)
	p.DiffWeapons(p.soldiers)
}

func (p *player) GetBricks() int   { return p.bricks }
func (p *player) GetWeapons() int  { return p.weapons }
func (p *player) GetCrystals() int { return p.crystals }
