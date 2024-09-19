package main

import "game_loop/window"

func main() {
	window := window.NewWindow(5, 5)
  window.Init()
  window.Draw()
}
