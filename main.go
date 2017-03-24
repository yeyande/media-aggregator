package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/yeyande/media-aggregator/app"
)

const delta = 0.2

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

    _, h := g.Size()
    title := app.NewTitleWidget(g, "title", 0, 0, "Media Player")
    title.AddMenuWidget(app.NewMenuWidget("gplay", "Google Music"))
    title.AddMenuWidget(app.NewMenuWidget("pandora", "Pandora"))
    title.AddMenuWidget(app.NewMenuWidget("local", "Local"))
    playlist := app.NewPlaylistWidget(0, 3, h)
    g.SetManager(title, playlist)

    setKeybinds(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func setKeybinds(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
}


const helpText = `KEYBINDINGS
Tab: Move between buttons
Enter: Push button
^C: Exit`
