package game

import (
	"reflect"
	"testing"
)

func PlayerTest(t *testing.T) {
	player := NewPlayer(0)

	if player.HandSize() != 0 || player.Hand() != nil {
		t.Error("New player's hand is not empty")
	}

	card0 := newCard(0)
	if player.HandContainsCard(card0) {
		t.Error("New player's hand should not contain card 0")
	}

	err := player.AddCardToHand(card0)
	if err != nil {
		t.Error("Error adding card 0 to player's empty hand")
	}

	err = player.AddCardToHand(card0)
	if err == nil {
		t.Error("No error when adding duplicate card to player's hand")
	}

	handSize := player.HandSize()
	if handSize != 1 {
		t.Errorf("Incorrect hand size after adding a card. Expected 1, got %v", handSize)
	}

	card2 := newCard(2)
	err = player.AddCardToHand(card2)

	card32 := newCard(32)
	err = player.AddCardToHand(card32)

	card13 := newCard(13)
	err = player.AddCardToHand(card13)

	handSize = player.HandSize()
	if handSize != 4 {
		t.Errorf("Incorrect hand size after adding a card. Expected 4, got %v", handSize)
	}

	if !player.HandContainsCard(card13) {
		t.Error("Hand does not contain card 13 after it was added")
	}

	expectedHand := []*Card{card0, card2, card13, card32}
	if !reflect.DeepEqual(player.Hand(), expectedHand) {
		t.Errorf("Did not get expected hand after adding 4 cards: got %v, expected: %v", player.Hand(), expectedHand)
	}

	card4 := newCard(4)
	err = player.RemoveCardFromHand(card4)
	if err == nil {
		t.Error("Did not get error when trying to remove missing card from hand")
	}

	err = player.RemoveCardFromHand(card2)
	if err != nil {
		t.Error("Got error trying to remove card 2 from hand")
	}

	handSize = player.HandSize()
	if handSize != 3 {
		t.Errorf("Got hand size of %v after removing a card, expected 3", handSize)
	}
}
