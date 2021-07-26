package parser

import "strconv"

func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}
