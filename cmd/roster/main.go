package main

import (
	"context"
	"fmt"
	"os"

	"github.com/StephenGriese/roster/app"
)

var (
	builder   = "Undeclared"
	buildTime = "Undeclared"
	goversion = "Undeclared"
	version   = "Undeclared"
)

func main() {
	ctx := context.Background()

	buildInfo := app.BuildInfo{
		Builder:   builder,
		BuildTime: buildTime,
		Goversion: goversion,
		Version:   version,
	}

	if err := app.Run(ctx, os.Stdin, os.Stdout, os.Stderr, os.Getenv, os.Getwd, buildInfo); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
