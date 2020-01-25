package src

import "math/rand"

type Card interface {
	Render() string
	Play(me Player, opponent Player)
	GetConsumption() CardConsumption
}

type Material int

const (
	Bricks Material = iota
	Weapons
	Crystals
)

func RenderMaterial(material Material) string {
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

type CardConsumption struct {
	Amount   int
	Material Material
}

type CardsPackage interface {
	GetNextCard() Card
}

type randomCardsPackage struct {
	possibleCards []Card
}

func NewCardsPackage() CardsPackage {
	cardsPackage := &randomCardsPackage{}
	cardsPackage.possibleCards = append(cardsPackage.possibleCards,
		&BansheeCard{},
		&PlatoonCard{},
		&AttackCard{},
		&BaseCard{},
		&FenceCard{},
		&SchoolCard{},
	)
	return cardsPackage
}

func (p *randomCardsPackage) GetNextCard() Card {
	return p.possibleCards[rand.Intn(len(p.possibleCards))]
}

// Banshee card

type BansheeCard struct{}

func (_ *BansheeCard) Play(me Player, opponent Player) { opponent.Attack(32) }
func (_ *BansheeCard) Render() string                  { return "Banshee - attack +32" }
func (_ *BansheeCard) GetConsumption() CardConsumption { return CardConsumption{28, Weapons} }

// Platoon Card

type PlatoonCard struct{}

func (_ *PlatoonCard) Play(me Player, opponent Player) { opponent.Attack(6) }
func (_ *PlatoonCard) Render() string                  { return "Platoon - attack +6" }
func (_ *PlatoonCard) GetConsumption() CardConsumption { return CardConsumption{4, Weapons} }

// Attack Card

type AttackCard struct{}

func (_ *AttackCard) Play(me Player, opponent Player) { opponent.Attack(12) }
func (_ *AttackCard) Render() string                  { return "Attack - attack +12" }
func (_ *AttackCard) GetConsumption() CardConsumption { return CardConsumption{10, Weapons} }

// Base Card

type BaseCard struct{}

func (_ *BaseCard) Play(me Player, opponent Player) { me.DiffCastle(2) }
func (_ *BaseCard) Render() string                  { return "Base - Castle +2" }
func (_ *BaseCard) GetConsumption() CardConsumption { return CardConsumption{1, Bricks} }

// Fence Card

type FenceCard struct{}

func (_ *FenceCard) Play(me Player, opponent Player) { me.DiffFence(22) }
func (_ *FenceCard) Render() string                  { return "Fence - Fence +20" }
func (_ *FenceCard) GetConsumption() CardConsumption { return CardConsumption{12, Bricks} }

// Fence Card

type SchoolCard struct{}

func (_ *SchoolCard) Play(me Player, opponent Player) { me.DiffBuilders(1) }
func (_ *SchoolCard) Render() string                  { return "School - Builders +1" }
func (_ *SchoolCard) GetConsumption() CardConsumption { return CardConsumption{8, Bricks} }
