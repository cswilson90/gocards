package game

// A Card represents a standard playing card
type Card struct {
	id uint

	suit  uint
	value uint
}

func newCard(id uint) *Card {
	return &Card{
		id:    id,
		suit:  id / 13,
		value: id % 13,
	}
}

// Suit returns the suit of the card as a number from 0 to 3.
func (c *Card) Suit() uint {
	return c.suit
}

// Value returns the value of the card as a number from 0 to 12.
func (c *Card) Value() uint {
	return c.value
}
