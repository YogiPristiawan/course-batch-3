package presentation

import (
	"course/domain"

	"github.com/gin-gonic/gin"
)

func ReadRestIn[T interface{}](c *gin.Context, in *T) bool {
	if err := c.ShouldBindJSON(in); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return false
	}

	return true
}

func WriteRestOut[T interface{}](c *gin.Context, out T, cr *domain.CommonResult) {
	if cr.ResErrorCode == 0 {
		c.JSON(200, out)
		return
	}

	if cr.ResErrorCode >= 400 && cr.ResErrorCode < 500 {
		c.AbortWithStatusJSON(cr.ResErrorCode, struct {
			Message string `json:"message"`
		}{
			Message: cr.ResErrorMessage,
		})
		return
	}

	c.AbortWithStatusJSON(500, gin.H{"message": "internal server error"})
	return
}
