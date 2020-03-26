package sorting

func exchange(a []byte, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func less(a, b byte) bool {
	return a < b
}
