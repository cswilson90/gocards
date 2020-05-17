package game

import (
	"fmt"
	"sort"
)

// A Player represents the state of a single player.
type Player struct {
	id int

	hand []*Card
}

// NewPlayer creates and returns a new Player with the given id.
func NewPlayer(id int) *Player {
	return &Player{id: id}
}

// Hand returns the hand the player currently has.
// The hand will always be sorted by suit and card value.
func (p *Player) Hand() []*Card {
	return p.hand
}

// AddCardToHand adds the given card to the player's hand.
// Returns an error if the is already in teh player's hand.
func (p *Player) AddCardToHand(card *Card) error {
	if p.findCardInHand(card) != -1 {
		fmt.Errorf("Card %v already in player %v's hand", card.id, p.id)
	}

	p.hand = append(p.hand, card)
	p.sortHand()

	return nil
}

// HandSize returns the number of cards currently in the player's hand.
func (p *Player) HandSize() int {
	return len(p.hand)
}

// HandContainsCard checks whether the given card is in the player's hand.
func (p *Player) HandContainsCard(card *Card) bool {
	return p.findCardInHand(card) != -1
}

// RemoveCardFromHand removes the given card from the player's hand.
// Returns an error if the card is not in the player's hand
func (p *Player) RemoveCardFromHand(card *Card) error {
	i := p.findCardInHand(card)

	if i == -1 {
		return fmt.Errorf("Card %v not found in player %v's hand", card.id, p.id)
	}

	// Delete card from hand
	copy(p.hand[i:], p.hand[i+1:])
	p.hand[len(p.hand)-1] = nil
	p.hand = p.hand[:len(p.hand)-1]

	return nil
}

// findCardInHand returns the index of the given card in the player's hand.
// If the hand does not contain the card, return -1.
func (p *Player) findCardInHand(card *Card) int {
	numCards := len(p.hand)

	// Do a binary search of the hand
	index := sort.Search(numCards, func(i int) bool {
		return p.hand[i].id >= card.id
	})

	if index < numCards && p.hand[index].id == card.id {
		return index
	}
	return -1
}

// sortHand sorts the hand in ascending order by card ID.
func (p *Player) sortHand() {
	sort.Slice(p.hand, func(i, j int) bool {
		return p.hand[i].id < p.hand[j].id
	})
}
