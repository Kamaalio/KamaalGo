package ginErrors

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type errorMessage struct {
	Message string `json:"message"`
}

func ErrorHandler(context *gin.Context, errorObject Error) {
	context.AbortWithStatusJSON(errorObject.Status, errorMessage{Message: errorObject.Message})
}
