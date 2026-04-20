package sort

// 泛型  任何类型的slice
type byLenth[T any] [][]T

func (b byLenth[T]) Len() int {
	return len(b)
}

func (b byLenth[T]) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func (b byLenth[T]) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
