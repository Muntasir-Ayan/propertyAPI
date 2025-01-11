package routers

import (
    "propertyAPI/controllers"

     beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
    beego.Router("/v1/property/list", &controllers.PropertyController{})
    beego.Router("/v1/property/details", &controllers.PropertyController{})
}