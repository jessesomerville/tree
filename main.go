package main

import (
	"github.com/jessesomerville/tree/pkg/drawtree"
	"github.com/jessesomerville/tree/pkg/node"
)

func main() {
	t := node.NewNode(1)
	t.PopulateBigTree()
	drawtree.ReingoldTilford(t)
	drawtree.Draw(t, "test.png")
}
