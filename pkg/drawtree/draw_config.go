package drawtree

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/goregular"
)

// Font faces
const (
	Regular = iota
	Bold
	Italic
	BoldItalic
)

var fontMap = map[int][]byte{
	Regular:    goregular.TTF,
	Bold:       gobold.TTF,
	Italic:     goitalic.TTF,
	BoldItalic: gobolditalic.TTF,
}

// keys are 'face-size'
var fontCache = map[string]font.Face{}

func loadFont(face int, size float64) font.Face {
	key := fmt.Sprintf("%d-%f", face, size)
	if f, ok := fontCache[key]; ok {
		return f
	}

	ttf, err := truetype.Parse(fontMap[face])
	if err != nil {
		log.Fatalf("failed to load font: %v", err)
	}
	ff := truetype.NewFace(ttf, &truetype.Options{Size: size})
	fontCache[key] = ff
	return ff
}

type DrawConfig struct {
	BgColor  color.Color
	FgColor  color.Color
	NodeSize float64
	Stroke   float64
	FontSize float64
	Padding  int
}

func NewConfig(bg, fg color.Color, ns, stroke, fs float64, pad int) *DrawConfig {
	return &DrawConfig{
		BgColor:  bg,
		FgColor:  fg,
		NodeSize: ns,
		Stroke:   stroke,
		FontSize: fs,
		Padding:  pad,
	}
}

func DefaultConfig() *DrawConfig {
	return &DrawConfig{
		BgColor:  image.White,
		FgColor:  image.Black,
		NodeSize: 35.0,
		Stroke:   2.0,
		FontSize: 35.0 * 0.75,
		Padding:  35.0 * 3,
	}
}

func DefaultConfigDark() *DrawConfig {
	return &DrawConfig{
		BgColor:  color.RGBA{0x2a, 0x2a, 0x2a, 0xff},
		FgColor:  colornames.Darkgrey,
		NodeSize: 35.0,
		Stroke:   2.0,
		FontSize: 35.0 * 0.75,
		Padding:  35.0 * 3,
	}
}
