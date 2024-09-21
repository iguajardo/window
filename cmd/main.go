package main

import (
	"fmt"
	"window/window"
)

func main() {
	window := window.NewWindow(13, 5)
  err := window.Init()
  if err != nil {
    fmt.Printf("error initializing window %v\n", err)
  }
  window.Draw()
  window.Draw()
  fmt.Println()
}
