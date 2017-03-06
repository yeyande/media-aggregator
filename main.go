package main

import (
    "github.com/yeyande/media-aggregator/app"
    "github.com/nsf/termbox-go"
)

func main() {
    app.Init()
	defer termbox.Close()
    app.AddMenuItem("Local")
    app.AddMenuItem("Google Music")
    app.AddMenuItem("Pandora")
	app.Draw()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
            handleKeypress(ev)
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			app.Draw()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func handleKeypress(ev termbox.Event) {
    switch ev.Ch {
    case 'q':
        panic("exiting")
        break
    case 'l':
        break
    default:
        break
    }
    app.PrintKey(ev.Ch)
    //dispatch_press(&ev)
    //termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    app.Draw()
    termbox.Flush()
}

