package responses

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewErrorResponse(data interface{}) *ErrorResponse {
	return &ErrorResponse{
		"Bad Request",
		data,
	}
}

func (response *ErrorResponse) Send(c *gin.Context) {
	c.JSON(400, response)
}

func (response *ErrorResponse) SendWithStatus(c *gin.Context, status int) {
	c.JSON(status, response)
}
