package view

import (
	"github.com/gin-gonic/gin"
)

const errorStatus = "ERROR"

type ErrResp struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Errors string `json:"errors"`
} // @Name ErrorResponse

type ErrItem struct {
	Cause   string `json:"cause"`
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
} // @Name ErrorItemResponse

func MakeErrResp(c *gin.Context, code int, err string) {
	errResp := &ErrResp{
		Status: errorStatus,
		Code:   code,
		Errors: err,
	}
	if code == 400 {

	}
	c.JSON(errResp.Code, errResp)
}

// func getHTTPStatusCode(err error) int {
// 	return 422
// }

// func getRespErrors(err error) (errs []*ErrItem) {
// 	return errs
// }

// func getHTTPStatusCode(err error) (code int) {
// 	switch err := util.TypeOfErr(err); {
// 	case err.IsType(util.ConvertInputToDomainErr):
// 		return http.StatusBadRequest
// 	case err.IsType(util.ValidationCreateErr):
// 		return http.StatusUnprocessableEntity
// 	case err.IsType(util.ValidationUpdateErr):
// 		return http.StatusUnprocessableEntity
// 	case err.IsType(util.RepoReadErr):
// 		return http.StatusNoContent
// 	case err.IsType(util.RepoListErr):
// 		return http.StatusNoContent
// 	case err.IsType(util.RepoCreateErr):
// 		return http.StatusInternalServerError
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }

// func getRespErrors(err error) (errs []*ErrItem) {
// 	switch err.(type) {
// 	case *util.Error:
// 		return utilToResp(err.(*util.Error))
// 	default:
// 		ukErr := util.UnknownErr(err)
// 		return []*ErrItem{
// 			{
// 				Cause:   ukErr.Error(),
// 				Code:    ukErr.Code,
// 				SubCode: ukErr.SubCode,
// 			},
// 		}
// 	}
// }

// func utilToResp(err *util.Error) w {
// 	switch err := util.TypeOfErr(err); {
// 	case err.IsType(util.ValidationCreateErr):
// 		return validateToResp(err)
// 	case err.IsType(util.ValidationUpdateErr):
// 		return validateToResp(err)
// 	default:
// 		return []*ErrItem{
// 			{
// 				Cause:   err.Error(),
// 				Code:    err.Code,
// 				SubCode: err.SubCode,
// 			},
// 		}
// 	}
// }

// func validateToResp(err *util.Error) (errs []*ErrItem) {
// 	vErrs := err.Cause.(validator.ValidationErrors)
// 	errs = make([]*ErrItem, len(vErrs))
// 	for i, vErr := range vErrs {
// 		errs[i] = &ErrItem{
// 			Cause:   vErr.Translate(nil),
// 			Code:    vErr.Tag(),
// 			SubCode: vErr.Field(),
// 		}
// 	}

// 	return errs
// }
