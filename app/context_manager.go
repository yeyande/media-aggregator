package app

type Plugin struct {}

type ContextManager interface {
    Register(Plugin) error
    Draw() error
}

type PanelContextManager struct {
}

func (m *PanelContextManager) Register(p Plugin) error {
    return nil
}

func (m *PanelContextManager) Draw() error {
    return nil
}
