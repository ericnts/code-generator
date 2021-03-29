package orm

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"time"
)

var DB *gorm.DB

func init() {
	log.Info("开始初始化数据库...")

	masterDialector, err := getDialector()
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(masterDialector, &gorm.Config{
		Logger:      &Logger{LogLevel: logger.Info},
		PrepareStmt: true,
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetConnMaxLifetime(time.Hour * 24)
}

func getDialector() (gorm.Dialector, error) {
	u := url.URL{
		Scheme:   "mysql",
		User:     url.UserPassword("airobot", "i3o3aFYUPr"),
		Host:     "tcp(192.168.6.101:3306)",
		Path:     "airobot_dev",
		RawQuery: "charset=utf8&parseTime=True&loc=Local",
	}
	dialector := mysql.Open(u.String()[8:])

	return dialector, nil
}
