package engine

type Card interface {
	GetValue() string
	GetColor() string
	Play(string, []Card) (CardResult, error)
	CanPlay(string, []Card) (bool, error)
}
