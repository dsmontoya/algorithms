package sorting

import (
	"math"
)

func bottomUpMergesort(a []byte, lo, hi int) {
	n := len(a)
	for len := 1; len < n; len *= 2 {
		for lo := 0; lo < n-len; lo += len + len {
			mid := lo + len - 1
			hi := int(math.Min(float64(lo+len+len-1), float64(n-1)))
			merge(a, lo, mid, hi)
		}
	}
}

func exchange(a []byte, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

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
