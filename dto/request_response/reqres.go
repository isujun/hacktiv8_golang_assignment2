package request_response

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func SuccessCreateResponse(dataRes any) *Response {
	return &Response{
		Status: http.StatusCreated,
		Data:   dataRes,
	}
}

func SuccessFindResponse(dataRes any) *Response {
	return &Response{
		Status: http.StatusOK,
		Data: dataRes,
	}
}

func SuccessUpdateResponse(dataRes any) *Response {
	return &Response{
		Status: http.StatusOK,
		Data: dataRes,
	} 
}

func SuccessDeleteResponse(message string) *Response {
	return &Response{
		Status: http.StatusOK,
		Message: message,
	} 
}

func BadRequestResponse(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "invalid json request",
		Error:   err.Error,
	}
}

func InternalServerErrorResponse(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
		Error:   err.Error,
	}
}

func NotFoundResponse(message string) *Response {
	return &Response{
		Status: http.StatusNotFound,
		Message: message,
	}
}

func UnauthorizedResponse(message string) *Response {
	return &Response{
		Status: http.StatusUnauthorized,
		Message: message,
	}
}