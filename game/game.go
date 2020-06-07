package game

import "fmt"

// A Game represents a single card game.
type Game struct {
	deck    *Deck
	players []*Player

	roundNumber int
	maxRounds   int

	currentRound *gameRound
}

// NewGame creates a new game with the given number of players.
func NewGame(numPlayers int) (*Game, error) {

	if numPlayers < 3 {
		return nil, fmt.Errorf("%v players less than min of 3", numPlayers)
	}

	if numPlayers > 7 {
		return nil, fmt.Errorf("%v players more than max of 7", numPlayers)
	}

	// Initialise players slice
	players := make([]*Player, numPlayers)
	for i, _ := range players {
		players[i] = NewPlayer(i + 1)
	}

	return &Game{
		deck:      NewDeck(),
		players:   players,
		maxRounds: 7,
	}, nil
}

// Players returns the list of players in the game.
func (g *Game) Players() []*Player {
	return g.players
}

// StartRound starts the next round and deals cards to players
func (g *Game) StartRound() error {
	g.roundNumber++

	if g.roundNumber > g.maxRounds {
		return fmt.Errorf("Tried to start round %v which is more than the max number of rounds of %v", g.roundNumber, g.maxRounds)
	}

	round := newGameRound(g.deck, g.players, g.roundNumber)
	g.currentRound = round
	round.deal()

	return nil
}
