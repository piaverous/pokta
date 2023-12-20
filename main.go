package main

import (
	"fmt"
	"os"

	"github.com/piaverous/pokta/cmd"
	"github.com/piaverous/pokta/pokta"
)

func main() {
	app, err := pokta.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err)
		os.Exit(1)
	}

	if err := cmd.New(app).Execute(); err != nil {
		os.Exit(1)
	}
}
