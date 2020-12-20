package common

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/core/config"
	"time"
)

var (
	Logs = logs.GetBeeLogger()
)

func init() {
	Logs.EnableFuncCallDepth(true)
	Logs.Async()
	Logs.SetLogFuncCallDepth(2)
	logPath, _ := config.String("log_path")
	fileName := logPath + "/" + time.Now().Format("2006-01-02_15") + ".log"
	Logs.SetLogger(logs.AdapterFile, `{"filename":"`+fileName+`"}`)
}
