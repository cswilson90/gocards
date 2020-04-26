package game

import "testing"

func TestDeck(t *testing.T) {
	deck := NewDeck()

	if deck.Size() != 52 {
		t.Errorf("New deck should have 52 cards but it has %v", deck.Size())
	}

	if deck.Empty() {
		t.Error("New deck thinks it's empty when it shouldn't")
	}

	lastCard, err := deck.DrawCard()
	if err != nil {
		t.Error("Error drawing first card: " + err.Error())
	}

	if deck.Size() != 51 {
		t.Errorf("Deck should have 51 cards after drawing but has %v", deck.Size())
	}

	orderedAsc := true
	orderedDesc := true

	for i := 0; i < 51; i++ {
		nextCard, err := deck.DrawCard()
		if err != nil {
			t.Error("Couldn't draw card: " + err.Error())
		}

		if nextCard.id == lastCard.id {
			t.Error("Drew the same card twice in a row")
		} else if nextCard.id < lastCard.id {
			orderedAsc = false
		} else {
			orderedDesc = false
		}
	}

	if !deck.Empty() {
		t.Error("Deck not empty after drawing all cards")
	}

	if orderedAsc || orderedDesc {
		t.Error("Cards were drawn in order when they should be shuffled")
	}

	_, err = deck.DrawCard()
	if err == nil {
		t.Error("No error given when trying to draw from empty deck")
	}

	deck.Reset()

	if deck.Size() != 52 {
		t.Errorf("Reset deck should have 52 cards but it has %v", deck.Size())
	}
}
