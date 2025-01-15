// controllers/property_details.go
package controllers

import (
    "propertyAPI/details_service"
    "github.com/beego/beego/v2/server/web"
)

type PropertyDetailsController struct {
    web.Controller
}

func (c *PropertyDetailsController) Get() {
    details, err := details_service.GetPropertyDetails()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = details
    c.ServeJSON()
}