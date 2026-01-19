// main.go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zeebo/clingy"
)

func main() {
	ctx := context.Background()
	ok, err := clingy.Environment{}.Run(ctx, func(cmds clingy.Commands) {
		cmds.New("run", "analyze packages for issues", new(cmdRun))
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	if !ok || err != nil {
		os.Exit(1)
	}
}
