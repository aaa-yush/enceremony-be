package config

import "time"

type MysqlConfig struct {
	UserName          string
	Password          string
	Host              string
	Port              int
	DbName            string
	ConnectionTimeOut int64
	Verbose           bool          `json:"v"`
	ConnMaxIdleTime   time.Duration `json:"conn_max_idle_time,omitempty"`
	MaxOpenConns      int           `json:"max_open_conns,omitempty"`
	MaxIdleConns      int           `json:"max_idle_conns,omitempty"`
	PrepareStmt       bool          `json:"prepare_stmt"`
}
