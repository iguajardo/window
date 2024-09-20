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
	// topLeft, topRight := '╔', '╗'
	// bottomLeft, bottomRight := '╚', '╝'
	// horizontal, vertical := '═', '║'
	vertical := '\u2551'

	w.buffer = w.createBuffer()
	for y := 0; y < w.height+2; y++ {
		firstVerticalX := y * (w.width + 3)
		secondVerticalX := firstVerticalX + w.width + 1

		w.buffer[firstVerticalX] = vertical
		w.buffer[secondVerticalX] = vertical
		w.buffer[secondVerticalX+1] = '\n'

		for x := 1; x < w.width+1; x++ {
			w.buffer[firstVerticalX+x] = '#'
		}
	}
}

func (w *Window) createBuffer() []rune {
	buffer := make([]rune, (w.width+3)*(w.height+2))
	return buffer
}

func (w *Window) Draw() {
	fmt.Printf("%s", string(w.buffer))
	fmt.Println()
}
