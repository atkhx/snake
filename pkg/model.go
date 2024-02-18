package pkg

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Game struct {
	Snake  *Snake
	Field  *Field
	Timer  *time.Timer
	Paused bool

	width, height int
}

const appleTTL = 20
const gameTick = 200 * time.Millisecond

func NewGame(width, height int) *Game {
	field := NewField(width, height)

	x, y := width/2, height/2

	snake := NewSnake(x, y, field)
	field.objects[y][x] = snake

	return &Game{
		Snake:  snake,
		Field:  field,
		width:  width,
		height: height,
	}
}

var keyToDirection = map[termbox.Key]Direction{
	termbox.KeyArrowUp:    Up,
	termbox.KeyArrowDown:  Down,
	termbox.KeyArrowLeft:  Left,
	termbox.KeyArrowRight: Right,
}

func (g *Game) Start(ctx context.Context) chan struct{} {
	g.Timer = time.NewTimer(gameTick)
	doneChan := make(chan struct{})

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				if direction, ok := keyToDirection[ev.Key]; ok {
					g.Snake.ChangeDirection(direction)
					continue
				}

				switch ev.Key {
				case termbox.KeyEsc:
					cancel()
				case termbox.KeySpace:
					if g.Paused {
						g.Paused = false
					} else {
						g.Paused = true
					}
				}
			}
		}
	}()

	go func() {
		defer close(doneChan)
		for {
			select {
			case <-ctx.Done():
				return
			case <-g.Timer.C:
				if g.Snake.IsDead {
					fmt.Println("Game Over: snake is dead")
					return
				}

				{ // apple generator
					if probability := rand.Float64(); probability > 0.3 {
						x, y := g.Field.GetFreeCoordinates(0)
						apple := NewApple(x, y, appleTTL, g.Field)
						apple.ScheduleTicks(g, gameTick)
					}
				}

				// Make snake moves
				if !g.Paused {
					g.Snake.Tick()
				}

				g.Timer.Reset(gameTick)
				g.Field.ShowField()
			}
		}
	}()

	return doneChan
}
