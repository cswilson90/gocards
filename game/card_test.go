package game

import "testing"

func TestNewCard(t *testing.T) {
	tests := map[uint]map[string]uint{
		0: map[string]uint{
			"Suit":  0,
			"Value": 0,
		},
		12: map[string]uint{
			"Suit":  0,
			"Value": 12,
		},
		13: map[string]uint{
			"Suit":  1,
			"Value": 0,
		},
		34: map[string]uint{
			"Suit":  2,
			"Value": 8,
		},
		51: map[string]uint{
			"Suit":  3,
			"Value": 12,
		},
	}

	for id, result := range tests {
		card := newCard(id)

		expectedSuit := result["Suit"]
		expectedValue := result["Value"]

		if card.Suit() != expectedSuit {
			t.Errorf("New card ID %v expected suit %v got %v", id, expectedSuit, card.Suit())
		}

		if card.Value() != expectedValue {
			t.Errorf("New card ID %v expected value %v got %v", id, expectedValue, card.Value())
		}
	}
}
