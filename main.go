package main

//go:generate go-bindata -o data.go -prefix "data/" data

import (
	"bytes"
	"image/color"
	"image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	blindGopher = assetPictureData("blind-gopher-brown-beard.png")
	pinkyGopher = assetPictureData("blind-gopher-pinky-liscious.png")
	rightEye    = assetPictureData("right-eye.png")
	leftEye     = assetPictureData("left-eye.png")

	blindGopherSprite = pixel.NewSprite(blindGopher, blindGopher.Bounds())
	pinkyGopherSprite = pixel.NewSprite(pinkyGopher, pinkyGopher.Bounds())
	rightEyeSprite    = pixel.NewSprite(rightEye, rightEye.Bounds())
	leftEyeSprite     = pixel.NewSprite(leftEye, leftEye.Bounds())
)

func main() {
	pixelgl.Run(run)
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, float64(1280), float64(720)),
		VSync:       true,
		Resizable:   false,
		Undecorated: true,
	})
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		win.Clear(color.RGBA{0x87, 0xce, 0xeb, 0xff})

		if win.Pressed(pixelgl.MouseButtonLeft) {
			pinkyGopherSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		} else {
			blindGopherSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		}

		rightEyeSprite.Draw(win, pixel.IM.Moved(rightEyePosition(win.MousePosition())))
		leftEyeSprite.Draw(win, pixel.IM.Moved(leftEyePosition(win.MousePosition())))

		win.Update()
	}
}

func rightEyePosition(mouse pixel.Vec) pixel.Vec {
	p := pixel.Vec{450, 338}

	if mouse.Y > 405 {
		p.Y = 405
	} else if mouse.Y < 312 {
		p.Y = 312
	} else {
		p.Y = mouse.Y
	}

	if mouse.X > 250 && mouse.X < 470 {
		p.X = mouse.X
	} else if mouse.X < 250 {
		p.X = 250
	} else {
		p.X = 470
	}

	return p
}

func leftEyePosition(mouse pixel.Vec) pixel.Vec {
	p := pixel.Vec{830, 338}

	if mouse.Y > 410 {
		p.Y = 410
	} else if mouse.Y < 300 {
		p.Y = 300
	} else {
		p.Y = mouse.Y
	}

	if mouse.X > 834 && mouse.X < 1060 {
		p.X = mouse.X
	} else if mouse.X < 834 {
		p.X = 834
	} else {
		p.X = 1060
	}

	return p
}

func assetPictureData(name string) *pixel.PictureData {
	data, err := Asset(name)
	if err != nil {
		panic(err)
	}

	m, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return pixel.PictureDataFromImage(m)
}
