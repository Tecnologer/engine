package engine

//CardResult defines the result for the game after play a card
type CardResult interface {
	//IsReverse returns true if the card change the direction of the game
	IsReverse() bool
	//IsSkip returns true if the card will cause the next player be skiped
	IsSkip() bool
	//GetDrawCount returns the number of card that the next player should draw
	GetDrawCount() int
	//GetNextColor returns the next color of the game
	GetNextColor() string
	//GetDeck returns the deck after a card played
	GetDeck() []Card
}
