package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	name := c.GetString("name")
	sex := c.GetString("sex")
	if name == "" {
		c.Ctx.WriteString("Hello Guest!")
		return
	} else if sex == "" {
		c.Data["Name"] = name
		c.Data["Sex"] = ""
	} else if sex == "man" {
		c.Data["Name"] = name
		c.Data["Sex"] = "Mr "
	} else {
		c.Data["Name"] = name
		c.Data["Sex"] = "Miss "
	}
	c.TplName = "index.tpl"
}
