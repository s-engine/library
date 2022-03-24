package log

import "go.uber.org/zap"

// Logger logger interface
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})

	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
}

var xLogger XLogger

type XLogger struct {
	logger Logger
}

func init() {

}

func SetLogger(log Logger) {
	if log == nil {
		zapLogger, _ := zap.NewDevelopment()
		xLogger.logger = zapLogger.Sugar()
		return
	}
	xLogger.logger = log
}

func Info(args ...interface{}) {
	xLogger.logger.Info(args)
}

func Warn(args ...interface{}) {
	xLogger.logger.Info(args)
}

func Error(args ...interface{}) {
	xLogger.logger.Info(args)
}

func Debug(args ...interface{}) {
	xLogger.logger.Info(args)
}

func Infof(fmt string, args ...interface{}) {
	xLogger.logger.Infof(fmt, args)
}

func Warnf(fmt string, args ...interface{}) {
	xLogger.logger.Infof(fmt, args)
}

func Errorf(fmt string, args ...interface{}) {
	xLogger.logger.Infof(fmt, args)
}

func Debugf(fmt string, args ...interface{}) {
	xLogger.logger.Infof(fmt, args)
}
