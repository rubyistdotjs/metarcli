package main

import (
	"fmt"
	"os"

	"github.com/rubyistdotjs/metarcli/internal/cmd"
)

func main() {
	command := cmd.NewRootCmd()
	commandErr := command.Execute()

	if commandErr != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error '%s'", commandErr)
		os.Exit(1)
	}
}
