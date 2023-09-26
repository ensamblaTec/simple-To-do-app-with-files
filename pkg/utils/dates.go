package utils

import (
	"strconv"
	"time"
)

func StringToDate(info string) (time.Time, error) {
	format := "02/01/2006 15:04:05"
	date, err := time.Parse(format, info)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func StringToUint(number string) (uint, error) {
	convertInt, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}
	return uint(convertInt), nil
}
