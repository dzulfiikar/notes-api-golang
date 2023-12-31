package responses

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		"Success",
		data,
	}
}

func (response *SuccessResponse) Send(c *gin.Context) {
	c.JSON(200, response)
}

func (response *SuccessResponse) SendWithStatus(c *gin.Context, status int) {
	c.JSON(status, response)
}
