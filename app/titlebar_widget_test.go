package app

import "testing"
import "reflect"
import "github.com/jroimartin/gocui"

/**
  * A new menu widget should be placed at the same height of the title bar and
  * properly spaced after the last entry
  */
func TestAddMenuWidget(t *testing.T) {
    cases := []struct {
        in []*MenuWidget
        want []MenuWidget
    }{
        // Case 1 : Single menu item
        {
            []*MenuWidget{createMenuWidget("foo", 0, 0, 3, 0, "foo")},
            []MenuWidget{*createMenuWidget("foo", 21, 0, 3, 2, "foo")},
        },
        // Case 2: Two menu items
        {
            []*MenuWidget{
                createMenuWidget("foo", 0, 0, 3, 0, "foo"),
                createMenuWidget("bar", 0, 0, 3, 0, "bar"),
            },
            []MenuWidget{
                *createMenuWidget("foo", 21, 0, 3, 2, "foo"),
                *createMenuWidget("bar", 25, 0, 3, 2, "bar"),
            },

        },
        // Case 3: Three menu items of varying body lengths
        {
            []*MenuWidget{
                createMenuWidget("foo", 0, 0, 3, 0, "foo"),
                createMenuWidget("bar", 0, 0, 4, 0, "barn"),
                createMenuWidget("baz", 0, 0, 3, 0, "baz"),
            },
            []MenuWidget{
                *createMenuWidget("foo", 21, 0, 3, 2, "foo"),
                *createMenuWidget("bar", 25, 0, 4, 2, "barn"),
                *createMenuWidget("baz", 30, 0, 3, 2, "baz"),
            },

        },
    }
    for _, c := range cases {
        title := titlebarWidget{
            Widget{"title", 0, 0, 20, 2, "title", false},
            0,
            []MenuWidget{},
        }
        for _, in := range c.in {
            title.AddMenuWidget(in)
        }
        if !reflect.DeepEqual(title.menuEntries, c.want) {
            t.Errorf("TitlebarWidget.AddMenuWidget(%q) should cause %q, got %q",
                    c.in, c.want, title.menuEntries)
        }
    }
}

type mockGui struct { gocui.Gui }

func (g *mockGui) View(name string) (*gocui.View, error) {
    return &gocui.View{}, nil
}

func TestSelectLeft(t *testing.T) {
    title := titlebarWidget{
        Widget{"title", 0, 0, 20, 2, "title", false},
        0,
        []MenuWidget{
            *createMenuWidget("foo", 0, 0, 3, 0, "foo"),
            *createMenuWidget("bar", 0, 0, 3, 0, "bar"),
        },
    }
    title.selectLeft(&mockGui{}, nil)
}

func createMenuWidget(name string, x, y, w, h int, body string) *MenuWidget {
    return &MenuWidget{Widget{name, x, y, w, h, body, false}}
}
