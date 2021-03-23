//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//"fmt"
type Suit uint
type Rank uint
//Represents card suits
const (
	Diamond Suit = iota
	Heart
	Spade
	Club
	Joker
)
var suits = [...]Suit{Diamond, Heart, Spade, Club}
//Represents card ranks
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)


const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

type Deck []Card
//New function creates a new deck of 52 cards 
//and execute on card all options, that were provided
func New(opts...func([]Card) []Card) Deck{
	deck := make([]Card, 0, 52)
	for _, suit := range suits{
		for rank := minRank; rank <= maxRank; rank++{
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts{
		deck = opt(deck)
	}
	return deck
}

func (c Card) String() string {
	if c.Suit == Joker{
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func absRank (c Card) int{
	return int(c.Suit) * int(maxRank) + int(c.Rank)
}

//DefaultSort function sort deck in default order
func DefaultSort(deck []Card) []Card{
	sort.Slice(deck, func(i, j int) bool{
		return absRank(deck[i]) < absRank(deck[j])
	})
	return deck
}

func Sort(deck []Card, less func(i, j int) bool) []Card{
	sort.Slice(deck, less)
	return deck
}
//Shuffle function gives cars random positions in the deck
func ShuffleDeck(deck []Card) []Card{
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(deck), func(i, j int){
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}
//AddJockers function add n joker in the end of the deck
func AddJokers(n int) func([]Card) []Card{
	return func(deck []Card) []Card{
		for i := 0; i < n; i++{
			deck = append(deck, Card{Suit: Joker, Rank: Rank(i)})
		} 
		return deck
	}
}
//CreateDeck function create simple deck or multideck
func CreateDeck(n int) func([]Card) []Card{
	return func(deck []Card) []Card{
		deck = append(deck, deck...)
		var multideck []Card
		for i := 0; i < n; i++{
			multideck = append(multideck, deck...)
		}
		return multideck
	}
}
//Filter function returns function to filter deck according f function
func Filter(f func(card Card) bool) func([]Card) []Card{
	return func(deck []Card) []Card {
			var newDeck []Card
			for _, card := range deck{
				if f(card){
					newDeck = append(newDeck, card)
				}
			}
		return newDeck
	}
}