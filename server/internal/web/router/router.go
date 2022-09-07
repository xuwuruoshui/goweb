package router

import (
	"fmt"
	"github.com/google/wire"
	"server/internal/config"
	"server/internal/engine"
	"server/internal/web/middleware"
)

type Router struct {
	*engine.WebEngine
}

var Provider = wire.NewSet(NewRouter)

func NewRouter(config *config.Config,engine *engine.WebEngine) *Router{

	// 1.全局异常
	engine.Engine.Use(middleware.Recover())

	User(engine)

	engine.Engine.Run(fmt.Sprintf(":%s",config.Web.Port))
	return &Router{
		WebEngine: engine,
	}
}
