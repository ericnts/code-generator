package orm

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

type Logger struct {
	LogLevel logger.LogLevel
}

// LogMode log mode
func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *Logger) Info(_ context.Context, sql string, params ...interface{}) {
	log.Info(sql, params)
}
func (l *Logger) Warn(_ context.Context, sql string, params ...interface{}) {
	log.Warn(sql, params)
}
func (l *Logger) Error(_ context.Context, sql string, params ...interface{}) {
	log.Error(sql, params)
}

func (l Logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rowsAffected := fc()
	sql = strings.ReplaceAll(sql, "\n", " ")
	sql = strings.ReplaceAll(sql, "\t", "")

	log.Debug(fmt.Sprintf("<%v> slow sql statement: %s, rowsAffected: %d",
		time.Since(begin), sql, rowsAffected))

}
