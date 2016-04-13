package main

import (
	"./routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &routers.AdminIndex{}, "get:Index")
	beego.Router("/login", &routers.Login{}, "get:Login")
	beego.Run()
}
