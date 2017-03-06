package app

import (
    "github.com/nsf/termbox-go"
)

type MenuItem struct {
    start int
    name string
    selected bool
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
