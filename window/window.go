package window

import "fmt"

type Window struct {
	width  int
	height int
	buffer []rune
}

func NewWindow(width int, height int) *Window {
	return &Window{width: width, height: height}
}

func (w *Window) Init() {
	w.buffer = w.createBuffer()
  for i := 0; i < len(w.buffer); i++ {
    if i % w.width == 0 {
      w.buffer[i] = '\n'
    } else {
      w.buffer[i] = '#'
    }
  }
}

func (w *Window) createBuffer() []rune {
	buffer := make([]rune, w.width*w.height)
	return buffer
}

func (w *Window) Draw() {
  fmt.Printf("%s", string(w.buffer))
  fmt.Println()
}
