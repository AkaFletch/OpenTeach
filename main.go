package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow(app.Title("OpenTeach"), app.Size(unit.Dp(800), unit.Dp(650)))
		uiLoop(w)
	}()
	app.Main()
}

func uiLoop(w *app.Window) error {
	gofont.Register()
	th := material.NewTheme()
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e.Queue, e.Config, e.Size)
			label := material.H1(th, "Hello!")
			layout.Rigid(label.Layout)
			label.Layout(gtx)
			e.Frame(&ops)
		}
	}
}
