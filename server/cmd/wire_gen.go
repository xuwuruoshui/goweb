// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"server/internal/app"
	"server/internal/config"
	"server/internal/db"
	"server/internal/engine"
	"server/internal/web/router"
)

// Injectors from wire.go:

func InitApp() (*app.App, error) {
	configConfig := config.NewConfig()
	gormDB, err := db.NewDB(configConfig)
	if err != nil {
		return nil, err
	}
	webEngine := engine.NewWebEngine(configConfig)
	routerRouter := router.NewRouter(configConfig, webEngine)
	appApp := app.NewApp(configConfig, gormDB, routerRouter)
	return appApp, nil
}