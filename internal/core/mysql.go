package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"otter-dingtalk/internal/global"
)

func Gorm() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.GL_MYSQL_USER, global.GL_MYSQL_PASS, global.GL_MYSQL_HOST, global.GL_MYSQL_DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		global.GL_LOG.Fatalln("mysql conntent error!")
	} else {
		global.GL_LOG.Info("mysql conntent success..")
	}
	return db
}
