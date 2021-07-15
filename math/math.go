package math

func Add(x, y int) (res int) {
	return x + y
}

func Subtract(x, y int) (res int) {
	return x - y
}

func Multiply(x, y int64) (res int64) {
	return x * y
}

func Divide(x, y float64) (res float64) {
	if y != 0 {
		return x / y
	}

	return -1
}
