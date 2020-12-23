package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/RaphaelPour/aoc2020/util"
)

var (
	inputFile      = "input"
	depthHistogram = map[int]int{}
	lastUpdate     = time.Now()
	startTime      = time.Now()
)

const (
	PLAYER1 = iota
	PLAYER2
)

type Player struct {
	cards []int
}

func NewPlayer() Player {
	p := Player{}
	p.cards = make([]int, 0)
	return p
}

func (p Player) NewSubPlayer(subCardCount int) Player {
	other := Player{}
	other.cards = make([]int, subCardCount)
	copy(other.cards, p.cards[1:subCardCount+1])
	return other
}

func (p Player) String() string {
	out := ""
	for _, card := range p.cards {
		out += fmt.Sprintf("%d ", card)
	}
	return out
}

func (p Player) CanGoRecursive() bool {
	return len(p.cards)-1 >= p.cards[0]
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

type Game struct {
	p1, p2  Player
	history map[string]bool
}

func NewGame() Game {
	g := Game{}
	g.p1 = NewPlayer()
	g.p2 = NewPlayer()
	g.history = make(map[string]bool, 0)
	return g
}

func (g Game) NewSubGame(p1SubCardCount, p2SubCardCount int) Game {
	other := Game{}
	other.p1 = g.p1.NewSubPlayer(p1SubCardCount)
	other.p2 = g.p2.NewSubPlayer(p2SubCardCount)
	other.history = make(map[string]bool, 0)
	return other
}

func (g Game) IsTurnReocurring() bool {
	if len(g.history) == 0 {
		return false
	}

	_, ok := g.history[g.CurrentTurnID()]
	return ok
}

func (g Game) CurrentTurnID() string {
	return fmt.Sprintf("%v%v", g.p1.cards, g.p2.cards)
}

func (g *Game) AddTurnToHistory() {
	g.history[g.CurrentTurnID()] = true
}

func (g Game) Print(depth int, msg string) {
	return
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Println(msg)
}

func (g *Game) Play(depth int) int {
	if _, ok := depthHistogram[depth]; !ok {
		depthHistogram[depth] = 0
	}
	depthHistogram[depth]++

	if time.Since(lastUpdate).Seconds() > 3 {
		fmt.Printf("\r%s: %v", time.Since(startTime), depthHistogram)
		lastUpdate = time.Now()
	}
	for {

		/* Abort if turn is reocurring and let player 1 win */
		if g.IsTurnReocurring() {
			g.Print(depth, "Turn reoccuring")
			return PLAYER1
		}
		g.AddTurnToHistory()

		/*
		 * Check if each player has at least as many cards remaining
		 * as the top card
		 */
		if g.p1.CanGoRecursive() && g.p2.CanGoRecursive() {
			g.Print(depth, "New Subgame")
			/* Go into recursion */
			g2 := g.NewSubGame(g.p1.cards[0], g.p2.cards[0])

			winner := g2.Play(depth + 1)
			if winner == PLAYER1 {
				g.p1.PutFirstCardToBack()
				g.p1.cards = append(g.p1.cards, g.p2.RemoveFirstCard())
				g.Print(depth, "P1 won subgame/round")
			} else {
				g.p2.PutFirstCardToBack()
				g.p2.cards = append(g.p2.cards, g.p1.RemoveFirstCard())
				g.Print(depth, "P2 won subgame/round")
			}
			continue
		}

		/* Check regular rule: Winner is the one with the higher top card*/
		if g.p1.cards[0] > g.p2.cards[0] {
			g.p1.PutFirstCardToBack()
			g.p1.cards = append(g.p1.cards, g.p2.RemoveFirstCard())
			g.Print(depth, "P1 won round")
		} else {
			g.p2.PutFirstCardToBack()
			g.p2.cards = append(g.p2.cards, g.p1.RemoveFirstCard())
			g.Print(depth, "P2 won round")
		}

		/* If a player has no card left: the other one wins */
		if len(g.p1.cards) == 0 {
			g.Print(depth, "P1 has no cards left")
			return PLAYER2
		} else if len(g.p2.cards) == 0 {
			g.Print(depth, "P2 has no cards left")
			return PLAYER1
		}
	}
}

func (g Game) WinnersScore() int {
	cards := g.p1.cards
	if len(g.p1.cards) == 0 {
		cards = g.p2.cards
	}

	score := 0
	for i, card := range cards {
		score += card * (len(cards) - i)
	}

	return score
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

	game := NewGame()
	game.p1 = player1
	game.p2 = player2

	game.Play(0)

	fmt.Println("\n", game.WinnersScore())
}
