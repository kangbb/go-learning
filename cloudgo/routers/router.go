package routers

import (
	"github.com/kangbb/go-learning/cloudgo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
