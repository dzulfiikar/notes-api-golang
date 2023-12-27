package responses

import (
	"github.com/gin-gonic/gin"
)

type NotFoundError struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewNotFoundError(data interface{}) *NotFoundError {
	return &NotFoundError{
		"Not Found",
		data,
	}
}

func (response *NotFoundError) Send(c *gin.Context) {
	c.JSON(404, response)
}

func (response *NotFoundError) SendWithStatus(c *gin.Context, status int) {
	c.JSON(status, response)
}

func (response *NotFoundError) SendWithMessage(c *gin.Context, message string) {
	c.JSON(404, response)
}

func (response *NotFoundError) SendWithStatusAndMessage(c *gin.Context, status int, message string) {
	response.Message = message
	c.JSON(status, response)
}

func (response *NotFoundError) SendWithStatusAndData(c *gin.Context, status int, data interface{}) {
	response.Data = data
	c.JSON(status, response)
}

func (response *NotFoundError) SendWithStatusMessageAndData(c *gin.Context, status int, message string, data interface{}) {
	response.Message = message
	response.Data = data
	c.JSON(status, response)
}

func (response *NotFoundError) SendWithStatusMessageAndDataAndError(c *gin.Context, status int, message string, data interface{}, err error) {
	response.Message = message
	response.Data = data
	c.JSON(status, response)
}
