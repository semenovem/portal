package util

func NumMapToArr[T uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | string](a map[T]struct{}) []T {
	b := make([]T, 0, len(a))
	for k := range a {
		b = append(b, k)
	}

	return b
}
