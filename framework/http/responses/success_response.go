package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		"Success",
		200,
		data,
	}
}

func (response *SuccessResponse) Send(c *gin.Context) {
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
