package update

func IfChanged[T comparable](oldData, newData T) (empty T) {
	if newData != empty {
		return newData
	}
	return oldData
}

func IfNilChanged[T comparable](oldData, newData *T) *T {
	var empty T
	if newData != nil && *newData != empty {
		return newData
	}
	return oldData
}
