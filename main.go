package main

import (
	"github.com/jessesomerville/tree/pkg/drawtree"
	"github.com/jessesomerville/tree/pkg/node"
)

var treeFiles = []string{
	"test_trees/binary_tree_basic.yaml",
	"test_trees/binary_tree_overlap.yaml",
	"test_trees/binary_tree_small_1.yaml",
	"test_trees/binary_tree_small_2.yaml",
	"test_trees/binary_tree_small_3.yaml",
}

func main() {
	tree := node.ParseYAML("test_trees/binary_tree_small_3.yaml")

	drawtree.Draw(
		tree,
		drawtree.ReingoldTilford,
		"binary_tree_small_4.png",
		drawtree.DefaultConfigDark(),
	)
}
