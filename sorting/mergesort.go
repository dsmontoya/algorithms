package sorting

func merge(a []byte, lo, mid, hi int) {
	aux := make([]byte, len(a))
	i := lo
	j := mid + 1
	copy(aux, a)
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if less(aux[j], aux[i]) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func topDownMergesort(a []byte, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	topDownMergesort(a, lo, mid)
	topDownMergesort(a, mid+1, hi)
	merge(a, lo, mid, hi)
}
