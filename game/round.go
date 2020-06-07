package game

import "fmt"

// A GameRound represents a single round of a game.
type gameRound struct {
	deck    *Deck
	players []*Player

	roundNumber int
	numCards    int
	dealer      int
}

// newGameRound creates a new game round.
func newGameRound(deck *Deck, players []*Player, roundNumber int) *gameRound {
	// The dealer rotates each round
	dealer := (roundNumber - 1) % len(players)

	// The number of cards dealt increases every round
	numCards := roundNumber

	return &gameRound{
		deck:        deck,
		players:     players,
		roundNumber: roundNumber,
		numCards:    numCards,
		dealer:      dealer,
	}
}

// deal deals out the cards for the round to the players
func (r *gameRound) deal() error {
	r.deck.Reset()

	for i := 0; i < r.numCards; i++ {
		for _, player := range r.players {
			card, err := r.deck.DrawCard()
			if err != nil {
				return fmt.Errorf("Ran out of card to deal on round %v for %v players", r.roundNumber, len(r.players))
			}

			err = player.AddCardToHand(card)
			if err != nil {
				return fmt.Errorf("Error adding card to player %v's hand: %v", player.id, err.Error())
			}
		}
	}

	return nil
}
