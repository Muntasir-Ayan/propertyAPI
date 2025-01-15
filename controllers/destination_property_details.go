// controllers/destination_property_details.go
package controllers

import (
    "propertyAPI/destination_service"
    "github.com/beego/beego/v2/server/web"
)

type DestinationPropertyDetailsController struct {
    web.Controller
}

func (c *DestinationPropertyDetailsController) Get() {
    destID := c.GetString("dest_id")
    if destID == "" {
        c.Data["json"] = map[string]interface{}{
            "error": "dest_id parameter is required",
        }
        c.ServeJSON()
        return
    }

    details, err := destination_service.GetPropertyDetailsByDestination(destID)
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