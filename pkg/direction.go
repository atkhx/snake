package pkg

type (
	Direction string
	Position  struct {
		X, Y int
	}
)

const (
	Up    = Direction("up")
	Down  = Direction("down")
	Left  = Direction("left")
	Right = Direction("right")
)
