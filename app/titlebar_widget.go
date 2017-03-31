package app

import (
	"github.com/jroimartin/gocui"
)

type titlebarWidget struct {
    Widget
    selectedEntry int
    menuEntries []MenuWidget
}

func NewTitleWidget(g ConsoleInterface, name string,
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
        w.menuEntries[w.selectedEntry].highlight(g)
        g.SetCurrentView(w.name)
        w.initialized = true
    }
	return nil
}

func (w *titlebarWidget) setBindings(g ConsoleInterface) {
    setBinding(g, w.name, 'h', gocui.ModNone, w.selectLeft)
    setBinding(g, w.name, 'l', gocui.ModNone, w.selectRight)
    setBinding(g, w.name, 'q', gocui.ModNone, quit)
}

func (w *titlebarWidget) selectLeft(g *gocui.Gui, v *gocui.View) error {
    return w.changeMenuEntry(g, -1)
}

func (w *titlebarWidget) selectRight(g *gocui.Gui, v *gocui.View) error {
    return w.changeMenuEntry(g, 1)
}

func (w *titlebarWidget) changeMenuEntry(g ConsoleInterface, dir int) error {
    w.menuEntries[w.selectedEntry].dehighlight(g)
    w.shiftSelectedEntry(dir)
    w.menuEntries[w.selectedEntry].highlight(g)
	return nil
}

func (w *titlebarWidget) shiftSelectedEntry(pos int) {
    if len(w.menuEntries) != 0 && len(w.menuEntries) != 1 {
        switch {
        case w.selectedEntry == 0 && pos < 0:
            w.selectedEntry = len(w.menuEntries) - 1
        case w.selectedEntry == len(w.menuEntries) - 1 && pos > 0:
            w.selectedEntry = 0
        default:
            w.selectedEntry += pos
        }
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
