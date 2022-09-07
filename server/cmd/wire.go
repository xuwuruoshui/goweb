// +build wireinject

package main

import (
	"github.com/google/wire"
	"server/internal/app"
	"server/internal/config"
	"server/internal/db"
	"server/internal/engine"
	"server/internal/web/router"
)

func InitApp()(*app.App,error){

	panic(wire.Build(config.Provider,
		db.Provider,
		engine.Provider,
		router.Provider,
		app.NewApp))
}