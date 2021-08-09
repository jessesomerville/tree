package node

type Node struct {
	Value    int
	Children []*Node
	Thread   *Node
	X, Y     int
	Mod      int
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) AddChildren(children ...*Node) {
	n.Children = append(n.Children, children...)
}

func (n *Node) GetHeight() int {
	height := 0
	for _, c := range n.Children {
		if h := c.GetHeight(); h > height {
			height = h
		}
	}
	return height + 1
}
