package util

import (
	errPkg "LinkShortening/internals/myerror"
	"strconv"
)

func InterfaceConvertInt(value interface{}) (int, error) {
	var intConvert int
	var errorConvert error
	switch value.(type) {
	case string:
		intConvert, errorConvert = strconv.Atoi(value.(string))
		if errorConvert != nil {
			return errPkg.IntNil, &errPkg.Errors{
				Text: errPkg.ErrAtoi,
			}
		}
		return intConvert, nil
	case int:
		intConvert = value.(int)
		return intConvert, nil
	default:
		return errPkg.IntNil, &errPkg.Errors{
			Text: errPkg.ErrNotStringAndInt,
		}
	}
}
