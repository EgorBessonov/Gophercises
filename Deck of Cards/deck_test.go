package deck

import (
	"fmt"
	"testing"
)

func CardExample() {
	fmt.Println(Card{Rank: Ace, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Diamond
	// Joker
}

func TestDefaultSort(t *testing.T) {
	deck := New(DefaultSort)
	cardExample := Card{Suit: Diamond, Rank: Ace}
	if deck[0] != cardExample {
		t.Error("Expected card: Ace of Diamond, received card:", deck[0])
	}
}

func TestAddJockers(t *testing.T) {
	deck := New(AddJokers(2))
	i := 0
	for _, card := range deck {
		if card.Suit == Joker {
			i++
		}
	}
	if i != 2 {
		t.Error("Expected jokers: 2, received:", i)
	}
}

func TestCreateDeck(t *testing.T) {
	deck := New(CreateDeck(2))
	if len(deck) != 104 {
		t.Errorf("Expected 104 cards, received %v", len(deck))
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		if card.Rank == King || card.Suit == Diamond {
			return false
		}
		return true
	}
	deck := New(Filter(filter))
	for _, card := range deck {
		if card.Rank == King || card.Suit == Diamond {
			t.Error("Expected all Kings and Diamond would be denied")
		}
	}
}
