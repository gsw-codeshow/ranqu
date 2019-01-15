package main

import (
	_ "ranqu/initial"
	_ "ranqu/routers"

	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.Listen.EnableAdmin = true

	beego.Run()
}
