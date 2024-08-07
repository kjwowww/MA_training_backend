package handler

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(code int, status bool, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}
