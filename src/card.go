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
		&SwatCard{},
		&FortCard{},
		&WainCard{},
		&BabylonCard{},
		&ConjureBricksCard{},
		&SorcererCard{},
		&RecruitingCard{},
		&PixiesCard{},
		&ThiefCard{},
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

func (_ *BaseCard) Play(me Player, opponent Player) { me.SetCastle(me.GetCastle() + 2) }
func (_ *BaseCard) Render() string                  { return "Base - Castle +2" }
func (_ *BaseCard) GetConsumption() CardConsumption { return CardConsumption{1, Bricks} }

// Fence Card

type FenceCard struct{}

func (_ *FenceCard) Play(me Player, opponent Player) { me.SetFence(me.GetFence() + 22) }
func (_ *FenceCard) Render() string                  { return "Fence - Fence +20" }
func (_ *FenceCard) GetConsumption() CardConsumption { return CardConsumption{12, Bricks} }

// Fence Card

type SchoolCard struct{}

func (_ *SchoolCard) Play(me Player, opponent Player) { me.SetBuilders(me.GetBuilders() + 1) }
func (_ *SchoolCard) Render() string                  { return "School - Builders +1" }
func (_ *SchoolCard) GetConsumption() CardConsumption { return CardConsumption{8, Bricks} }

// Swat Card

type SwatCard struct{}

func (_ *SwatCard) Play(me Player, opponent Player) { opponent.SetCastle(opponent.GetCastle() - 10) }
func (_ *SwatCard) Render() string                  { return "SWAT - Castle of your enemy -10" }
func (_ *SwatCard) GetConsumption() CardConsumption { return CardConsumption{18, Weapons} }

// Fort Card

type FortCard struct{}

func (_ *FortCard) Play(me Player, opponent Player) { me.SetCastle(me.GetCastle() + 20) }
func (_ *FortCard) Render() string                  { return "Fort - Castle of your enemy -10" }
func (_ *FortCard) GetConsumption() CardConsumption { return CardConsumption{18, Bricks} }

// Wain Card

type WainCard struct{}

func (_ *WainCard) Play(me Player, opponent Player) {
	me.SetCastle(me.GetCastle() + 8)
	opponent.SetCastle(opponent.GetCastle() - 4)
}
func (_ *WainCard) Render() string                  { return "Wain - Castle +8, Castle of your enemy -4" }
func (_ *WainCard) GetConsumption() CardConsumption { return CardConsumption{10, Bricks} }

// Babylon Card

type BabylonCard struct{}

func (_ *BabylonCard) Play(me Player, opponent Player) { me.SetCastle(me.GetCastle() + 32) }
func (_ *BabylonCard) Render() string                  { return "Babylon - Castle +32" }
func (_ *BabylonCard) GetConsumption() CardConsumption { return CardConsumption{39, Bricks} }

// Conjure bricks Card

type ConjureBricksCard struct{}

func (_ *ConjureBricksCard) Play(me Player, opponent Player) { me.SetBricks(me.GetBricks() + 8) }
func (_ *ConjureBricksCard) Render() string                  { return "Conjure bricks - Bricks +8" }
func (_ *ConjureBricksCard) GetConsumption() CardConsumption { return CardConsumption{4, Crystals} }

// Sorcerer Card

type SorcererCard struct{}

func (_ *SorcererCard) Play(me Player, opponent Player) { me.SetMages(me.GetMages() + 1) }
func (_ *SorcererCard) Render() string                  { return "Sorcerer - Soldiers +1" }
func (_ *SorcererCard) GetConsumption() CardConsumption { return CardConsumption{8, Crystals} }

// Recruiting

type RecruitingCard struct{}

func (_ *RecruitingCard) Play(me Player, opponent Player) { me.DiffSoldiers(1)  }
func (_ *RecruitingCard) Render() string                  { return "Recruiting - Soldiers +1"}
func (_ *RecruitingCard) GetConsumption() CardConsumption { return CardConsumption{8, Weapons}}

// Pixies

type PixiesCard struct{}

func (_ *PixiesCard) Play(me Player, opponent Player) { me.DiffCastle(22)}
func (_ *PixiesCard) Render() string                  { return "Pixies - Castle +22"  }
func (_ *PixiesCard) GetConsumption() CardConsumption { return CardConsumption{22, Crystals}}

// Thief

type ThiefCard struct{}

func transfereMaterial(
	meDiff func(int), 
	opponentDiff func(int),
	opponentGet func() int,
) {
	opponentDiff(-5)
	meDiff(5)

	if b := opponentGet(); b < 0 {
		meDiff(b)
		opponentDiff(-b)
	}
}

func (_ *ThiefCard) Play(me Player, opponent Player) {
	transfereMaterial(me.DiffBricks, opponent.DiffBricks, opponent.GetBricks)
	transfereMaterial(me.DiffWeapons, opponent.DiffWeapons, opponent.GetWeapons)
	transfereMaterial(me.DiffCrystals, opponent.DiffCrystals, opponent.GetCrystals)
}
func (_ *ThiefCard) Render() string                  { return "Thief - Transfer stocks of your enemy 5"}
func (_ *ThiefCard) GetConsumption() CardConsumption { return CardConsumption{15, Weapons}}


