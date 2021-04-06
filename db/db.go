package db

import (
	"fmt"
	"github.com/ericnts/code-generator/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

var DB *gorm.DB

func init() {
	log.Info("开始初始化数据库...")
	u := url.URL{
		Scheme:   "mysql",
		User:     url.UserPassword(config.DBUsername, config.DBPassword),
		Host:     fmt.Sprintf("tcp(%s:%s)", config.DBHost, config.DBPort),
		Path:     config.DBDatabase,
		RawQuery: "charset=utf8&parseTime=True&loc=Local",
	}
	dialector := mysql.Open(u.String()[8:])
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger:      &Logger{LogLevel: logger.Info},
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	DB = db
}
