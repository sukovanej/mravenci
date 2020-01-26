package src

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Game interface {
	Start()
	Render() string
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

func clearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func validateConsumption(player Player, consumption CardConsumption) bool {
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

func validateMove(player Player, card Card) bool {
	consumption := card.GetConsumption()
	return validateConsumption(player, consumption)
}

func consumeMaterial(player Player, consumption CardConsumption) {
	switch consumption.Material {
	case Bricks:
		player.DiffBricks(-consumption.Amount)
	case Weapons:
		player.DiffWeapons(-consumption.Amount)
	case Crystals:
		player.DiffCrystals(-consumption.Amount)
	}
	return
}

type CardAction int

const (
	CardPlay CardAction = iota
	CardDiscard
)

func (game *game) cli() (CardAction, Card, int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Only homos type in here: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		if text == "help" {
			fmt.Printf(" - type the number of the card you desire to play\n - type `discard <number>` or `x <number>` to discard the card\n\n")
			continue
		} else {
			numberMatched, _ := regexp.MatchString(`^[0-9]$`, text)
			if numberMatched {
				i, _ := strconv.Atoi(text)
				card := game.currentPlayer.GetCards()[i]

				if !validateMove(game.currentPlayer, card) {
					fmt.Println("You can't do that move :(")
					continue
				}

				return CardPlay, card, i
			}

			dicardNumberMatched, _ := regexp.MatchString(`^(discard|x) [0-9]$`, text)
			if dicardNumberMatched {
				i, _ := strconv.Atoi(strings.Fields(text)[1])
				return CardDiscard, nil, i
			}
		}

		fmt.Println("I don't understand that, type help for help!")
	}

	panic("This shouldn't never happen. Contact Milan because he is a f*cking morron and made this shit wrong!")
}

func (game *game) Start() {
	for {
		clearTerminal()
		game.currentPlayer.DiffAllStats()
		fmt.Println(game.Render())

		action, card, index := game.cli()

		if action == CardPlay {
			card.Play(game.currentPlayer, game.opponentPlayer)
			consumeMaterial(game.currentPlayer, card.GetConsumption())
		} else if action == CardDiscard {
			// dicarding the card
		}

		game.currentPlayer.ReplaceCard(index, game.cardsPackage.GetNextCard())
		game.currentPlayer, game.opponentPlayer = game.opponentPlayer, game.currentPlayer
	}
}

func (game *game) Render() string {
	board := fmt.Sprintf("Red player: \n%s\n\n---------\n\nBlack player:\n%s\n", game.redPlayer.Render(), game.blackPlayer.Render())

	if game.currentPlayer == game.blackPlayer {
		board += "Black player's move:"
	} else {
		board += "Red player's move:"
	}

	board += "\n\n----------\n\nCards:\n"
	for i, card := range game.currentPlayer.GetCards() {
		consumption := card.GetConsumption()

		formatStart := "\033[0m"
		formatEnd := "\033[0m"

		if !validateConsumption(game.currentPlayer, consumption) {
			formatStart = "\033[2m"
		}

		board += fmt.Sprintf("%s(%d) %s [%d %s]%s\n",
			formatStart, i, card.Render(), consumption.Amount, RenderMaterial(consumption.Material), formatEnd)
	}

	return board
}
