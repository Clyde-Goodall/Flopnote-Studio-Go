package main

import "image/color"

type ThemeColor int

const (
	LightGrey = iota
	OffWhite
	Orange
	LightOrange
	Yellow
)

func (t ThemeColor) ColorPalette() color.Color {
	return [...]color.Color{
		color.Color(color.RGBA{238, 238, 238, 200}),
		color.Color(color.RGBA{255, 255, 242, 255}),
		color.Color(color.RGBA{242, 123, 45, 255}),
		color.Color(color.RGBA{242, 238, 226, 255}),
		color.Color(color.RGBA{255, 253, 123, 255}),
	}[t-1]
}
