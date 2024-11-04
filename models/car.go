package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"math/rand"
	"sync"
	"time"
)

type Car struct {
	id int
	parkingTime time.Duration
	image *canvas.Image
	space int 
	exitImage *canvas.Image
}

func NewCar(id int) *Car {
	image := canvas.NewImageFromURI(storage.NewFileURI("./assets/lancer.png"))
	exitImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/lancerExit.png"))
	return &Car{
		id:id,
		parkingTime: time.Duration(rand.Intn(10)+10) * time.Second,
		image: image,
		space: 0,
		exitImage: exitImage,
	}
}
func (c *Car) Enter(p *Parking, carsContainer *fyne.Container, carCountText *canvas.Text) {
	p.GetSpaces() <- c.GetId()
	p.GetEntrance().Lock()

	spaceArray := p.GetSpacesArray()
	fmt.Printf("Auto %d ha entrado. Espacios ocupados: %d\n",c.GetId(), len(p.GetSpaces()))
	for i:=0; i<5; i++{
		c.image.Move(fyne.NewPos(c.image.Position().X+20, c.image.Position().Y))
		
		time.Sleep(200 * time.Millisecond)
	}
	p.GetEntrance().Unlock()
	for i:=0; i <len(spaceArray); i++{
		if spaceArray[i]== false {
			spaceArray[i] = true
			c.space = i
			c.image.Move(fyne.NewPos(290,float32(290+(i*30))))
			
			break
		}
	}
	p.SetSpacesArray(spaceArray)
	carsContainer.Refresh()

	// Actualizar el texto con el número actual de autos
	carCountText.Text = fmt.Sprintf("Autos en el estacionamiento: %d", p.GetCurrentCarCount())
	carCountText.Refresh()
}

func (c *Car) Leave(p *Parking, carsContainer *fyne.Container, carCountText *canvas.Text) {
	p.GetEntrance().Lock()
	<- p.GetSpaces()

	spaceArray := p.GetSpacesArray()
	spaceArray[c.space] = false
	p.SetSpacesArray(spaceArray)

	carsContainer.Refresh()
	fmt.Printf("Auto %d ha salido. Espacios ocupados: %d\n", c.GetId(), len(p.GetSpaces()))
	p.GetEntrance().Unlock()
	for i:=0; i< 10; i++{
		c.exitImage.Move(fyne.NewPos(c.exitImage.Position().X-30, c.exitImage.Position().Y))
		time.Sleep(200 * time.Millisecond)
	}
	carsContainer.Remove(c.exitImage)
	carsContainer.Refresh()

	// Actualizar el texto con el número actual de autos
	carCountText.Text = fmt.Sprintf("Autos en el estacionamiento: %d", p.GetCurrentCarCount())
	carCountText.Refresh()
}
func (c *Car) Park(p *Parking, carsContainer *fyne.Container, carCountText *canvas.Text, wg *sync.WaitGroup) {
	for i:=0; i <7; i++{
		c.exitImage.Move(fyne.NewPos(c.exitImage.Position().X-30, c.exitImage.Position().Y))
	
		time.Sleep(time.Millisecond *200)
	}
	c.Enter(p,carsContainer, carCountText)
	time.Sleep(c.parkingTime)
	carsContainer.Remove(c.image)
	c.exitImage.Resize(fyne.NewSize(50,30))
	p.ExitQueue(carsContainer, c.exitImage)
	c.Leave(p,carsContainer, carCountText)
	wg.Done()
}
func (c *Car) GetId() int{
	return c.id
}

func (c *Car) GetCarImage() *canvas.Image{
	return c.image
}
