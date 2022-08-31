package log

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/utils"
)

const (
	TRACE = iota
	DEBUG
	INFO
	NOTICE
	WARN
	ERROR
	PANIC
	ACCESS //这个最大
)

var levelStr map[int]string = map[int]string{DEBUG: "DEBUG", INFO: "INFO", NOTICE: "NOTICE", WARN: "WARN", ERROR: "ERROR", ACCESS: "ACCESS"}

type LogConf struct {
	Level int
	Path  string
}

var Logger *LogConf

type LogInterface interface {
	Debug(*gin.Context, string, map[string]interface{})
	Info(*gin.Context, string, map[string]interface{})
	Notice(*gin.Context, string, map[string]interface{})
	Warn(*gin.Context, string, map[string]interface{})
	Error(*gin.Context, string, map[string]interface{})
}

func Init(logConf *LogConf) {
	Logger = logConf
}

func Access(c *gin.Context, tag string, data map[string]interface{}) {
	if Logger.Level > ACCESS {
		return
	}
	writeFile(c, "ACCESS", ACCESS, tag, data)
}

func Debug(c *gin.Context, tag string, data map[string]interface{}) {
	if Logger.Level > DEBUG {
		return
	}
	writeFile(c, "DEBUG", DEBUG, tag, data)
}

func Notice(c *gin.Context, tag string, data map[string]interface{}) {
	if Logger.Level > NOTICE {
		return
	}
	writeFile(c, "NOTICE", NOTICE, tag, data)
}

func Info(c *gin.Context, tag string, data map[string]interface{}) {
	if Logger.Level > INFO {
		return
	}
	writeFile(c, "INFO", INFO, tag, data)
}

func Warn(c *gin.Context, tag string, data map[string]interface{}) {
	if Logger.Level > WARN {
		return
	}
	writeFile(c, "WARN", WARN, tag, data)
}

func Error(c *gin.Context, tag string, data map[string]interface{}) {
	if Logger.Level > ERROR {
		return
	}
	writeFile(c, "ERROR", ERROR, tag, data)
}

func Panic(c *gin.Context, tag string, data map[string]interface{}) {
	if Logger.Level > ERROR {
		return
	}
	writeFile(c, "PANIC", PANIC, tag, data)
}

//指定写入的文件名与日志等级,方便以后灵活调用
func LogFile(c *gin.Context, fileName string, level int, tag string, data map[string]interface{}) {
	if Logger.Level > level {
		return
	}
	writeFile(c, fileName, level, tag, data)
}

func writeFile(c *gin.Context, fileName string, level int, tag string, data map[string]interface{}) error {
	var buf strings.Builder

	buf.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	buf.WriteString("`")
	buf.WriteString(utils.GetTraceIdFromCtx(c))
	buf.WriteString("`")
	buf.WriteString(levelStr[level])
	buf.WriteString("`")
	buf.WriteString(tag)
	buf.WriteString("`")
	dataStr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	buf.WriteString(string(dataStr))
	buf.WriteString("\n")
	fileName = Logger.Path + fileName + "_" + time.Now().Format("2006-01-02") + ".log"
	return utils.WriteFile(c, fileName, buf.String())
}
