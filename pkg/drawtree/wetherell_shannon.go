package drawtree

import (
	"log"

	"github.com/jessesomerville/tree/pkg/node"
)

func WetherellShannonMinWidth(n *node.Node, nextX []int, depth int) {
	if nextX == nil {
		nextX = make([]int, n.GetHeight())
	}
	if n.X != 0 || n.Y != 0 {
		log.Fatalf("node %d has more than one direct ancestor", n.Value)
	}
	n.X = nextX[depth]
	n.Y = depth
	nextX[depth]++
	for _, c := range n.Children {
		WetherellShannonMinWidth(c, nextX, depth+1)
	}
}
