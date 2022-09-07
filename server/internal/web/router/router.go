package router

import (
	"fmt"
	"github.com/google/wire"
	"server/internal/config"
	"server/internal/engine"
)

type Router struct {
	*engine.WebEngine
}

var Provider = wire.NewSet(NewRouter)

func NewRouter(config *config.Config,engine *engine.WebEngine) *Router{

	User(engine)

	engine.Engine.Run(fmt.Sprintf(":%s",config.Web.Port))
	return &Router{
		WebEngine: engine,
	}
}
