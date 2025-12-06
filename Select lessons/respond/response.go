package respond

import (
	"errors"
)

type Response struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}
type FinalResponse struct {
	Status string      `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

var (
	UserNotFound     = Response{Status: "1001", Info: "用户不存在"}
	WrongPassword    = Response{Status: "1002", Info: "参数错误"}
	ResourceNotFound = Response{Status: "1004", Info: "资源不存在"}
	NoLeft           = Response{Status: "1005", Info: "课程已被选完"}
)

func (r Response) Error() string {
	return r.Info
}
func InternalError(err error) Response {
	return Response{
		Status: "500",
		Info:   err.Error(),
	}
}

func HandleError(err error, data interface{}) FinalResponse {
	if err == nil {
		return FinalResponse{
			Status: "200",
			Info:   "success",
			Data:   data,
		}
	}
	var resp Response
	if errors.As(err, &resp) {
		return FinalResponse{
			Status: resp.Status,
			Info:   resp.Info,
			Data:   data,
		}
	} else {
		final := InternalError(err)
		return FinalResponse{
			Status: final.Status,
			Info:   final.Info,
			Data:   data,
		}
	}

}
