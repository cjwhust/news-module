package routers

import (
	"news-module/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Get")
    beego.Router("/ueditor", &controllers.UeditorController{}, "*:UEController")
    beego.Router("/news", &controllers.MainController{}, "get:GetOne")
    beego.Router("/content", &controllers.MainController{}, "get:Content")
    beego.Router("/save", &controllers.MainController{}, "post:Save")
    beego.Router("/getFlags", &controllers.MainController{}, "post:GetFlags")
    beego.Router("/getInfo", &controllers.MainController{}, "get:GetInfo")
}
