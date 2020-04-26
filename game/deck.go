package game

import (
	"errors"
	"math/rand"
	"time"
)

// A Deck represents a deck of playing cards.
type Deck struct {
	cards []*Card
}

// NewDeck creates and returns a new Deck with shuffled cards.
func NewDeck() *Deck {
	deck := &Deck{}
	deck.Reset()

	return deck
}

// Empty returns whether the deck is empty or not.
func (d *Deck) Empty() bool {
	if d.Size() == 0 {
		return true
	} else {
		return false
	}
}

// DrawCard returns the next card in the deck.
// Trying to take a card from an empty deck gives an error
func (d *Deck) DrawCard() (*Card, error) {
	if d.Empty() {
		return nil, errors.New("Tried to take card from empty deck")
	}

	var nextCard *Card
	nextCard, d.cards = d.cards[0], d.cards[1:]

	return nextCard, nil
}

// Reset sets the number of cards in the deck back to the default number.
// It also shuffles the cards.
func (d *Deck) Reset() {
	d.cards = make([]*Card, 52)

	for i := 0; i < 52; i++ {
		d.cards[i] = newCard(uint(i))
	}

	d.Shuffle()
}

// Shuffle shuffles all cards left in the Deck.
// Trying to shuffle an empty deck gives an error
func (d *Deck) Shuffle() error {
	if d.Empty() {
		return errors.New("Tried to shuffle empty deck")
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	return nil
}

// Size returns the number of cards left in the deck.
func (d *Deck) Size() int {
	return len(d.cards)
}
