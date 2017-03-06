package app

import (
    "fmt"
    "github.com/nsf/termbox-go"
)

func Draw() {
    width, height := termbox.Size()
    drawBorders(Window{width-1, height-1, 0, 0})
    for i := 1; i <= width; i++ {
        drawChar(i, 2, BoxHoriz)
    }
    drawMenuBar(width)
    drawQueueWindow(Window{20, height, 0, 3})
    drawConnectors(width-1, height-1)
	termbox.Flush()
}
type Window struct {
    width int
    height int
    offsetx int
    offsety int
}

const (
    BoxHoriz = 0x2500
    BoxVert = 0x2502
    BoxHorizUp = 0x2534
    BoxVertRight = 0x251C
    BoxVertLeft = 0x2524
    BoxDownHoriz = 0x252C
    BoxDownRight = 0x250C
    BoxDownLeft = 0x2510
    BoxUpRight = 0x2514
    BoxUpLeft = 0x2518
)

type MenuItem struct {
    start int
    name string
    selected bool
}

var menuItems = []MenuItem{
    {22, "Local", true},
    {30, "Google Music", false},
    {45, "Pandora", false},
}

func drawMenuBar(width int) {
	printfTb(4, 1, termbox.ColorWhite|termbox.AttrBold, termbox.ColorDefault, "Media Player")
    for _, item := range menuItems {
        printMenuItem(item)
        drawChar(item.start+len(item.name)+1, 1, BoxVert)
    }
}

func printMenuItem(item MenuItem) {
    font := termbox.ColorWhite
    var bgcolor termbox.Attribute
    if (item.selected) {
        font |= termbox.AttrBold
        bgcolor = termbox.ColorBlack
    } else {
        bgcolor = termbox.ColorDefault
    }
	printfTb(item.start, 1, font, bgcolor, item.name)
}

func drawConnectors(width, height int) {
    drawChar(20, 1, BoxVert)
    drawChar(20, 0, BoxDownHoriz)
    drawChar(20, height, BoxHorizUp)
    drawChar(0, 2, BoxVertRight)
    drawChar(width, 2, BoxVertLeft)
    drawChar(0, 0, BoxDownRight)
    drawChar(width, 0, BoxDownLeft)
    drawChar(0, height, BoxUpRight)
    drawChar(width, height, BoxUpLeft)
}

func drawChar(x int, y int, ch rune) {
    termbox.SetCell(x, y, ch, termbox.ColorWhite, termbox.ColorDefault)
}

func drawBorders(w Window) {
	for i := 1; i <= w.width; i++ {
        drawChar(i, 0, BoxHoriz)
        drawChar(i, w.height, BoxHoriz)
	}
	for i := 1; i < w.height; i++ {
        drawChar(0, i, BoxVert)
        drawChar(w.width, i, BoxVert)
	}
}

func drawQueueWindow(w Window) {
    for i := w.offsety; i < w.height; i++ {
        drawChar(w.width, i, BoxVert)
    }
    termbox.SetCell(w.width, 2, 0x253C, termbox.ColorWhite, termbox.ColorDefault)
}

func PrintKey(key termbox.Key) {
	printfTb(3, 19, termbox.ColorWhite, termbox.ColorDefault, "Key: ")
	printfTb(8, 19, termbox.ColorYellow, termbox.ColorDefault, "decimal: %d", key)
}

func printTb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printfTb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	printTb(x, y, fg, bg, s)
}

func Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}
