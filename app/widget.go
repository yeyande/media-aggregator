package app

import (
    "fmt"

	"github.com/jroimartin/gocui"
)

type Widget struct {
    name string
    x, y int
    w, h int
    body string
    initialized bool
}

func (w *Widget) Layout(g *gocui.Gui) error {
    return w.render(g)
}

func (w *Widget) render(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
        fmt.Fprint(v, w.body)
	}
    return nil
}
