package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HTTP200(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, GenH(http.StatusOK, msg, data))
}

func HTTP400(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, GenH(http.StatusBadRequest, msg, data))
}

func HTTP403Abort(c *gin.Context, msg string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusForbidden, GenH(http.StatusBadRequest, msg, data))
}

func GenH(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}
