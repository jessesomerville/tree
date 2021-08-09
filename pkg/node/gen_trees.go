package node

func (root *Node) Populate() {
	var (
		two   = NewNode(2)
		three = NewNode(3)
		four  = NewNode(4)
		five  = NewNode(5)
		six   = NewNode(6)
		seven = NewNode(7)
		eight = NewNode(8)
		nine  = NewNode(9)
	)
	root.AddChildren(two, three, four)
	two.AddChildren(five, six)
	four.AddChildren(seven)
	six.AddChildren(eight)
	six.AddChildren(nine)
}

func (root *Node) Populate2() {
	var (
		two      = NewNode(2)
		three    = NewNode(3)
		four     = NewNode(4)
		five     = NewNode(5)
		six      = NewNode(6)
		seven    = NewNode(7)
		eight    = NewNode(8)
		nine     = NewNode(9)
		ten      = NewNode(10)
		eleven   = NewNode(11)
		twelve   = NewNode(12)
		thirteen = NewNode(13)
		fourteen = NewNode(14)
		fifteen  = NewNode(15)
	)
	root.AddChildren(two, three)
	two.AddChildren(four, five)
	four.AddChildren(six, seven)
	six.AddChildren(eight, nine)
	eight.AddChildren(ten, eleven)
	three.AddChildren(twelve, thirteen)
	twelve.AddChildren(fourteen, fifteen)
}

func (root *Node) PopulateBinaryTree() {
	var (
		two   = NewNode(2)
		three = NewNode(3)
		four  = NewNode(4)
		five  = NewNode(5)
		six   = NewNode(6)
		seven = NewNode(7)
		eight = NewNode(8)
		nine  = NewNode(9)
		ten   = NewNode(10)
	)
	root.AddChildren(two, three)
	two.AddChildren(four)
	three.AddChildren(five, six)
	four.AddChildren(seven, eight)
	six.AddChildren(nine, ten)
}

func (root *Node) PopulateBinaryTreeOverlap() {
	var (
		two   = NewNode(2)
		three = NewNode(3)
		four  = NewNode(4)
		five  = NewNode(5)
		six   = NewNode(6)
		seven = NewNode(7)
		eight = NewNode(8)
		nine  = NewNode(9)
	)
	root.AddChildren(two, three)
	two.AddChildren(four)
	three.AddChildren(five)
	four.AddChildren(six, seven)
	five.AddChildren(eight, nine)
}
