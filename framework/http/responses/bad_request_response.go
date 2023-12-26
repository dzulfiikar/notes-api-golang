package responses

import (
	"github.com/gin-gonic/gin"
)

type BadRequestError struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBadRequestError(data interface{}) *BadRequestError {
	return &BadRequestError{
		"Bad Request",
		data,
	}
}

func (response *BadRequestError) Send(c *gin.Context) {
	c.JSON(400, response)
}

func (response *BadRequestError) SendWithStatus(c *gin.Context, status int) {
	c.JSON(status, response)
}

func (response *BadRequestError) SendWithStatusAndMessage(c *gin.Context, status int, message string) {
	response.Message = message
	c.JSON(status, response)
}

func (response *BadRequestError) SendWithStatusAndData(c *gin.Context, status int, data interface{}) {
	response.Data = data
	c.JSON(status, response)
}

func (response *BadRequestError) SendWithStatusMessageAndData(c *gin.Context, status int, message string, data interface{}) {
	response.Message = message
	response.Data = data
	c.JSON(status, response)
}

func (response *BadRequestError) SendWithStatusMessageAndDataAndError(c *gin.Context, status int, message string, data interface{}, err error) {
	response.Message = message
	response.Data = data
	c.JSON(status, response)
}
