package views

import (
	"parking/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MainView struct{}

func NewMainView() *MainView{
	return &MainView{}
}
func (v *MainView) Run(){
	app := app.New()
	window := app.NewWindow("Parking")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(1000, 1000))

	mainScene := scenes.NewMainScene(window)
	mainScene.Show()
	go mainScene.Run()
	window.ShowAndRun()
}
