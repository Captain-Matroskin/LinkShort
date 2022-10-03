package myerror

import (
	"encoding/json"
	"net/http"
)

func (c *CheckError) CheckErrorCreateLinkShort(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case LSHCreateLinkShortNotInsertUnique:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusConflict,
				Explain: LSHCreateLinkShortNotInsertUnique,
			})
			if errMarshal != nil {
				println(errMarshal.Error())
				//c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &MyErrors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			println(err.Error())
			//c.Logger.Errorf("%s, requestId: %d", ADeleteCookieCookieNotDelete, c.RequestId)
			return &MyErrors{
					Text: ErrCheck,
				},
				result, http.StatusOK

		default:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				println(errMarshal.Error())
				//c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &MyErrors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			println(err.Error())
			//c.Logger.Errorf("%s, requestId: %d", ADeleteCookieCookieNotDelete, c.RequestId)
			return &MyErrors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError

		}

	}
	return nil, nil, IntNil
}
