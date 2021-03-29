package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var DB *gorm.DB

func init() {
	log.Info("开始初始化数据库...")
	u := url.URL{
		Scheme:   "mysql",
		User:     url.UserPassword("airobot", "i3o3aFYUPr"),
		Host:     "tcp(192.168.6.101:3306)",
		Path:     "airobot_dev",
		RawQuery: "charset=utf8&parseTime=True&loc=Local",
	}
	dialector := mysql.Open(u.String()[8:])
	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	DB = db
}

