package utils

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Height int
	Width  int
	Length int
	ColorR uint8
	ColorG uint8
	ColorB uint8
	ColorA uint8
}

func (c *Captcha) DriverString() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          c.Height,
		Width:           c.Width,
		NoiseCount:      0,
		ShowLineOptions: 0,
		Length:          c.Length,
		Source:          "0123456789",
		BgColor: &color.RGBA{
			R: c.ColorR,
			G: c.ColorG,
			B: c.ColorB,
			A: c.ColorA,
		},
		Fonts: nil,
	}

	return stringType
}

func (c *Captcha) DriverMath() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          c.Height,
		Width:           c.Width,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor: &color.RGBA{
			R: c.ColorR,
			G: c.ColorG,
			B: c.ColorB,
			A: c.ColorA,
		},
		Fonts: nil,
	}

	return mathType
}
