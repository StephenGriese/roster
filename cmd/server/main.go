package main

import (
	"context"
	"fmt"
	"os"

	"github.com/StephenGriese/roster/server"
)

var (
	builder   = "Undeclared"
	buildTime = "Undeclared"
	goversion = "Undeclared"
	version   = "Undeclared"
)

func main() {
	ctx := context.Background()

	buildInfo := server.BuildInfo{
		Builder:   builder,
		BuildTime: buildTime,
		Goversion: goversion,
		Version:   version,
	}

	if err := server.Run(ctx, os.Stdin, os.Stdout, os.Getenv, os.Getwd, buildInfo); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
