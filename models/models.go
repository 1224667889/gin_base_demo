package models

import (
	"fmt"
	"fzuhelper_launch_screen/pkg/setting"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

)

var db *gorm.DB
func Setup() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Name,
		),
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Error,
			},
		),
	})
	if err != nil {
		log.Fatalln(err)
	}

	//if err := db.AutoMigrate(&Game{}); err != nil {
	//	log.Fatalln(err)
	//}

	if sqlDB, err := db.DB(); err != nil {
		log.Fatalln(err)
	} else {
		sqlDB.SetMaxIdleConns(setting.DatabaseSetting.MaxIdle)
		sqlDB.SetMaxOpenConns(setting.DatabaseSetting.MaxOpen)
	}
}