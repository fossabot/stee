package main

import (
	"os"

	"github.com/milanrodriguez/stee/command"
)

func main() {
	os.Exit(command.ServerRun())
}
