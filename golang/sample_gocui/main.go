package main

import (
	"fmt"
	"github.com/pkg/profile"

	"github.com/jroimartin/gocui"
)

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer g.Close()

	g.SetManagerFunc(layout)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
		return
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		fmt.Println(err)
		return
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

