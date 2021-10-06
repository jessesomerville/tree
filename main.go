package main

import (
	"github.com/jessesomerville/tree/pkg/drawtree"
	"github.com/jessesomerville/tree/pkg/node"
)

func main() {
	tree := node.ParseYAML("test_trees/binary_tree_small_3.yaml")

	drawtree.Draw(
		tree,
		drawtree.ReingoldTilford,
		"binary_tree_small_4.png",
		drawtree.DefaultConfigDark(),
	)
}
