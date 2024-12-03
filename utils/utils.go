package utils

func Filter[T any](xs []T, test func(T) bool) (ys []T) {
	for _, x := range xs {
		if test(x) {
			ys = append(ys, x)
		}
	}
	return
}

func ElemCount[T any](xs []T, test func(T) bool) int {
	return len(Filter(xs, test))
}
