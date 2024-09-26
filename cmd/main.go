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

  err = window.Set('X', 13, 4)
  if err != nil {
    panic(fmt.Sprintf("error settings values: %v", err))
  }
  window.Draw()
  fmt.Println()
}
