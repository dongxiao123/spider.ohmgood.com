package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	"spider.ohmgood.com/common"
	"spider.ohmgood.com/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()

	job := models.Job{Name: "slene"}
	// insert
	id, err := o.Insert(&job)
	common.Logs.Warning("Insert")
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
