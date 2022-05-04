package rest

type BizCode struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

var (
	CodeSuccess      = 0
	CodeInValidParam = 100001
	CodeNotFound     = 100002
	CodeUnKnown      = 100003

	CodeMessages = map[int]string{
		CodeSuccess:      "success",
		CodeNotFound:     "data not found",
		CodeInValidParam: "invalid input param",
		CodeUnKnown:      "unknown error",
	}
)

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"message"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}

type PageRequest struct {
	PageNum  int `form:"page" json:"page" xml:"page"`
	PageSize int `form:"size" json:"size" xml:"size"`
}

type PageResponse struct {
	Total     int64 `json:"total"`
	PageTotal int64 `json:"page_total"`
}

type PageItems struct {
	PageRequest
	PageResponse
	Items interface{} `json:"items"`
}

func NewResponse(code int, data interface{}) *Response {
	res := &Response{
		Code: code,
	}
	res.Msg, _ = CodeMessages[code]
	if code == CodeSuccess {
		res.Data = data
	} else {
		if err, ok := data.(error); ok {
			res.Detail = err.Error()
		} else {
			res.Detail, _ = data.(string)
		}
	}
	return res
}
