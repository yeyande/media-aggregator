package app

import (
    "fmt"
    "github.com/nsf/termbox-go"
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
