package client

import (
	"fmt"
	"image/color"
	"time"

	"github.com/TheInvader360/hack-emulator/hack"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	gfxW = 512
	gfxH = 256
)

// Client - Struct
type Client struct {
	VM  *hack.Hack
	fps int
}

// Run - https://pkg.go.dev/github.com/faiface/pixel/pixelgl#Run
func (c *Client) Run() {
	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, gfxW, gfxH),
		VSync:  true,
	}

	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	canvas := pixelgl.NewCanvas(pixel.R(0, 0, gfxW, gfxH))

	lastFrame := time.Now()

	for !window.Closed() {
		if window.Pressed(pixelgl.KeyLeftControl) {
			if window.Pressed(pixelgl.KeyQ) {
				return
			}
			if window.Pressed(pixelgl.KeyR) {
				(*c.VM).Reset()
			}
		}
		c.handleInput(window)
		c.handleOutput(window, canvas)
		dt := time.Since(lastFrame).Seconds()
		c.fps = int(1 / dt)
		lastFrame = time.Now()
		window.SetTitle(fmt.Sprintf("Hack Emulator | %d fps", c.fps))
	}
}

func (c *Client) handleInput(window *pixelgl.Window) {
	keyCode := 0
	for i, key := range Keys {
		if window.Pressed(key) {
			keyCode = i
			break
		}
	}
	(*c.VM).SetKeyboard(int16(keyCode))
}

func (c *Client) handleOutput(window *pixelgl.Window, canvas *pixelgl.Canvas) {
	pixels := []uint8{}
	for _, word := range (*c.VM).GetScreen() {
		for i := 15; i >= 0; i-- {
			pixelRGB := uint8(255) // white
			if (word>>i)&0b1 == 1 {
				pixelRGB = uint8(0) // black
			}
			pixels = append(pixels, pixelRGB)
			pixels = append(pixels, pixelRGB)
			pixels = append(pixels, pixelRGB)
			pixels = append(pixels, 255)
		}
	}
	canvas.Clear(color.White)
	canvas.SetPixels(pixels)
	canvas.Draw(window, pixel.IM.Moved(pixel.V(gfxW/2, gfxH/2)).ScaledXY(pixel.V(gfxW/2, gfxH/2), pixel.V(1, -1)))
	window.Update()
}
