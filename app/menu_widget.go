package app

import (
	"github.com/jroimartin/gocui"
)

type MenuWidget struct {
    Widget
    p Panel
}

type Panel struct {
    Widget
}

func (w *MenuWidget) highlight(g ConsoleInterface) {
    w.setEntryAttrs(g, gocui.ColorWhite, gocui.ColorBlack | gocui.AttrBold)
}

func (w *MenuWidget) dehighlight(g ConsoleInterface) {
    w.setEntryAttrs(g, gocui.ColorBlack, gocui.ColorWhite)
}

func (w *MenuWidget) deactivate(g ConsoleInterface) {
}

func (w *MenuWidget) activate(g ConsoleInterface) {
}

func (w *MenuWidget) setEntryAttrs(g viewGetter, bg, fg gocui.Attribute) {
    if view, err := g.View(w.name); err == nil {
        view.BgColor, view.FgColor = bg, fg
    }
}

func NewMenuWidget(name, body string) *MenuWidget {
    return &MenuWidget{Widget{name, 0, 0, len(body)+1, 0, body, false}, Panel{}}
}
