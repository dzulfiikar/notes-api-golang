package responses

import (
	"github.com/gin-gonic/gin"
)

type CreatedResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewCreatedResponse(data interface{}) *CreatedResponse {
	return &CreatedResponse{
		"Created",
		data,
	}
}

func (response *CreatedResponse) Send(c *gin.Context) {
	c.JSON(201, response)
}

func (response *CreatedResponse) SendWithStatus(c *gin.Context, status int) {
	c.JSON(status, response)
}

func (response *CreatedResponse) SendWithStatusAndMessage(c *gin.Context, status int, message string) {
	response.Message = message
	c.JSON(status, response)
}

func (response *CreatedResponse) SendWithStatusAndData(c *gin.Context, status int, data interface{}) {
	response.Data = data
	c.JSON(status, response)
}

func (response *CreatedResponse) SendWithStatusMessageAndData(c *gin.Context, status int, message string, data interface{}) {
	response.Message = message
	response.Data = data
	c.JSON(status, response)
}

func (response *CreatedResponse) SendWithStatusMessageAndDataAndError(c *gin.Context, status int, message string, data interface{}, err error) {
	response.Message = message
	response.Data = data
	c.JSON(status, response)
}
