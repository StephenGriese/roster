package main

import (
	"fmt"
	"sort"

	"github.com/StephenGriese/roster/nhle"
)

func main() {
	s := nhle.NewPlayerService()
	roster, err := s.Players("PHI")
	if err != nil {
		panic(err)
	}
	sort.Slice(roster, func(i, j int) bool {
		return roster[i].SweaterNumber < roster[j].SweaterNumber
	})

	for _, p := range roster {
		fmt.Printf("%2d  %s %s\n", p.SweaterNumber, p.FirstName, p.LastName)
	}
}
