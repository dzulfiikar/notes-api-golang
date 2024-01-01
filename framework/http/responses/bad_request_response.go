package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func NewErrorResponse(data interface{}) *ErrorResponse {
	return &ErrorResponse{
		"Error",
		400,
		data,
	}
}

func (response *ErrorResponse) Send(c *gin.Context) {
	if dataMap, ok := response.Data.(map[string]interface{}); ok {
		code, ok := dataMap["code"].(int)
		if ok {
			response.Code = code
		}
		delete(dataMap, "code")
	}
	response.Message = http.StatusText(response.Code)
	c.JSON(response.Code, response)
}
