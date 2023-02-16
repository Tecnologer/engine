package engine

type Game interface {
	//New creates a new instance the output channel of the game
	New() error
	//GetMetadata returns the Metadata of the game
	GetMetadata() Metadata
	//Start selects the first card and the first player
	Start() (Card, Player, error)
	//Shuffle shuffles the draw pile
	Shuffle(int)
	//GetDiscardedPile returns the cards that was discarded
	GetDiscardedPile() []Card
	//GetDrawPile returns the cards that can be drawn
	GetDrawPile() []Card
	//DrawCard draws a card from the draw pile
	DrawCard() Card
	//PlayCard players plays a card
	PlayCard(Player, Card) error
	//AddPlayers registers a player to the game, the string represents the name of the player
	AddPlayer(string) error
	//GetPlayers returns the list of players registered on the game
	GetPlayers() []Player
	//GetCurrentPlayer returns the player that is his turn
	GetCurrentPlayer() Player
	//SetNextPlayer set the next player as current player
	SetNextPlayer(Player)
	//SayUno the current player says UNO
	SayUno(Player)
	//GetDirection returns the current direction of the game
	GetDirection() Direction
	//SetDirection changes the current direction of the game
	SetDirection(Direction)
	//Close closes the output channel, so closes the game
	Close()
}
