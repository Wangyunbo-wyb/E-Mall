package orm

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DSN         string
	Active      int
	Idle        int
	IdleTimeout time.Duration
}

type DB struct {
	*gorm.DB
}

// use the go-zero log library logx to implement the logger.Interface interface of gorm
type ormLog struct {
	LogLevel logger.LogLevel
}

func (l *ormLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ormLog) Info(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Warn(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Error(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(format, v...)
}

// Trace helps developers monitor the performance of SQL queries for debugging and optimization
func (l *ormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//get the elapsed time since the beginning of the function
	elapsed := time.Since(begin)
	// get the SQL statement and the number of rows affected by the SQL statement
	sql, rows := fc()
	// log the SQL statement, the number of rows affected, and the elapsed time
	logx.WithContext(ctx).WithDuration(elapsed).Infof("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
}

func NewMysql(c *Config) *DB {
	if c == nil {
		panic("config cannot be nil")
	}

	if c.Idle == 0 {
		c.Idle = 10
	}
	if c.Active == 0 {
		c.Active = 100
	}
	if c.IdleTimeout == 0 {
		c.IdleTimeout = time.Hour
	}
	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{
		Logger: &ormLog{},
	})
	if err != nil {
		panic(err)
	}
	sdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sdb.SetMaxIdleConns(c.Idle)
	sdb.SetMaxOpenConns(c.Active)
	sdb.SetConnMaxLifetime(c.IdleTimeout)

	err = db.Use(NewCustomPlugin())
	if err != nil {
		panic(err)
	}
	return &DB{DB: db}
}
