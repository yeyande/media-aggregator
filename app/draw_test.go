package app

import (
    "testing"
    "github.com/nsf/termbox-go"
)

func TestGetConnector(t *testing.T) {
    cases := []struct {
        arg1 []termbox.Cell
        arg2 Window
        want rune
    }{
        {[]termbox.Cell{}, Window{}, 0x00},
    }
    for _, c := range cases {
        got := getConnector(c.arg1, c.arg2)
        if got != c.want {
            t.Errorf("getConnector(%q, %q) == %q, want %q", c.arg1, c.arg2, got, c.want)
        }
    }
}
