package es

func ToAnySlice[T any](in []T) []interface{} {
	res := make([]interface{}, 0)
	for _, v := range in {
		res = append(res, v)
	}
	return res
}
