package applog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

const logTimePattern = "2006-01-02 15:04:05.000"
const emptyPlaceHolder = "-"
const requestStartTimeKey = "__request_start_time__"
const queryPathKey = "__query_path__"
const ginContextLogKey = "__gin_context_log_key__"

type formatter struct{}

func (formatter) Format(entry *logrus.Entry) ([]byte, error) {
	logTime := entry.Time.Format(logTimePattern)
	level := emptyPlaceHolder

	switch entry.Level {
	case logrus.DebugLevel:
		level = "DEBUG"
	case logrus.InfoLevel:
		level = "INFO_"
	case logrus.WarnLevel:
		level = "WARN_"
	case logrus.ErrorLevel:
		level = "ERROR"
	case logrus.FatalLevel:
		level = "FATAL"
	case logrus.PanicLevel:
		level = "PANIC"
	}

	m := entry.Data
	queryPath := getFromMap(m, queryPathKey, emptyPlaceHolder)
	message := entry.Message

	s := fmt.Sprintf("%s|%s|%s|%s\n", logTime, level, queryPath, message)

	return []byte(s), nil
}

func getFromMap(m map[string]interface{}, key, default_ string) interface{} {
	if v, ok := m[key]; ok {
		return v
	}
	return default_
}

func ConfigLocalFilesystemLogger(logPath, logFileName string) {
	baseLogPath := path.Join(logPath, logFileName)

	accessLogWriter, err := rotatelogs.New(
		baseLogPath+".log.%Y-%m-%d",
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithRotationCount(30),
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	errorLogWriter, err := rotatelogs.New(
		baseLogPath+"-error.log.%Y-%m-%d",
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithRotationCount(30),
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: accessLogWriter,
		logrus.InfoLevel:  accessLogWriter,
		logrus.WarnLevel:  errorLogWriter,
		logrus.ErrorLevel: errorLogWriter,
		logrus.FatalLevel: errorLogWriter,
		logrus.PanicLevel: errorLogWriter,
	}, &formatter{})

	logrus.AddHook(lfHook)
}

func GinLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		fields := map[string]interface{}{
			requestStartTimeKey: time.Now(),
			queryPathKey:        context.Request.URL.Path,
		}

		context.Set(ginContextLogKey, logrus.WithFields(fields))
		context.Next()
	}
}

func GetFromGin(c *gin.Context) *logrus.Entry {
	return c.Value(ginContextLogKey).(*logrus.Entry)
}

func init() {
	logrus.SetFormatter(formatter{})
}
