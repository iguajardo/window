package window

import "fmt"

type Window struct {
	width     int
	height    int
	buffer    []rune
	initiated bool
}

func NewWindow(width int, height int) *Window {
	return &Window{
		width:  width,
		height: height,
		buffer: make([]rune, (width+3)*(height*2)),
	}
}

func (w *Window) Init() error {
	if w.initiated {
		return fmt.Errorf("window already initialized")
	}

	topLeft, topRight := '╔', '╗'
	bottomLeft, bottomRight := '╚', '╝'
	horizontal, vertical := '═', '║'
	fill := ' '

  // fill corners
	w.buffer[0], w.buffer[w.width+1] = topLeft, topRight
	w.buffer[(w.height+1)*(w.width+3)], w.buffer[(w.height+1)*(w.width+3)+w.width+1] = bottomLeft, bottomRight

  // fill top and bottom horizontal
	for x := 1; x <= w.width; x++ {
		w.buffer[x] = horizontal
		w.buffer[(w.height+1)*(w.width+3)+x] = horizontal
	}

  // fill vertical borders and content area
  for y := 1; y <= w.height; y++ {
    rowStart := y * (w.width+3)
    w.buffer[rowStart] = vertical
    w.buffer[rowStart+w.width+1] = vertical

    // fill content space
    for x := 1; x <= w.width; x++ {
      w.buffer[rowStart+x] = fill
    }

    // new line at the end of each row
    w.buffer[rowStart+w.width+2] = '\n'
  }

  w.buffer[w.width+2] = '\n'
  w.buffer[(w.height+1)*(w.width+3)+w.width+2] = '\n'

	w.initiated = true

	return nil
}

func (w *Window) Draw() error {
  if !w.initiated {
    return fmt.Errorf("window not initialized")
  }
  fmt.Print(string(w.buffer))

  return nil
}
