package main

import (
	"github.com/jessesomerville/tree/pkg/drawtree"
	"github.com/jessesomerville/tree/pkg/node"
)

func main() {
	t := node.NewNode(1)
	t.PopulateBinaryTreeOverlap()
	drawtree.ReingoldTilford(t)

	drawtree.Draw(t, "test.png")
}