package common

import (
	"github.com/unrolled/render"
	"io"
)

func init() {
	renderer = render.New()
}

var renderer *render.Render

type response struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}

func jsonParse(w io.Writer, data interface{}, code int, status, message string) {
	response := response{
		Data:    data,
		Code:    code,
		Status:  status,
		Message: message,
	}

	renderer.JSON(w, code, &response)

}

func Success(w io.Writer, data interface{}, code ...int) {
	if code == nil {
		code = append(code, 200)
	}
	jsonParse(w, data, code[0], "Success", "")
}
