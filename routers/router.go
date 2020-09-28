package routers

import (
	"kline/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("kline", &controllers.KLineController{}, "get:Show")
	beego.Router("kline/:instrument:string", &controllers.KLineController{}, "get:Show")
}
