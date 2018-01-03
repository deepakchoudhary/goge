package goge

import (
  "goge/extras"
)

type __Framer__ struct {
  Grid   *__Grid__
  Delay  int
}

func (framer *__Framer__) Frame() *__Framer__ {
  extras.ClearScreen()
  framer.Grid.__UpdatePos__()
  framer.Grid.Draw()
  extras.Delay(framer.Delay)
  
  return framer
}

func (grid *__Grid__) Framer(millis int) *__Framer__ {
  return &__Framer__{Grid: grid, Delay: millis}
}