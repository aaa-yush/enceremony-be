package logger

import (
	"enceremony-be/internal/common/logger/conf"
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(c *conf.LoggerConf) (*Logger, error) {

	z, err := NewZapLogger(c)

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
		return nil, err
	}

	l := &Logger{z}

	return l, nil
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.Errorw(msg, keysAndValues...)
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.Infow(msg, keysAndValues...)
}
func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.Debugw(msg, keysAndValues...)
}
func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.Warnw(msg, keysAndValues...)
}

// Printf For the in memory cache
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Debugf(format, v...)
}
