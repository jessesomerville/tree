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

	var lNode, rNode, lPrev, rPrev, lPrevPar, rPrevPar *node.Node
	var lMore, rMore bool
	lOffset, rOffset := left.Mod, right.Mod
	for {
		lNode, lMore = <-lChan
		rNode, rMore = <-rChan

		if !lMore && !rMore {
			// Max depth of both subtrees has been reached
			break
		} else if !lMore {
			if lPrevPar != nil && len(lPrevPar.Children) == 2 {
				// The last node in the left subtree had a sibling leaf so thread it instead
				lPrevPar.Children[0].Thread = rNode
			} else {
				// Thread the last node in the left subtree
				lPrev.Thread = rNode
			}
			break
		} else if !rMore {
			if rPrevPar != nil && len(rPrevPar.Children) == 2 {
				// The last node in the right subtree had a sibling leaf so thread it instead
				rPrevPar.Children[1].Thread = lNode
			} else {
				// Thread the last node in the right subtree
				rPrev.Thread = lNode
			}
			break
		}

		lPos := lNode.X + lOffset
		rPos := rNode.X + rOffset
		if dist := lPos - rPos; dist > distToMove {
			distToMove = dist
		}

		// Only add offsets if this node isn't a leaf
		if lNode.Thread == nil {
			lOffset += lNode.Mod
		}
		if rNode.Thread == nil {
			rOffset += rNode.Mod
		}

		lPrevPar, rPrevPar = lPrev, rPrev
		lPrev, rPrev = lNode, rNode
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
	// fmt.Printf("%d: %d\n", n.Value, n.Mod+currMod)
	n.X += currMod
	for _, child := range n.Children {
		addMod(child, currMod+n.Mod)
	}
}
