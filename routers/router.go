package routers

import (
    "propertyAPI/controllers"

    beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/v1/property/list", &controllers.PropertyListController{})
    beego.Router("/v1/property/details", &controllers.PropertyDetailsController{})

    beego.Router("/destination-property-details", &controllers.DestinationPropertyDetailsController{})
    beego.Router("/hotel-details", &controllers.HotelDetailsController{})
}
