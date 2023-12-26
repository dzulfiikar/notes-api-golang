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

func (response *SuccessResponse) SendWithStatusAndMessage(c *gin.Context, status int, message string) {
	response.Message = message
	c.JSON(status, response)
}

func (response *SuccessResponse) SendWithStatusAndData(c *gin.Context, status int, data interface{}) {
	response.Data = data
	c.JSON(status, response)
}

func (response *SuccessResponse) SendWithStatusMessageAndData(c *gin.Context, status int, message string, data interface{}) {
	response.Message = message
	response.Data = data
	c.JSON(status, response)
}

func (response *SuccessResponse) SendWithStatusMessageAndDataAndError(c *gin.Context, status int, message string, data interface{}, err error) {
	response.Message = message
	response.Data = data
	c.JSON(status, response)
}
