package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	Logs = logs.GetBeeLogger()
)

func init() {
	Logs.EnableFuncCallDepth(true)
	Logs.Async()
	Logs.SetLogFuncCallDepth(2)
	//beego.LoadAppConfig("ini","C:\\Users\\dongxiao\\Desktop\\goLearning\\src\\spider.ohmgood.com\\conf\\app.conf")
	logPath := beego.AppConfig.String("log_path")
	fileName := logPath + "/" + time.Now().Format("2006-01-02_15") + ".log"
	Logs.SetLogger(logs.AdapterFile, `{"filename":"`+fileName+`"}`)
}
