package app

import (
    "fmt"

	"github.com/jroimartin/gocui"
)

type ConsoleInterface interface {
    viewGetter
    SetView(name string, x0, y0, x1, y1 int) (*gocui.View, error)
    SetCurrentView(name string) (*gocui.View, error)
    SetKeybinding(viewname string, key interface{}, mod gocui.Modifier,
                    handler func(*gocui.Gui, *gocui.View) error) error
}

type viewGetter interface {
    View(name string) (*gocui.View, error)
}

type Window struct {
    x, y int
    w, h int
}

type Widget struct {
    Window
    name string
    body string
    initialized bool
}

func (w *Widget) Layout(g ConsoleInterface) error {
    return w.render(g)
}

func (w *Widget) render(g ConsoleInterface) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
        fmt.Fprint(v, w.body)
	}
    return nil
}
