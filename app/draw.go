package app

import (
    "fmt"
    "github.com/nsf/termbox-go"
)

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

func PrintKey(key rune) {
	printfTb(3, 19, termbox.ColorWhite, termbox.ColorDefault, "Key: %c", key)
}

func drawChar(x int, y int, ch rune) {
    termbox.SetCell(x, y, ch, termbox.ColorWhite, termbox.ColorDefault)
}

func drawBorders(w Window) {
	for i := 1; i <= w.width; i++ {
        drawChar(i, 0, boxHoriz)
        drawChar(i, w.height, boxHoriz)
	}
	for i := 1; i < w.height; i++ {
        drawChar(0, i, boxVert)
        drawChar(w.width, i, boxVert)
	}
}

func printfTb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
    for _, c := range fmt.Sprintf(format, args...) {
		termbox.SetCell(x, y, c, fg, bg)
		x++
    }
}

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

func getConnector(cells []termbox.Cell, w Window) rune {
    return 'c'
}
