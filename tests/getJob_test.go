package test

import (
	"github.com/astaxie/beego"
	"path/filepath"
	"runtime"
	"spider.ohmgood.com/controllers/command"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestGetJob(t *testing.T) {
	command.Main()
}
