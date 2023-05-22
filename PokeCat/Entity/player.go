package entity
type Player struct {
	Entity
	Username string
	Password string
	PokeList []string
	// movement Movement
}

const (
	W string = "w"
	A string = "a"
	S string = "s"
	D string = "d"
	E string = "e"
)