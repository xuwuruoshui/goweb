package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Database database `json:"database"`
	Web database `json:"web"`
}

type Web struct {
	Port int `json:"port"`
}

type database struct {
	Type string `json:"type"`
	Url string `json:"url"`
	Port string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBname  string `json:"dbname"`
	Params string `json:"params"`
	MaxIdleConns int `json:"maxIdleConns"`
	MaxOpenConns int `json:"MaxOpenConns"`
}

type cache struct {
	Type string `json:"type"`
	Url string `json:"url"`
	Port string `json:"port"`
	Password string `json:"password"`
}

var Provider = wire.NewSet(NewConfig)

func NewConfig() *Config{
	viper.SetConfigFile("../config/app.yml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Read file error:" + err.Error())
	}
	var conf = &Config{}
	if err := viper.Unmarshal(conf); err != nil {
		panic("File exchange error:" + err.Error())
	}
	// 监控配置文件的变化
	viper.WatchConfig()
	// viper文件发送变化事件
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(conf); err != nil {
			panic("File exchange error:" + err.Error())
		}
	})
	return conf
}
