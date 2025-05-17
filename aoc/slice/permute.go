package slice

func Permute[T any](data []T, yield func() bool) {
	if len(data) == 0 {
		return
	}

	var gen func(int) bool
	gen = func(idx int) bool {
		if idx == 0 {
			return yield()
		}

		if !gen(idx - 1) {
			return false
		}

		for i := 0; i < idx; i++ {
			if idx%2 == 1 {
				data[idx], data[i] = data[i], data[idx]
			} else {
				data[0], data[idx] = data[idx], data[0]
			}

			if !gen(idx - 1) {
				return false
			}
		}

		return true
	}

	gen(len(data) - 1)
}
