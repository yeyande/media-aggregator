package app

import (
    "log"

	"github.com/jroimartin/gocui"
)


type PlaylistWidget struct { Widget }

func (w *PlaylistWidget) Layout(g *gocui.Gui) error {
    _, h := g.Size()
    w.h = h-w.y-1
    return w.render(g)
}

func setBinding(g ConsoleInterface, view string, key interface{},
                mod gocui.Modifier, cb func(g *gocui.Gui, v *gocui.View) error) {
	if err := g.SetKeybinding(view, key, mod, cb); err != nil {
		log.Panicln(err)
	}
}

func NewPlaylistWidget(x, y, h int) *PlaylistWidget {
    return &PlaylistWidget{Widget{"playlist", x, y, 20, h-y-1, "", false}}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
