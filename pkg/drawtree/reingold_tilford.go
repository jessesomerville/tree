package drawtree

import (
	"log"

	"github.com/jessesomerville/tree/pkg/node"
)

func ReingoldTilford(n *node.Node) {
	build(n, 0)
	addMod(n, 0)
}

func build(n *node.Node, depth int) *node.Node {
	if len(n.Children) > 2 {
		log.Fatal("tree must be a binary tree")
	}

	n.Y = depth
	if len(n.Children) == 0 {
		n.X = 0
		return n
	}

	if len(n.Children) == 1 {
		child := n.Children[0]
		n.X = build(child, depth+1).X
		return n
	}
	lChild, rChild := n.Children[0], n.Children[1]
	left := build(lChild, depth+1)
	right := build(rChild, depth+1)
	n.X = fixSubtrees(left, right)
	return n
}

func addMod(n *node.Node, mod int) {
	n.X += mod
	for _, child := range n.Children {
		addMod(child, mod+n.Mod)
	}
}

func nextRight(n *node.Node) *node.Node {
	if n.Thread != nil {
		return n.Thread
	}
	if len(n.Children) > 0 {
		return n.Children[len(n.Children)-1]
	}
	return nil
}

func nextLeft(n *node.Node) *node.Node {
	if n.Thread != nil {
		return n.Thread
	}
	if len(n.Children) > 0 {
		return n.Children[0]
	}
	return nil
}

func contour(left, right, leftOuter, rightOuter *node.Node, maxOffset, lOffset, rOffset int) (*node.Node, *node.Node, *node.Node, *node.Node, int, int, int) {
	delta := (left.X + lOffset) - (right.X + rOffset)
	if maxOffset == 0 || delta > maxOffset {
		maxOffset = delta
	}

	if leftOuter == nil {
		leftOuter = left
	}
	if rightOuter == nil {
		rightOuter = right
	}

	lo := nextLeft(leftOuter)
	li := nextRight(left)
	ri := nextLeft(right)
	ro := nextRight(rightOuter)

	if li != nil && ri != nil {
		lOffset += left.Mod
		rOffset += right.Mod
		return contour(li, ri, lo, ro, maxOffset, lOffset, rOffset)
	}
	return li, ri, leftOuter, rightOuter, maxOffset, lOffset, rOffset
}

func fixSubtrees(left, right *node.Node) int {
	li, ri, lo, ro, diff, lOffset, rOffset := contour(left, right, nil, nil, 0, 0, 0)
	diff += 1
	diff += (right.X + diff + left.X) % 2

	right.Mod = diff
	right.X += diff

	if len(right.Children) > 0 {
		rOffset += diff
	}

	if ri != nil && li == nil {
		lo.Thread = ri
		lo.Mod = rOffset - lOffset
	} else if ri == nil && li != nil {
		ro.Thread = li
		ro.Mod = lOffset - rOffset
	}

	return (left.X + right.X) / 2
}
