package pkg

import (
	"math/rand"
)

type Snake struct {
	Body []Position

	Direction Direction
	IsDead    bool
	Field     *Field
}

func NewSnake(x, y int, field *Field) *Snake {
	return &Snake{
		Body:      []Position{{X: x, Y: y}},
		Direction: []Direction{Up, Down, Left, Right}[rand.Intn(4)],
		Field:     field,
	}
}

var opposite = map[Direction]Direction{
	Up:    Down,
	Down:  Up,
	Left:  Right,
	Right: Left,
}

func (s *Snake) ChangeDirection(direction Direction) {
	if len(s.Body) == 1 || s.Direction != opposite[direction] {
		s.Direction = direction
	}
}

func (s *Snake) String() string {
	if s.IsDead {
		return "😵"
	}

	return "🟩"
}

func (s *Snake) Tick() {
	var length = len(s.Body)

	var (
		nx, ny int                                      // new coordinates
		cx, cy = s.Body[0].X, s.Body[0].Y               // current head coordinates
		tx, ty = s.Body[length-1].X, s.Body[length-1].Y // last tail piece coordinates
	)

	switch s.Direction {
	case Up:
		nx, ny = cx, cy-1
	case Down:
		nx, ny = cx, cy+1
	case Left:
		nx, ny = cx-1, cy
	case Right:
		nx, ny = cx+1, cy
	}

	if nx < 0 || nx >= s.Field.width || ny < 0 || ny >= s.Field.height {
		// No move, just dead :'(
		s.IsDead = true
		return
	}

	if obj := s.Field.GetObject(nx, ny); obj != nil {
		if _, ok := obj.(*Apple); ok {
			s.Body = append(s.Body, Position{})
		} else {
			s.IsDead = true
			return
		}
	}

	copy(s.Body[1:], s.Body)

	s.Body[0].X = nx
	s.Body[0].Y = ny

	s.Field.SetObject(tx, ty, nil)
	s.Field.SetObject(nx, ny, s)
}
