package responses

import "github.com/gin-gonic/gin"

type Error struct {
	Type int
	Message string
}

func HTTPResponseNotFound(c *gin.Context, message string) {
	e := Error{404, message}
	c.JSON(404, gin.H{
		"type": e.Type,
		"message": e.Message,
	})
}

func HTTPResponseInternalServerError(c *gin.Context, message string) {
	e := Error{500, message}
	c.JSON(500, gin.H{
		"type": e.Type,
		"message": e.Message,
	})
}