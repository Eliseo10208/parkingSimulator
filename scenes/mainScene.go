package scenes

import (
	"parking/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"sync"
	"time"
)
type MainScene struct{
	window fyne.Window
}
func NewMainScene(w fyne.Window) *MainScene{
	return &MainScene{window: w}
}
var carsContainer  = container.NewWithoutLayout()

func (s *MainScene) Show() {
	rectangle := canvas.NewRectangle(color.Transparent)
	rectangle.StrokeWidth = 2
	rectangle.StrokeColor = color.White
	rectangle.Resize(fyne.NewSize(400, 620))
	rectangle.Move(fyne.NewPos(200, 10))

	gate := canvas.NewRectangle(color.White)
	gate.Resize(fyne.NewSize(10, 100))
	gate.Move(fyne.NewPos(195, 300))

	carsContainer.Add(rectangle)
	carsContainer.Add(gate)
	s.window.SetContent(carsContainer)
}
func (s *MainScene) Run() {
	p := models.NewParking(make(chan int, 20), &sync.Mutex{})
	poissonDist := models.NewPoissonDist()

	var wg sync.WaitGroup

	carCountText := canvas.NewText("Autos en el estacionamiento: 0", color.White)
	carCountText.Move(fyne.NewPos(300, 10))
	carsContainer.Add(carCountText)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			car := models.NewCar(id)
			carImage := car.GetCarImage()
			carImage.Resize(fyne.NewSize(50, 30))
			carImage.Move(fyne.NewPos(-20, 310))

			carsContainer.Add(carImage)
			carsContainer.Refresh()

			car.Park(p, carsContainer, carCountText, &wg)
		}(i)
		var randPoissonNumber = poissonDist.Generate(float64(2))
		time.Sleep(time.Second * time.Duration(randPoissonNumber))
	}

	wg.Wait()
	fmt.Println("terminado")
}
