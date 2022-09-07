package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"server/internal/config"
	"server/internal/model"
	"time"
)

var Provider = wire.NewSet(NewDB)

func NewDB(c *config.Config)(*gorm.DB,error){
	var ormLogger logger.Interface

	if gin.Mode()=="debug"{
		ormLogger = logger.Default.LogMode(logger.Info)
	}else{
		ormLogger = logger.Default
	}

	var dsn string
	if c.Database.Type=="mysql"{
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",c.Database.Username,c.Database.Password,c.Database.Url,
			c.Database.Port,c.Database.DBname,c.Database.Params)
	}


	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 不加s
		},
	})
	if err!=nil{
		return nil,err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second*10)


	return db,nil
}

func Migration(db *gorm.DB){
	db.Set("gorm:table_options", "ENGINE=InnoDB charset=utf8mb4")
	db.AutoMigrate(&model.User{})
}