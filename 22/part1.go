package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2020/util"
)

var (
	inputFile = "input"
)

type Player struct {
	cards []int
}

func NewPlayer() Player {
	p := Player{}
	p.cards = make([]int, 0)
	return p
}

func (p Player) String() string {
	out := ""
	for _, card := range p.cards {
		out += fmt.Sprintf("%d ", card)
	}
	return out
}

func (p *Player) PutFirstCardToBack() {
	if len(p.cards) < 2 {
		return
	}
	p.cards = append(p.cards[1:], p.cards[0])
}

func (p *Player) RemoveFirstCard() int {
	if len(p.cards) == 0 {
		return -1
	}

	card := p.cards[0]
	p.cards = p.cards[1:]

	return card
}

func main() {

	// re := regexp.MustCompile(`^$`)

	player1 := NewPlayer()
	player2 := NewPlayer()
	currentPlayer := &player1
	for i, line := range util.LoadString(inputFile) {

		if strings.HasPrefix(line, "Player") {
			continue
		}

		if line == "" {
			currentPlayer = &player2
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error converting line %d: %s\n", i, line)
			return
		}

		currentPlayer.cards = append(currentPlayer.cards, num)
	}

	fmt.Println("P1", player1.cards)
	fmt.Println("P2", player2.cards)

	for len(player1.cards) > 0 && len(player2.cards) > 0 {
		if player1.cards[0] > player2.cards[0] {
			player1.PutFirstCardToBack()
			player1.cards = append(player1.cards, player2.RemoveFirstCard())
		} else if player2.cards[0] > player1.cards[0] {
			player2.PutFirstCardToBack()
			player2.cards = append(player2.cards, player1.RemoveFirstCard())
		} else {
			fmt.Println("PAT", player1.cards[0], player2.cards[0])
		}
	}

	result := 0
	if len(player1.cards) > 0 {
		fmt.Println("Player 1 wins")
		currentPlayer = &player1
	} else {
		fmt.Println("Player 2 wins")
		currentPlayer = &player2
	}

	fmt.Println("P1", player1.cards)
	fmt.Println("P2", player2.cards)
	for i, card := range currentPlayer.cards {
		result += card * (len(currentPlayer.cards) - i)
	}
	fmt.Println(result)
}
