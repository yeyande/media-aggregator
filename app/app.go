package app

import (
    "github.com/nsf/termbox-go"
)

type Window struct {
    width int
    height int
    offsetx int
    offsety int
}

func Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func Draw() {
    width, height := termbox.Size()
    drawBorders(Window{width-1, height-1, 0, 0})
    for i := 1; i <= width; i++ {
        drawChar(i, 2, boxHoriz)
    }
    drawMenuBar(width)
    drawQueueWindow(Window{20, height, 0, 3})
    drawConnectors(width-1, height-1)
	termbox.Flush()
}

func drawQueueWindow(w Window) {
    for i := w.offsety; i < w.height; i++ {
        drawChar(w.width, i, boxVert)
    }
    termbox.SetCell(w.width, 2, 0x253C, termbox.ColorWhite, termbox.ColorDefault)
}
