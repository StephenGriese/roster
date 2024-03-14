package main

import (
	"context"
	"fmt"
	"github.com/StephenGriese/roster/rest"
	"io"
	"os"
	"sort"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(_ context.Context, w io.Writer, _ []string) error {
	ps := rest.NewPlayerService()
	players, err := ps.Players()
	if err != nil {
		return fmt.Errorf("error getting players: %w", err)
	}
	sort.Slice(players, func(i, j int) bool {
		return players[i].SweaterNumber < players[j].SweaterNumber
	})
	for _, b := range players {
		fmt.Fprintf(w, "%2d   %-25s %s\n", b.SweaterNumber, b.FirstName+" "+b.LastName, b.Position)
	}
	return nil

}
