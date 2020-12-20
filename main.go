package main

import (
	"github.com/astaxie/beego"
	_ "spider.ohmgood.com/common"
	_ "spider.ohmgood.com/routers"
)

func main() {
	beego.Run()
}
