package node

import (
	"math/rand"
	"time"
)

func GenerateRandTree(numOfNodes int) *Node {
	randSeq := generateRandSeq(numOfNodes)
	return treeFromRandSeq(randSeq, 1)
}

func treeFromRandSeq(seq []int, value int) *Node {
	var stack []*Node
	var root, currNode *Node
	insertRight := false

	for _, num := range seq {
		if num == 1 {
			prevNode := currNode
			currNode = NewNode(value)
			value++

			if root == nil {
				root = currNode
			} else if insertRight {
				prevNode.AddChildren(currNode)
				insertRight = false
			} else {
				prevNode.AddChildren(currNode)
			}

			stack = append(stack, currNode)
		} else {
			n := len(stack) - 1
			currNode = stack[n]
			stack = stack[:n]
			insertRight = true
		}
	}
	return root
}

func generateRandSeq(n int) []int {
	seq := make([]int, n*2)
	for i := range seq {
		if i%2 == 0 {
			seq[i] = -1
		} else {
			seq[i] = 1
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(seq), func(i, j int) {
		seq[i], seq[j] = seq[j], seq[i]
	})
	return balanceBrackets(seq)
}

func balanceBrackets(seq []int) []int {
	var prefix, suffix, word []int
	var partialSum int

	for _, s := range seq {
		word = append(word, s)
		partialSum += s
		if partialSum == 0 {
			if s == -1 {
				prefix = append(prefix, word...)
			} else {
				prefix = append(prefix, 1)
				tmpSuf := []int{-1}
				for _, x := range word[1 : len(word)-1] {
					tmpSuf = append(tmpSuf, -x)
				}
				suffix = append(tmpSuf, suffix...)
			}
			word = []int{}
		}
	}
	return append(prefix, suffix...)
}
