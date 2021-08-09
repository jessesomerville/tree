package drawtree

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/jessesomerville/tree/pkg/node"
)

const (
	nodeSize = 35.0
	padding  = nodeSize * 3
	stroke   = 2.0
	fontSize = nodeSize * 0.75
)

var (
	fontFace font.Face
)

func init() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	fontFace = truetype.NewFace(font, &truetype.Options{Size: fontSize})
}

func Draw(n *node.Node, filename string) {
	maxX, maxY := scaleNodePositions(n)
	w, h := int(maxX+padding), int(maxY+padding)
	dc := getCanvas(n, w, h, image.White)
	drawTree(dc, n)
	dc.SavePNG(filename)
}

func getCanvas(t *node.Node, w, h int, bgColor color.Color) *gg.Context {
	canvas := gg.NewContext(w, h)
	canvas.SetColor(bgColor)
	canvas.Clear()
	return canvas
}

func drawTree(dc *gg.Context, n *node.Node) {
	for _, child := range n.Children {
		drawEdge(dc, n, child)
		drawTree(dc, child)
	}
	drawNode(dc, n)
}

func scaleNodePositions(n *node.Node) (int, int) {
	n.X = (n.X + 1) * padding
	n.Y = (n.Y + 1) * padding
	maxX, maxY := n.X, n.Y

	for _, child := range n.Children {
		x, y := scaleNodePositions(child)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	return maxX, maxY
}

func drawNode(dc *gg.Context, n *node.Node) {
	drawNodeCircle(dc, n.X, n.Y)

	dc.SetFontFace(fontFace)
	dc.SetColor(image.Black)
	text := fmt.Sprintf("%d", n.Value)
	dc.DrawStringAnchored(text, float64(n.X), float64(n.Y), 0.5, 0.5)
}

func drawNodeCircle(dc *gg.Context, x, y int) {
	dc.DrawCircle(float64(x), float64(y), nodeSize)

	dc.SetLineWidth(stroke)
	dc.SetColor(image.Black)
	dc.StrokePreserve()

	dc.SetColor(image.White)
	dc.Fill()
}

func drawEdge(dc *gg.Context, n1, n2 *node.Node) {
	dc.SetColor(image.Black)
	dc.DrawLine(float64(n1.X), float64(n1.Y), float64(n2.X), float64(n2.Y))
	dc.SetLineWidth(stroke)
	dc.Stroke()
}
