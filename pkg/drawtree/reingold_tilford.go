package drawtree

import (
	"github.com/jessesomerville/tree/pkg/node"
)

func ReingoldTilford(root *node.Node) {
	if len(root.Children) == 0 {
		return
	}
	rtBuild(root, 0)
	addMod(root, 0)
}

func rtBuild(n *node.Node, depth int) *node.Node {
	n.Y = depth

	if len(n.Children) == 0 {
		n.X = 0
		return n
	}

	if len(n.Children) == 1 {
		child := n.Children[0]
		n.X = rtBuild(child, depth+1).X
		return n
	}

	leftChild := n.Children[0]
	rightChild := n.Children[1]
	lSubtree := rtBuild(leftChild, depth+1)
	rSubtree := rtBuild(rightChild, depth+1)

	n.X = adjustSubtrees(lSubtree, rSubtree)
	return n
}

func adjustSubtrees(left, right *node.Node) int {
	distToMove := 0
	lChan := rightContour(left)
	rChan := leftContour(right)

	var lNode, rNode, lPrev, rPrev *node.Node
	var lMore, rMore bool
	lOffset, rOffset := left.Mod, right.Mod
	for {
		lNode, lMore = <-lChan
		rNode, rMore = <-rChan
		if !lMore && !rMore {
			break
		} else if !lMore {
			lPrev.Thread = rNode
			break
		} else if !rMore {
			rPrev.Thread = lNode
			break
		}
		lPos := lNode.X + lOffset
		rPos := rNode.X + rOffset
		if dist := lPos - rPos; dist > distToMove {
			distToMove = dist
		}
		lPrev, rPrev = lNode, rNode
		lOffset += lNode.Mod
		rOffset += rNode.Mod
	}
	distToMove += 1
	distToMove += (right.X + left.X + distToMove) % 2
	right.Mod = distToMove
	right.X += distToMove

	return (right.X + left.X) / 2
}

func leftContour(n *node.Node) <-chan *node.Node {
	c := make(chan *node.Node)
	go findLeftContour(n, c)
	return c
}

func findLeftContour(n *node.Node, c chan<- *node.Node) {
	c <- n
	if next := nextLeft(n); next != nil {
		findLeftContour(next, c)
	} else {
		close(c)
	}
}

func nextLeft(n *node.Node) *node.Node {
	if len(n.Children) != 0 {
		return n.Children[0]
	}
	if n.Thread != nil {
		return n.Thread
	}
	return nil
}

func rightContour(n *node.Node) <-chan *node.Node {
	c := make(chan *node.Node)
	go findRightContour(n, c)
	return c
}

func findRightContour(n *node.Node, c chan<- *node.Node) {
	c <- n
	if next := nextRight(n); next != nil {
		findRightContour(next, c)
	} else {
		close(c)
	}
}

func nextRight(n *node.Node) *node.Node {
	if len(n.Children) == 1 {
		return n.Children[0]
	}
	if len(n.Children) == 2 {
		return n.Children[1]
	}
	if n.Thread != nil {
		return n.Thread
	}
	return nil
}

func addMod(n *node.Node, currMod int) {
	n.X += currMod
	for _, child := range n.Children {
		addMod(child, currMod+n.Mod)
	}
}
