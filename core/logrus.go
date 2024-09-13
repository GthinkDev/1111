package core

import (
	"bytes"
	"fmt"
	"gin-blog/global"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = "\033[31m"
	gray   = "\033[39m"
	yellow = "\033[33m"
	blue   = "\033[36m"
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的 level 输出不同的颜色
	var levelColor string
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 自定义日期格式
	log := global.Config.Logger
	timeFormat := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "%s%s [%s] --> [%s] %s  %s %s\n", levelColor, log.Prefix, timeFormat, entry.Level, funcVal, fileVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s%s [%s] --> [%s] %s\n", log.Prefix, levelColor, timeFormat, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                // 新建一个实例
	mLog.SetFormatter(&LogFormatter{})                  // 设置输出格式
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数和行号
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)      // 设置日志级别
	mLog.SetOutput(os.Stdout) // 设置输出位置
	InitDefaultLogger()
	return mLog
}

func InitDefaultLogger() {
	// 全局 log
	logrus.SetFormatter(&LogFormatter{})                  // 设置输出格式
	logrus.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数和行号
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)      // 设置日志级别
	logrus.SetOutput(os.Stdout) // 设置输出位置
}
