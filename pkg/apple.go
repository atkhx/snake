package pkg

import "time"

type Apple struct {
	TTL, X, Y int
	IsDead    bool
	field     *Field
}

func NewApple(x, y, ttl int, field *Field) *Apple {
	apple := &Apple{
		TTL:   ttl,
		X:     x,
		Y:     y,
		field: field,
	}
	field.SetObject(x, y, apple)
	return apple
}

func (a *Apple) String() string {
	return "🍏"
}

func (a *Apple) ScheduleTicks(g *Game, d time.Duration) {
	t := time.NewTimer(d)

	isAppleOnPlace := func(x, y int) bool {
		if obj := a.field.GetObject(x, y); obj != nil {
			if _, ok := obj.(*Apple); !ok {
				return false
			}
		}
		return true
	}

	go func() {
		for {
			select {
			case <-t.C:
				if !g.Paused {
					if a.TTL--; a.TTL < 0 {
						if isAppleOnPlace(a.X, a.Y) {
							a.field.SetObject(a.X, a.Y, nil)
						}
						return
					}
				}
				t.Reset(d)
			}
		}
	}()
}
