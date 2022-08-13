package slicepkg

type Number interface {
	int | int64 | int32 | float64
}

func Sum[V Number](array []V) V {
	var result V
	result = 0

	for _, item := range array {
		result += item
	}

	return result

}

func Average[V Number](array []V) float64 {
	sum := Sum(array)

	result := float64(sum) / float64(len(array))

	return result
}
