package util

func Pointer[T any](v T) *T {
	return &v
}

func SliceValueToSlicePointer[T any](sliceValue []T) []*T {
	slicePointer := []*T{}
	for i := range sliceValue {
		slicePointer = append(slicePointer, &sliceValue[i])
	}

	return slicePointer
}

func SlicePointerToSliceValue[T any](sliceValue []*T) []T {
	slicePointer := []T{}
	for i := range sliceValue {
		slicePointer = append(slicePointer, *sliceValue[i])
	}

	return slicePointer
}

func ConvertArray[K any, T any](arr []K, callback func(K) T) []T {
	nodes := []T{}

	for _, v := range arr {
		nodes = append(nodes, callback(v))
	}

	return nodes
}
