package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
)

func main() {
    init_termbox()
	defer termbox.Close()
	draw_main()
	for {
        termbox.PollEvent()
        break
        /*
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
            handle_keypress(ev)
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw_main()
		case termbox.EventMouse:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw_main()
		case termbox.EventError:
			panic(ev.Err)
		}
        fmt.Printf("Hi")
        */
	}
}
type Window struct {
    width int
    height int
    offsetx int
    offsety int
}

const (
    BOX_HORIZ = 0x2500
    BOX_VERT = 0x2502
    BOX_HORIZ_UP = 0x2534
    BOX_VERT_RIGHT = 0x251C
    BOX_VERT_LEFT = 0x2524
    BOX_DOWN_HORIZ = 0x252C
    BOX_DOWN_RIGHT = 0x250C
    BOX_DOWN_LEFT = 0x2510
    BOX_UP_RIGHT = 0x2514
    BOX_UP_LEFT = 0x2518
)

func draw_main() {
    width, height := termbox.Size()
    draw_borders(Window{width-1, height-1, 0, 0})
    for i := 1; i <= width; i++ {
        draw_char(i, 2, BOX_HORIZ)
    }
    draw_queue_window(Window{20, height, 0, 3})
    draw_connectors(width-1, height-1)
	printf_tb(4, 1, termbox.ColorWhite|termbox.AttrBold, termbox.ColorDefault, "Media Player")
	termbox.Flush()
}

func draw_connectors(width, height int) {
    draw_char(20, 1, BOX_VERT)
    draw_char(20, 0, BOX_DOWN_HORIZ)
    draw_char(20, height, BOX_HORIZ_UP)
    draw_char(0, 2, BOX_VERT_RIGHT)
    draw_char(width, 2, BOX_VERT_LEFT)
    draw_char(0, 0, BOX_DOWN_RIGHT)
    draw_char(width, 0, BOX_DOWN_LEFT)
    draw_char(0, height, BOX_UP_RIGHT)
    draw_char(width, height, BOX_UP_LEFT)
}

func draw_char(x int, y int, ch rune) {
    termbox.SetCell(x, y, ch, termbox.ColorWhite, termbox.ColorDefault)
}

func draw_borders(w Window) {
	for i := 1; i <= w.width; i++ {
        draw_char(i, 0, BOX_HORIZ)
        draw_char(i, w.height, BOX_HORIZ)
	}
	for i := 1; i < w.height; i++ {
        draw_char(0, i, BOX_VERT)
        draw_char(w.width, i, BOX_VERT)
	}
}

func draw_queue_window(w Window) {
    for i := w.offsety; i < w.height; i++ {
        draw_char(w.width, i, BOX_VERT)
    }
    termbox.SetCell(w.width, 2, 0x253C, termbox.ColorWhite, termbox.ColorDefault)
}

func handle_keypress(ev termbox.Event) {
    fmt.Printf("Pressed key")
    print_key(ev.Key)
    //dispatch_press(&ev)
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    draw_main()
    termbox.Flush()
}

func print_key(key termbox.Key) {
	printf_tb(3, 19, termbox.ColorWhite, termbox.ColorDefault, "Key: ")
	printf_tb(8, 19, termbox.ColorYellow, termbox.ColorDefault, "decimal: %d", key)
}

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}

func init_termbox() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

