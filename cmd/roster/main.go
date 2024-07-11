package main

import (
	"context"
	"fmt"
	"os"

	"github.com/StephenGriese/roster/app"
)

func main() {
	ctx := context.Background()
	if err := app.Run(ctx, os.Stdin, os.Stdout, os.Stderr, os.Getenv, os.Getwd); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
