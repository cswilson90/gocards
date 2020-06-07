package game

import "fmt"

// A gameTrick represents a single trick in a game.
// A trick id one card played by each player in order.
type gameTrick struct {
	players     []*Player
	playedCards []*Card
	trumpSuit   uint
}

// newGameTrick creates a new trick.
// Accepts a list of players in the trick and the player who goes first.
// Also accepts a trump suit, set to nil for no trump suit.
func newGameTrick(players []*Player, firstPlayer *Player, trumpSuit uint) *gameTrick {
	// Order players slice for trick in order of play
	orderedPlayers := make([]*Player, len(players))

	var firstPlayerIndex int
	for i, player := range players {
		if player.id == firstPlayer.id {
			firstPlayerIndex = i
			break
		}
	}

	for i := 0; i < len(players); i++ {
		orderedPlayers[i] = players[(firstPlayerIndex+i)%len(players)]
	}

	return &gameTrick{
		players:     orderedPlayers,
		playedCards: make([]*Card, 0, len(players)),
		trumpSuit:   trumpSuit,
	}
}

// playCard plays a card in the trick for the given player.
// Will return an error if all cards in the trick have been played.
// Will return an error if player is trying to play out of turn.
func (t *gameTrick) playCard(player *Player, card *Card) error {
	nextPlayer := t.nextPlayer()
	if nextPlayer == nil {
		return fmt.Errorf("Tried to play card in complete trick")
	}

	if nextPlayer.id != player.id {
		return fmt.Errorf("Tried to play card for player %v but player %v is next", player.id, nextPlayer.id)
	}

	leadSuit, err := t.leadSuit()
	if err != nil {
		if card.Suit() != leadSuit && player.hasCardOfSuit(leadSuit) {
			return fmt.Errorf("Player %v not following suit when they could", player.id)
		}
	}

	err = player.RemoveCardFromHand(card)
	if err != nil {
		return fmt.Errorf("Tried to play card not in player %v's hand", player.id)
	}

	t.playedCards = append(t.playedCards, card)
	return nil
}

// allCardsPlayed returns whether all cards in the trick have been played.
func (t *gameTrick) allCardsPlayed() bool {
	return len(t.players) == len(t.playedCards)
}

// nextPlayer returns the next player who should play a card.
// If the trick is complete returns nil.
func (t *gameTrick) nextPlayer() *Player {
	if t.allCardsPlayed() {
		return nil
	}

	return t.players[len(t.playedCards)]
}

// leadSuit returns the suit that was lead in this trick.
// Returns an error if no card has been played in the trick.
func (t *gameTrick) leadSuit() (uint, error) {
	if len(t.playedCards) == 0 {
		return 0, fmt.Errorf("Can't get lead suit of trick as no card has been played")
	}

	return t.playedCards[0].Suit(), nil
}

// checkForWinner checks if the trick is finished and returns the winner.
// Returns nil if the trick has not yet finished
func (t *gameTrick) checkForWinner() *Player {
	if !t.allCardsPlayed() {
		return nil
	}

	winnerIndex := 0
	winningCard := t.playedCards[0]

	for i := 1; i < len(t.players); i++ {
		nextCard := t.playedCards[i]

		if (nextCard.Suit() == winningCard.Suit() && nextCard.Value() > winningCard.Value()) || nextCard.Suit() == t.trumpSuit {
			winnerIndex = i
			winningCard = nextCard
		}
	}

	return t.players[winnerIndex]
}

// PlayedCards returns a slice containing all cards played in the trick.
func (t *gameTrick) PlayedCards() []*Card {
	return t.playedCards
}
