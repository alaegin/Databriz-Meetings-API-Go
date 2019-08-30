package utils

import (
	"strconv"
)

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
