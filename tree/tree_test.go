package tree

var testTree = &BinaryNode{
	1,
	nil,
	&BinaryNode{
		2,
		&BinaryNode{3,
			&BinaryNode{5,
				nil,
				nil},
			&BinaryNode{6,
				nil,
				nil}},
		&BinaryNode{4,
			nil,
			nil,
		},
	},
}

var testTree2 = &BinaryNode{
	1,
	&BinaryNode{
		2,
		&BinaryNode{
			4,
			nil,
			nil,
		},
		&BinaryNode{
			5,
			nil,
			nil,
		},
	},
	&BinaryNode{
		3,
		&BinaryNode{
			6,
			nil,
			nil,
		},
		&BinaryNode{
			7,
			nil,
			nil,
		},
	},
}
