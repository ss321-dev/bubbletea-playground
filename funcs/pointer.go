package funcs

func ToValue[T any](ptr *T) T {
	if ptr == nil {
		return *new(T)
	}
	return *ptr
}

func ToPtr[T any](value T) *T {
	return &value
}
