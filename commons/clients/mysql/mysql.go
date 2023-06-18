package mysql

import (
	"database/sql"
	"enceremony-be/commons/clients/mysql/config"
	"enceremony-be/internal/common/logger"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

type Connection interface {
	GetInstance() *gorm.DB
}

type ConnectionImpl struct {
	conn   *gorm.DB
	conf   *config.MysqlConfig
	logger *logger.Logger
}

func NewMysqlConnection(config *config.MysqlConfig, logger *logger.Logger) (Connection, error) {
	var conn *gorm.DB
	var err error

	logger.Infow("MysqlConnection", zap.String("ctx", "init"))
	conn, err = createConnection(config, logger)
	if err != nil {
		return nil, err
	}

	return &ConnectionImpl{
		conf:   config,
		logger: logger,
		conn:   conn,
	}, nil
}

func (m *ConnectionImpl) GetInstance() *gorm.DB {

	return m.conn
}

func createConnection(conf *config.MysqlConfig, logger *logger.Logger) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Errorw("MysqlConnection", zap.String("ctx", "MysqlOpen"), zap.Error(err))
		return nil, err
	}

	dbOrm, err := gorm.Open(
		mysql.New(
			mysql.Config{
				Conn: db,
			}),
		&gorm.Config{
			Logger:      getZapGormLogger(conf, logger),
			PrepareStmt: conf.PrepareStmt,
		})

	if err != nil {
		logger.Errorw("MysqlConnection", zap.String("ctx", "GoOrmOpen"), zap.Error(err))
		return nil, err
	}

	sqlDB, err := dbOrm.DB()
	if err != nil {
		logger.Errorw("MysqlConnection", zap.String("ctx", "DB"), zap.Error(err))
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnectionTimeOut) * time.Minute)
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	return dbOrm, nil
}

func getZapGormLogger(conf *config.MysqlConfig, logger *logger.Logger) zapgorm2.Logger {
	zapGormLogger := zapgorm2.New(logger.Desugar())
	zapGormLogger.SetAsDefault()

	var ll gormLogger.LogLevel
	if conf.Verbose {
		ll = gormLogger.Info
	} else {
		ll = gormLogger.Warn
	}
	zapGormLogger.LogLevel = ll
	zapGormLogger.SlowThreshold = 1 * time.Second
	zapGormLogger.SkipCallerLookup = true
	zapGormLogger.IgnoreRecordNotFoundError = true

	return zapGormLogger
}
