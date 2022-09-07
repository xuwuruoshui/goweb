package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"server/comm/result"
)

// 全局异常处理
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Errorf("panic %v\n", r)
				debug.PrintStack()
				c.JSON(http.StatusInternalServerError, result.UNKNOW_ERROR.SetMsg(r.(error).Error()))
				c.Abort()
			}
		}()
		c.Next()
	}

}
