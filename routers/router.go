package routers

import (
	"ranqu/controllers/users"

	"github.com/astaxie/beego"
)

func init() {

	//首页
	beego.Router("/", &users.MainController{})
	beego.Router("/homepage", &users.MainController{})
}
