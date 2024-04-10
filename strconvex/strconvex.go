package strconvex

import "strconv"

func ParseInt32(s string, base int) int32 {
	i, _ := strconv.ParseInt(s, base, 32)
	return int32(i)
}

func TryParseInt32(s string, base int) (int32, error) {
	i, err := strconv.ParseInt(s, base, 32)
	return int32(i), err
}

func ParseInt64(s string, base int) int64 {
	i, _ := strconv.ParseInt(s, base, 64)
	return i
}

func TryParseInt64(s string, base int) (int64, error) {
	i, err := strconv.ParseInt(s, base, 64)
	return i, err
}

func ParseFloat32(s string) float32 {
	float, _ := strconv.ParseFloat(s, 32)
	return float32(float)
}

func TryParseFloat32(s string) (float32, error) {
	float, err := strconv.ParseFloat(s, 32)
	return float32(float), err
}

func ParseFloat64(s string) float64 {
	float, _ := strconv.ParseFloat(s, 64)
	return float
}

func TryParseFloat64(s string) (float64, error) {
	float, err := strconv.ParseFloat(s, 64)
	return float, err
}
