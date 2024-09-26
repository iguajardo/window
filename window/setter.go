package window

import "fmt"

type void struct{}

var walls = map[rune]bool{
	'╝': true,
	'╗': true,
	'╔': true,
	'║': true,
	'═': true,
	'╚': true,
}

func (w *Window) Set(char rune, x int, y int) error {
  if x < 0 || y < 0 {
    return fmt.Errorf("invalid negative values for coordinates")
  }
	position := (y+1)*(w.width+3) + x + 1

	if walls[w.buffer[position]] {
		return fmt.Errorf("invalid coordinates: x -> %v, y -> %v", x, y)
	}

  w.buffer[position] = char

	return nil
}
