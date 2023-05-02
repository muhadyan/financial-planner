package utils

import "math"

func ConvertToSliceInterface(data ...string) []interface{} {
	types := make([]interface{}, 0)
	for _, v := range data {
		types = append(types, v)
	}
	return types
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func TotalPage(lenData int, limit int) int {
	res := math.Ceil(float64((lenData / limit))) + 1

	return int(res)
}
