package main

import (
	"fmt"
	"os"
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
		_, _ = fmt.Fprintf(os.Stdout, "%2d  %s %s\n", p.SweaterNumber, p.FirstName, p.LastName)
	}
}
