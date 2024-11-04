package main

import (
	"context"
	"fmt"
	_ "go.uber.org/automaxprocs"
	"os"
)

const (
	exitOK = iota
	exitError
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(exitError)
	}
	os.Exit(exitOK)
}

func run(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
}
