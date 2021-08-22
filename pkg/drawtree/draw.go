package drawtree

import (
	"fmt"

	"github.com/fogleman/gg"

	"github.com/jessesomerville/tree/pkg/node"
)

func Draw(n *node.Node, algo func(*node.Node), filename string, cfg *DrawConfig) {
	algo(n)

	if cfg == nil {
		cfg = DefaultConfig()
	}

	w, h := scaleNodePositions(n, cfg.Padding)
	dc := gg.NewContext(w+cfg.Padding, h+cfg.Padding)
	dc.SetColor(cfg.BgColor)
	dc.Clear()

	drawTree(dc, n, cfg)
	dc.SavePNG(filename)
}

func scaleNodePositions(n *node.Node, padding int) (int, int) {
	n.X = (n.X + 1) * padding
	n.Y = (n.Y + 1) * padding
	maxX, maxY := n.X, n.Y

	for _, child := range n.Children {
		x, y := scaleNodePositions(child, padding)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	return maxX, maxY
}

func drawTree(dc *gg.Context, n *node.Node, cfg *DrawConfig) {
	// if tNode := n.Thread; tNode != nil {
	// 	drawThread(dc, n, tNode, cfg)
	// }
	for _, child := range n.Children {
		drawEdge(dc, n, child, cfg)
		drawTree(dc, child, cfg)
	}
	drawNode(dc, n, cfg)
}

func drawNode(dc *gg.Context, n *node.Node, cfg *DrawConfig) {
	drawNodeCircle(dc, n.X, n.Y, cfg)

	fontFace := loadFont(n.Font, cfg.FontSize)
	dc.SetFontFace(fontFace)
	dc.SetColor(cfg.FgColor)
	text := fmt.Sprintf("%d", n.Value)
	dc.DrawStringAnchored(text, float64(n.X), float64(n.Y), 0.5, 0.5)
}

func drawNodeCircle(dc *gg.Context, x, y int, cfg *DrawConfig) {
	dc.DrawCircle(float64(x), float64(y), cfg.NodeSize)

	dc.SetLineWidth(cfg.Stroke)
	dc.SetColor(cfg.FgColor)
	dc.StrokePreserve()

	dc.SetColor(cfg.BgColor)
	dc.Fill()
}

func drawEdge(dc *gg.Context, n1, n2 *node.Node, cfg *DrawConfig) {
	dc.SetColor(cfg.FgColor)
	dc.DrawLine(float64(n1.X), float64(n1.Y), float64(n2.X), float64(n2.Y))
	dc.SetLineWidth(cfg.Stroke)
	dc.Stroke()
}

func drawThread(dc *gg.Context, n1, n2 *node.Node, cfg *DrawConfig) {
	dc.SetColor(cfg.FgColor)
	dc.DrawLine(float64(n1.X), float64(n1.Y), float64(n2.X), float64(n2.Y))
	dc.SetDash(16, 24)
	dc.SetLineWidth(cfg.Stroke)
	dc.Stroke()
	dc.SetDash()
}
