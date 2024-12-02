package roster

import (
	"fmt"
	"time"
)

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
	ID             int       `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	SweaterNumber  int       `json:"sweaterNumber"`
	Position       Position  `json:"position"`
	HeightInInches int       `json:"heightInInches"`
	WeightInPounds int       `json:"weightInPounds"`
	BirthDate      time.Time `json:"birthDate"`
}

func (p Player) HeightInFeetAndInches() (feet, inches int) {
	feet = p.HeightInInches / 12
	inches = p.HeightInInches % 12
	return
}

func FeetAndInchesToString(feet, inches int) string {
	return fmt.Sprintf("%d'%d\"", feet, inches)
}

func (p Player) Age() int {
	now := time.Now()
	years := now.Year() - p.BirthDate.Year()
	if now.YearDay() < p.BirthDate.YearDay() {
		years--
	}
	return years
}

func (p Player) FullName() string {
	return p.LastName + ", " + p.FirstName
}
