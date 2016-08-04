package routers

import (
	"news-module/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/ueditor", &controllers.UeditorController{}, "*:UEController")
}
