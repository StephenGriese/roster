package roster

import "fmt"

type Position uint

const (
	Forward Position = iota
	Defense
	Goalie
)

func (p Position) String() string {
	switch p {
	case Forward:
		return "forward"
	case Defense:
		return "defense"
	case Goalie:
		return "goalie"
	default:
		return "unknown"
	}
}

type Player struct {
	ID             int      `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	SweaterNumber  int      `json:"sweaterNumber"`
	Position       Position `json:"position"`
	HeightInInches int      `json:"heightInInches"`
	WeightInPounds int      `json:"weightInPounds"`
}

func (p Player) HeightInFeetAndInches() (feet, inches int) {
	feet = p.HeightInInches / 12
	inches = p.HeightInInches % 12
	return
}

func FeetAndInchesToString(feet, inches int) string {
	return fmt.Sprintf("%d'%d\"", feet, inches)
}
