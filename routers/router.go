package routers

import (
	"energy_post_dom/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1/api",
		//	api cache checking
		//beego.NSBefore(models.ReadApiCache),

		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		),
	)

	beego.AddNamespace(ns)
}
