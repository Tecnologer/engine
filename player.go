package engine

type Player interface {
	GetName() string
	GetCards() []Card
	GetIndex() int
	GetOutput() chan TurnResult
	SendResult(TurnResult) error
}
