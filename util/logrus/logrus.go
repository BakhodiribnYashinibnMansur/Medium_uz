package logrus

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (log *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{log.WithField(k, v)}
}

func init() {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		EnvironmentOverrideColors: true,
		DisableColors:             false,
		FullTimestamp:             true,
	}
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(log)
}

// logger.WithFields(logrus.Fields{
// 		"details": logrus.Fields{
// 			"step1": "hello",
// 			"step2": "hi",
// 		},
// 	}).Debugln("some debug info")
