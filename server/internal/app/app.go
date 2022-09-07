package app

import (
	"gorm.io/gorm"
	"server/internal/config"
	"server/internal/web/router"
)

type App struct { // 最终需要的对象
	DB *gorm.DB
	Router *router.Router
}

func NewApp(config *config.Config,db *gorm.DB,router *router.Router)  *App{
	return &App{
		DB:db,
		Router: router,
	}
}
