package engine

//TurnResult defines the result of the turn
type TurnResult interface {
	//GetDrawCount returns the amount of card to draw for the player on turn
	GetDrawCount() int
	//GetCard returns the card played
	GetCard() Card
	//GetPlayer returns the player on turn
	GetPlayer() Player
	//GetDirection returns the current direction of the game after play a card
	GetDirection() Direction
}
