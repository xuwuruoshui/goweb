package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"server/comm/result"
	"server/internal/config"
)

// router
type WebEngine struct {
	Engine *gin.Engine
}

var Provider = wire.NewSet(NewWebEngine)

func NewWebEngine(c *config.Config) *WebEngine {
	engine := gin.Default()
	// 中间件
	return &WebEngine{
		Engine: engine,
	}
}

type RouterHandler func(ctx *gin.Context) *result.Result

func (r *WebEngine) GET(path string, handler RouterHandler, handlers ...gin.HandlerFunc) {
	postHanlder := PostHandler(handler)
	handlers = append(handlers, postHanlder)
	r.Engine.GET(path, handlers...)
}

func (r *WebEngine) POST(path string, handler RouterHandler, handlers ...gin.HandlerFunc) {
	postHanlder := PostHandler(handler)
	handlers = append(handlers, postHanlder)
	r.Engine.POST(path, handlers...)
}

func (r *WebEngine) PUT(path string, handler RouterHandler, handlers ...gin.HandlerFunc) {
	postHanlder := PostHandler(handler)
	handlers = append(handlers, postHanlder)
	r.Engine.PUT(path, handlers...)
}

func (r *WebEngine) DELETE(path string, handler RouterHandler, handlers ...gin.HandlerFunc) {
	postHanlder := PostHandler(handler)
	handlers = append(handlers, postHanlder)
	r.Engine.DELETE(path, handlers...)
}

func (r *WebEngine) Group(path string, handlers ...gin.HandlerFunc) *RouterGroupPlus {
	return &RouterGroupPlus{
		routerGroup: r.Engine.Group(path, handlers...),
	}
}

type RouterGroupPlus struct {
	routerGroup *gin.RouterGroup
}

func (r *RouterGroupPlus) GET(path string, handler RouterHandler, handlers ...gin.HandlerFunc) {
	postHanlder := PostHandler(handler)
	handlers = append(handlers, postHanlder)
	r.routerGroup.GET(path, handlers...)
}

func (r *RouterGroupPlus) POST(path string, handler RouterHandler, handlers ...gin.HandlerFunc) {
	postHanlder := PostHandler(handler)
	handlers = append(handlers, postHanlder)
	r.routerGroup.POST(path, handlers...)
}

func (r *RouterGroupPlus) PUT(path string, handler RouterHandler, handlers ...gin.HandlerFunc) {
	postHanlder := PostHandler(handler)
	handlers = append(handlers, postHanlder)
	r.routerGroup.PUT(path, handlers...)
}

func (r *RouterGroupPlus) DELETE(path string, handler RouterHandler, handlers ...gin.HandlerFunc) {
	postHanlder := PostHandler(handler)
	handlers = append(handlers, postHanlder)
	r.routerGroup.DELETE(path, handlers...)
}

// 统一返回
func PostHandler(handler RouterHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, handler(ctx))
	}
}

