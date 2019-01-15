package users

import (
	"ranqu/controllers"
)

//用户主页
type MainController struct {
	controllers.BaseController
}

func (this *MainController) Get() {
	this.TplName = "index.tpl"
}
