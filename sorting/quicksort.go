package sorting

import "fmt"

func partition(a []byte, lo, hi int) int {
	i := lo
	j := hi + 1
	v := a[lo]
	for {
		for {
			i++
			if !less(a[i], v) {
				break
			}
			if i == hi {
				break
			}
		}
		for {
			j--
			if !less(v, a[j]) {
				break
			}
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		// fmt.Println("i", i)
		// fmt.Println("a[i]", string(a[i]))
		// fmt.Println("j", j)
		// fmt.Println("a[j]", string(a[j]))
		exchange(a, i, j)
		// fmt.Println("after exchange", string(a))
	}
	// fmt.Println("final j", j)
	// fmt.Println("final a[j]", string(a[j]))
	// fmt.Println("before final exchange", string(a))
	exchange(a, lo, j)
	// fmt.Println("after final exchange", string(a))
	return j
}

func quicksort(a []byte, lo, hi int) {
	fmt.Println("complete:", string(a))
	fmt.Println("subarray:", string(a[lo:hi+1]))
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	fmt.Println("complete after partition", string(a))
	fmt.Println("subarray after partition", string(a[lo:hi+1]))
	fmt.Println("==")
	fmt.Println("LEFT")
	quicksort(a, lo, j-1)
	fmt.Println("RIGHT")
	quicksort(a, j+1, hi)
}
