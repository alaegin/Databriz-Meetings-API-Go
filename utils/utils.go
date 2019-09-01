package utils

import (
	"strconv"
)

func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
