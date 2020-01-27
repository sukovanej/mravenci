package src

type Player interface {
	GetBricks() int
	GetWeapons() int
	GetCrystals() int

	GetBuilders() int
	GetSoldiers() int
	GetMages() int

	GetCastle() int
	GetFence() int

	SetCrystals(value int)
	SetBricks(value int)
	SetWeapons(value int)

	SetMages(value int)
	SetBuilders(value int)
	SetSoldiers(value int)

	SetCastle(value int)
	SetFence(value int)

	Attack(value int)

	GetCards() []Card
	ReplaceCard(pos int, newCard Card)
}

type player struct {
	cards []Card

	crystals, bricks, weapons int
	mages, builders, soldiers int
	castle, fence             int
}

func NewPlayer(cardsPackage CardsPackage) Player {
	player := &player{}
	player.crystals = 2
	player.bricks = 2
	player.weapons = 2

	player.mages = 2
	player.builders = 2
	player.soldiers = 2

	player.castle = 30
	player.fence = 10

	for i := 0; i < 10; i++ {
		player.cards = append(player.cards, cardsPackage.GetNextCard())
	}

	return player
}

func (p *player) SetCrystals(value int) { p.crystals = value }
func (p *player) SetBricks(value int)   { p.bricks = value }
func (p *player) SetWeapons(value int)  { p.weapons = value }

func (p *player) SetMages(value int)    { p.mages = value }
func (p *player) SetBuilders(value int) { p.builders = value }
func (p *player) SetSoldiers(value int) { p.soldiers = value }

func (p *player) SetCastle(value int) { p.castle += value }
func (p *player) SetFence(value int)  { p.fence += value }

func (p *player) Attack(value int) {
	p.fence -= value

	if p.fence < 0 {
		p.castle += p.fence
		p.fence = 0
	}
}

func (p *player) GetCards() []Card { return p.cards }

func (p *player) ReplaceCard(pos int, newCard Card) {
	p.cards[pos] = newCard
}

func (p *player) GetBricks() int   { return p.bricks }
func (p *player) GetWeapons() int  { return p.weapons }
func (p *player) GetCrystals() int { return p.crystals }

func (p *player) GetBuilders() int { return p.builders }
func (p *player) GetSoldiers() int { return p.soldiers }
func (p *player) GetMages() int    { return p.mages }

func (p *player) GetCastle() int { return p.castle }
func (p *player) GetFence() int  { return p.fence }
