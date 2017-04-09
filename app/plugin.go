package app

type PluginDisplay interface {
    Draw(*Window, ConsoleInterface)
}

type Plugin struct {
    w, h int
}

func (p *Plugin) Draw(w *Window, cui ConsoleInterface) {
}

/*
    Plugin.Register(app.ContextManager)
    ContextManager.Draw()
        plugin.Draw(x, y, w, h)
            cui.Draw()
*/
