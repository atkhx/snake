package pkg

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type (
	Object any
	Field  struct {
		objects       [][]Object
		width, height int
	}
)

func NewField(width, height int) *Field {
	objects := make([][]Object, height)
	for i := range objects {
		objects[i] = make([]Object, width)
	}

	return &Field{
		objects: objects,
		width:   width,
		height:  height,
	}
}

func (f *Field) GetObject(x, y int) Object {
	return f.objects[y][x]
}

func (f *Field) SetObject(x, y int, obj Object) {
	f.objects[y][x] = obj
}

func (f *Field) GetFreeCoordinates(offset int) (x, y int) {
	for {
		x = offset + rand.Intn(len(f.objects[0])-offset)
		y = offset + rand.Intn(len(f.objects)-offset)

		if f.objects[y][x] != nil {
			continue
		}
		return
	}
}

func (f *Field) ShowField() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for y := 0; y < len(f.objects); y++ {
		for x := 0; x < len(f.objects[y]); x++ {
			if obj := f.objects[y][x]; obj != nil {
				fmt.Print(obj)
			} else {
				fmt.Print("⬜️")
			}
		}
		fmt.Print("    ")
		fmt.Println()
	}
}
