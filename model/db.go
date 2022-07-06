package model

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"go.mod/config"
)

var (
	db  *gorm.DB
	err error
)

// InitDB to test connect db
func InitDB() {
	// todo golang to get config,这里考虑下是用配置文件的方式好呢还是，传参的方式好
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	_, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数:%s", err)
		os.Exit(1)
	}
}
