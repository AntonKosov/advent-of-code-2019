package math

type Numbers interface {
	int8 | int | int64 | uint64
}

func Abs[T Numbers](v T) T {
	if v < 0 {
		return -v
	}

	return v
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Sign[T Numbers](v T) int {
	if v == 0 {
		return 0
	}

	if v < 0 {
		return -1
	}

	return 1
}

func GCD[T Numbers](nums ...T) T {
	if len(nums) < 2 {
		panic("incorrect number of parameters")
	}

	a, b := nums[0], nums[1]
	if len(nums) > 2 {
		b = GCD(nums[1:]...)
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM[T Numbers](nums ...T) T {
	if len(nums) < 2 {
		panic("incorrect number of parameters")
	}

	a, b := nums[0], nums[1]
	if len(nums) > 2 {
		b = LCM(nums[1:]...)
	}

	return a * b / GCD(a, b)
}
