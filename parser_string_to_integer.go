package parser

import "strconv"

// ToInt ToInt
func ToInt(s string) (int, error) {
	val, err := ToInt64(s)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

// ToInt32 ToInt32
func ToInt32(s string) (int32, error) {
	val, err := ToInt64(s)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// ToInt64 ToInt64
func ToInt64(s string) (int64, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// ToUint ToUint
func ToUint(s string) (uint, error) {
	val, err := ToUint64(s)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

// ToUint32 ToUint32
func ToUint32(s string) (uint32, error) {
	val, err := ToUint64(s)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

// ToUint64 ToUint64
func ToUint64(s string) (uint64, error) {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, err
}
