package node

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Node struct {
	Value    int     `yaml:"value"`
	Children []*Node `yaml:"children,omitempty"`
	Thread   *Node
	X, Y     int
	Mod      int
	Font     int
}

func NewNode(value int, children ...*Node) *Node {
	return &Node{Value: value, Children: children}
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

func ParseYAML(file string) *Node {
	fBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", file, err)
	}

	var root Node
	if err := yaml.Unmarshal(fBytes, &root); err != nil {
		log.Fatalf("failed to unmarshal tree: %v", err)
	}
	return &root
}
