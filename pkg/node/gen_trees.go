package node

func (root *Node) PopulateBasic() {
	root.AddChildren(NewNode(2), NewNode(3))
}

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

func (n *Node) setPos(x, y int) {
	n.X = x
	n.Y = y
}

func (root *Node) PopulateBinaryTree2() {
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
	)
	root.AddChildren(two, three)
	two.AddChildren(four, five)
	three.AddChildren(six, seven)
	six.AddChildren(eight, nine)
	seven.AddChildren(ten, eleven)
	eight.AddChildren(twelve)
	ten.AddChildren(thirteen)

	// root.setPos(3, 0)
	// two.setPos(1, 1)
	// three.setPos(5, 1)
	// four.setPos(0, 2)
	// five.setPos(2, 2)
	// six.setPos(3, 2)
	// seven.setPos(7, 2)
	// eight.setPos(2, 3)
	// nine.setPos(4, 3)
	// ten.setPos(6, 3)
	// eleven.setPos(8, 3)
	// twelve.setPos(2, 4)
	// thirteen.setPos(6, 4)

	// four.Thread = twelve
	// eleven.Thread = thirteen
}

func (root *Node) PopulateBinaryTree3() {
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

func (root *Node) PopulateBigTree() {
	var (
		a = NewNode(0)
		b = NewNode(1)
		c = NewNode(2)
		d = NewNode(3)
		e = NewNode(4)
		f = NewNode(5)
		g = NewNode(6)
		h = NewNode(7)
		i = NewNode(8)
		j = NewNode(9)
		k = NewNode(10)
		l = NewNode(11)
		m = NewNode(12)
		n = NewNode(13)
		o = NewNode(14)
		p = NewNode(15)
		q = NewNode(16)
		r = NewNode(17)
		s = NewNode(18)
		t = NewNode(19)
		u = NewNode(20)
		v = NewNode(21)
		w = NewNode(22)
		x = NewNode(23)
		y = NewNode(24)
		z = NewNode(25)
	)
	root.AddChildren(a, b)
	a.AddChildren(c, d)
	b.AddChildren(e, f)
	e.AddChildren(g)
	f.AddChildren(h, i)
	g.AddChildren(j, k)
	h.AddChildren(l, m)
	l.AddChildren(n)
	m.AddChildren(o, p)
	n.AddChildren(q, r)
	o.AddChildren(s, t)
	p.AddChildren(u)
	q.AddChildren(v)
	u.AddChildren(w, x)
	t.AddChildren(y, z)
}
