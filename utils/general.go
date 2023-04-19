package utils

func ConvertToSliceInterface(data ...string) []interface{} {
	types := make([]interface{}, 0)
	for _, v := range data {
		types = append(types, v)
	}
	return types
}
