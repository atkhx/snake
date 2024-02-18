package main

import (
	"context"
	"fmt"

	"github.com/atkhx/snake/pkg"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	game := pkg.NewGame(40, 18)
	done := game.Start(ctx)

	select {
	case <-done:
	case <-ctx.Done():
	}

	fmt.Println("context is done")
}
