package types

type CommonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DataResponse[T any] struct {
	CommonResponse
	Data T `json:"data"`
}

func NewDataResponse[T any](code int, msg string, data T) *DataResponse[T] {
	return &DataResponse[T]{
		CommonResponse: CommonResponse{Code: code, Msg: msg},
		Data:           data,
	}
}

type ResultResponse[T any] struct {
	CommonResponse
	Result T `json:"result"`
}

func NewResultResponse[T any](code int, msg string, result T) *ResultResponse[T] {
	return &ResultResponse[T]{
		CommonResponse: CommonResponse{Code: code, Msg: msg},
		Result:         result,
	}
}

// Implementing BaseResponse interface for CommonBaseResponse
func (d *CommonResponse) GetCode() int {
	return d.Code
}

func (d *CommonResponse) GetMsg() string {
	return d.Msg
}

func (d *CommonResponse) SetCode(code int) {
	d.Code = code
}

func (d *CommonResponse) SetMsg(msg string) {
	d.Msg = msg
}
