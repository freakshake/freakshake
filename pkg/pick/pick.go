package pick

import "github.com/freakshake/type/optional"

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

func IfSome[T any](oldData, newData optional.Optional[T]) optional.Optional[T] {
	if newData.IsSome() {
		return newData
	}
	return oldData
}
