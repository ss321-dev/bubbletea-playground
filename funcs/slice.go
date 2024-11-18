package funcs

func Remove[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}
