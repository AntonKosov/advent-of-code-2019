package math

import "iter"

var primesCache []uint64

func init() {
	primesCache = []uint64{2}
}

func PrimeNumbers() iter.Seq[uint64] {
	findNext := func() {
	nextNum:
		for n := primesCache[len(primesCache)-1] + 1; ; n++ {
			for _, pn := range primesCache {
				if pn*pn > n {
					break
				}
				if n%pn == 0 {
					continue nextNum
				}
			}
			primesCache = append(primesCache, n)
			return
		}
	}

	return func(yield func(prime uint64) bool) {
		for i := 0; ; i++ {
			if i+1 > len(primesCache) {
				findNext()
			}

			if !yield(primesCache[i]) {
				return
			}
		}
	}
}
