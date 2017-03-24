package app

import (
	"github.com/jroimartin/gocui"
)

type titlebarWidget struct {
    Widget
    selectedEntry int
    menuEntries []MenuWidget
}

func NewTitleWidget(g *gocui.Gui, name string,
                    x, y int, body string) *titlebarWidget {
    titlebar := titlebarWidget{Widget{name, x, y, 20, 2, body, false}, 0, []MenuWidget{}}
    return &titlebar
}

func (w *titlebarWidget) Layout(g *gocui.Gui) error {
    w.render(g)
    for _, e := range w.menuEntries {
        e.render(g)
    }
    if !w.initialized {
        w.setBindings(g)
        g.SetCurrentView(w.name)
        w.selectEntry(g)
        w.initialized = true
    }
	return nil
}

func (w *titlebarWidget) setBindings(g *gocui.Gui) {
    setBinding(g, w.name, 'h', gocui.ModNone, w.selectLeft)
    setBinding(g, w.name, 'q', gocui.ModNone, quit)
}

func (w *titlebarWidget) selectLeft(g *gocui.Gui, v *gocui.View) error {
    w.deselectEntry(g)
    if w.selectedEntry == 0 {
        w.selectedEntry = len(w.menuEntries) - 1
    } else {
        w.selectedEntry--
    }
    w.selectEntry(g)
	return nil
}

func (w *titlebarWidget) deselectEntry(g *gocui.Gui) {
    w.setEntryAttrs(g, gocui.ColorBlack, gocui.ColorWhite)
}

func (w *titlebarWidget) selectEntry(g *gocui.Gui) {
    w.setEntryAttrs(g, gocui.ColorWhite, gocui.ColorBlack | gocui.AttrBold)
}

func (w *titlebarWidget) setEntryAttrs(g *gocui.Gui, bg, fg gocui.Attribute) {
    if view, err := g.View(w.menuEntries[w.selectedEntry].name); err == nil {
        view.BgColor = bg
        view.FgColor = fg
    }
}

func (w *titlebarWidget) AddMenuWidget(e *MenuWidget) {
    x := w.x + w.w + 1
    for _, e := range w.menuEntries {
        x += e.w +1
    }
    e.x, e.y, e.h = x, w.y, w.h
    w.menuEntries = append(w.menuEntries, *e)
}
