package config

import (
	mconfig "enceremony-be/commons/clients/mysql/config"
	"enceremony-be/internal/common/logger/conf"
	"enceremony-be/internal/config/aws_conf"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Config struct {
	Aws    string
	Logger conf.LoggerConf
	Mysql  mconfig.MysqlConfig
}

func NewConfig() (*Config, error) {
	var config Config
	var confStr string
	var err error

	envKey := "enceremony"

	secretName := fmt.Sprintf("prod/%s/all", envKey)

	confSvc := aws_conf.NewConfigurationService()
	if confStr, err = confSvc.GetConfigString(envKey, secretName); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(confStr), &config); err != nil {
		log.Fatalf("Couldn't unmarshal json => %s to conf. Error %s", confStr, err)
		return nil, err
	}

	return &config, nil
}

func NewLoggerConf(config *Config) *conf.LoggerConf {
	return &config.Logger
}

func NewMysqlConf(conf *Config) *mconfig.MysqlConfig {
	mc := conf.Mysql
	return &mconfig.MysqlConfig{
		UserName:          mc.UserName,
		Password:          mc.Password,
		Host:              mc.Host,
		Port:              mc.Port,
		DbName:            mc.DbName,
		ConnectionTimeOut: int64(time.Hour),
		//Verbose:           IsStage(config),
		Verbose:      true,
		MaxOpenConns: 15,
		MaxIdleConns: 10,
		PrepareStmt:  true,
	}
}
