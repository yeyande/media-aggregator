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

const (
    boxHoriz = 0x2500
    boxVert = 0x2502
    boxHorizUp = 0x2534
    boxVertRight = 0x251C
    boxVertLeft = 0x2524
    boxDownHoriz = 0x252C
    boxDownRight = 0x250C
    boxDownLeft = 0x2510
    boxUpRight = 0x2514
    boxUpLeft = 0x2518
)

func drawConnectors(width, height int) {
    drawChar(20, 1, boxVert)
    drawChar(20, 0, boxDownHoriz)
    drawChar(20, height, boxHorizUp)
    drawChar(0, 2, boxVertRight)
    drawChar(width, 2, boxVertLeft)
    drawChar(0, 0, boxDownRight)
    drawChar(width, 0, boxDownLeft)
    drawChar(0, height, boxUpRight)
    drawChar(width, height, boxUpLeft)
}

func drawQueueWindow(w Window) {
    for i := w.offsety; i < w.height; i++ {
        drawChar(w.width, i, boxVert)
    }
    termbox.SetCell(w.width, 2, 0x253C, termbox.ColorWhite, termbox.ColorDefault)
}
