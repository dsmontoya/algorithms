package tree

type DepthCalculator struct {
	maximum uint
}

func (b *BinaryNode) MaximumDepthTopDown() uint {
	d := &DepthCalculator{}
	return d.MaximumDepth(b)
}

func (b *BinaryNode) MaximumDepthBottomUp() uint {
	if b == nil {
		return 0
	}
	left := b.Left.MaximumDepthBottomUp()
	right := b.Right.MaximumDepthBottomUp()
	if left > right {
		return left + 1
	}
	return right + 1
}

func (d *DepthCalculator) MaximumDepth(node *BinaryNode) uint {
	d.maximumDepth(node, 0)
	return d.maximum
}

func (d *DepthCalculator) maximumDepth(node *BinaryNode, depth uint) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if d.maximum < depth {
			d.maximum = depth
		}
	}
	d.maximumDepth(node.Left, depth+1)
	d.maximumDepth(node.Right, depth+1)
}
