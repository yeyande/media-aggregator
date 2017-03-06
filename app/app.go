package app

import (
    "fmt"
    "github.com/nsf/termbox-go"
)

type MenuItem struct {
    start int
    name string
    selected bool
}

type Window struct {
    width int
    height int
    offsetx int
    offsety int
}

func AddMenuItem(name string) {
    item := MenuItem{getMenuItemOffset(), name, false}
    if len(menuItems) == 0 {
        item.selected = true
    }
    menuItems = append(menuItems, item)
}

func getMenuItemOffset() int {
    offset := 22
    for _, item := range menuItems {
        offset += 3 + len(item.name)
    }
    return offset
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

func PrintKey(key rune) {
	printfTb(3, 19, termbox.ColorWhite, termbox.ColorDefault, "Key: %c", key)
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

var menuItems = []MenuItem{}

func drawMenuBar(width int) {
	printfTb(4, 1, termbox.ColorWhite|termbox.AttrBold, termbox.ColorDefault, "Media Player")
    for _, item := range menuItems {
        printMenuItem(item)
        drawChar(item.start+len(item.name)+1, 1, boxVert)
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

func drawQueueWindow(w Window) {
    for i := w.offsety; i < w.height; i++ {
        drawChar(w.width, i, boxVert)
    }
    termbox.SetCell(w.width, 2, 0x253C, termbox.ColorWhite, termbox.ColorDefault)
}

func printfTb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
    for _, c := range fmt.Sprintf(format, args...) {
		termbox.SetCell(x, y, c, fg, bg)
		x++
    }
}
