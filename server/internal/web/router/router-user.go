package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/comm/result"
	"server/internal/engine"
)

func User(r *engine.WebEngine) {
	r.GET("/user", func(ctx *gin.Context) *result.Result {
		fmt.Println("success")
		return result.SuccessResult(map[string]interface{}{"id":1})
	})
}

