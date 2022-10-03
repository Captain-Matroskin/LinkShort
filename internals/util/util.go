package util

import (
	errPkg "LinkShortening/internals/myerror"
	"strconv"
)

const (
	LenLinkShort = 10
	LinkDomain   = "ozon.click.ru" // test domain
)

type Result struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body,omitempty"`
}

func InterfaceConvertInt(value interface{}) (int, error) {
	var intConvert int
	var errorConvert error
	switch value.(type) {
	case string:
		intConvert, errorConvert = strconv.Atoi(value.(string))
		if errorConvert != nil {
			return errPkg.IntNil, &errPkg.MyErrors{
				Text: errPkg.ErrAtoi,
			}
		}
		return intConvert, nil
	case int:
		intConvert = value.(int)
		return intConvert, nil
	default:
		return errPkg.IntNil, &errPkg.MyErrors{
			Text: errPkg.ErrNotStringAndInt,
		}
	}
}
