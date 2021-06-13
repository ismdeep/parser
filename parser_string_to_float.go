package parser

import "strconv"

func ToFloat32(s string) (float32, error) {
	val, err := ToFloat64(s)
	val32 := float32(val)
	return val32, err
}

func ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
